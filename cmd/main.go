package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buaazp/fasthttprouter"
	fasthttpprometheus "github.com/flf2ko/fasthttp-prometheus"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"

	"anagramsvc/pkg/anagramsvc"
	"anagramsvc/pkg/anagramsvc/httperrors"
	"anagramsvc/pkg/anagramsvc/httpserver"
	"anagramsvc/pkg/api/v1"
	"anagramsvc/pkg/cache"
	"anagramsvc/pkg/processor"
)

var (
	serviceVersion = "v1.0.0"
	errorCreator   = httperrors.NewError
	methodError    = []string{"method", "error"}
)

type configuration struct {
	Debug bool   `envconfig:"DEBUG" default:"true"`
	Port  string `envconfig:"PORT" default:"8080"`

	ReadTimeout        time.Duration `envconfig:"READ_TIMEOUT" default:"1s"`
	ReadBufferSize     int           `envconfig:"READ_BUFFER_SIZE" default:"163840"`
	MaxRequestBodySize int           `envconfig:"MAX_REQUEST_BODY_SIZE" default:"10485760"`

	/*Metrics*/
	MetricsNamespace    string `envconfig:"METRICS_NAMESPACE" default:"anagramsvc"`
	MetricsSubsystem    string `envconfig:"METRICS_SUBSYSTEM" default:"anagramsvc"`
	MetricsNameCount    string `envconfig:"METRICS_NAME_COUNT" default:"request_count"`
	MetricsHelpCount    string `envconfig:"METRICS_HELP_COUNT" default:"Request count"`
	MetricsNameDuration string `envconfig:"METRICS_NAME_DURATION" default:"request_duration"`
	MetricsHelpDuration string `envconfig:"METRICS_HELP_DURATION" default:"Request duration"`
}

type dependency struct {
	cfg               configuration
	logger            log.Logger
	prometheusCounter metrics.Counter
	prometheusSummary metrics.Histogram
	prometheus        *fasthttpprometheus.Prometheus
	pcs               processor.Processor
}

func main() {
	// logger
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", serviceVersion)

	// configuration
	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}
	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	/*Prometheus variables*/
	prometheusCounter := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: cfg.MetricsNamespace,
		Subsystem: cfg.MetricsSubsystem,
		Name:      cfg.MetricsNameCount,
		Help:      cfg.MetricsHelpCount,
	}, methodError)
	prometheusSummary := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: cfg.MetricsNamespace,
		Subsystem: cfg.MetricsSubsystem,
		Name:      cfg.MetricsNameDuration,
		Help:      cfg.MetricsHelpDuration,
	}, methodError)

	// cache
	cacheNew := cache.NewGenerateCache()

	// service
	processorNew := processor.NewProcessor(
		logger,
		cacheNew,
	)

	p := fasthttpprometheus.NewPrometheus(cfg.MetricsSubsystem)

	dep := &dependency{
		cfg:               cfg,
		logger:            logger,
		pcs:               processorNew,
		prometheus:        p,
		prometheusCounter: prometheusCounter,
		prometheusSummary: prometheusSummary,
	}

	errorProcessor := httperrors.NewErrorProcessor(v1.ErrMap)

	service := anagramsvc.NewService(
		dep.pcs,
	)
	service = anagramsvc.NewValidatorMiddleware(service, dep.cfg.Debug, errorCreator)
	service = anagramsvc.NewLoggingMiddleware(dep.logger, service)
	service = anagramsvc.NewInstrumentingMiddleware(dep.prometheusCounter,
		dep.prometheusSummary,
		service,
	)

	router := fasthttprouter.New()
	httpserver.New(
		router,
		service,
		httperrors.DecodeJSONErrorCreator,
		httperrors.EncodeJSONErrorCreator,
		httperrors.EncodeQueryTypeIntErrorCreator,
		errorProcessor,
	)

	serverHandler := dep.prometheus.WrapHandler(router)

	fasthttpServer := &fasthttp.Server{
		Handler:            serverHandler,
		MaxRequestBodySize: dep.cfg.MaxRequestBodySize,
		ReadBufferSize:     dep.cfg.ReadBufferSize,
		ReadTimeout:        dep.cfg.ReadTimeout,
	}

	go func() {
		_ = level.Info(dep.logger).Log("msg", "starting http server", "port", dep.cfg.Port)
		if err := fasthttpServer.ListenAndServe(":" + dep.cfg.Port); err != nil {
			_ = level.Error(dep.logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)

		if err := fasthttpServer.Shutdown(); err != nil {
			_ = level.Error(logger).Log("msg", "server shutdown failure", "err", err)
		}

		_ = level.Info(logger).Log("msg", "goodbye")
	}(<-c)
}
