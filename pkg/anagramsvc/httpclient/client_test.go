// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/valyala/fasthttp"

	"anagramsvc/pkg/api/v1"
)

func Test_client_GetAnagrams(t *testing.T) {

	var searchWord *string
	_ = faker.FakeData(&searchWord)

	var response v1.GetAnagramsResponse
	_ = faker.FakeData(&response)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := response

		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportGetAnagrams := NewGetAnagramsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetAnagrams,
		httpMethodGetAnagrams,
	)

	type fields struct {
		cli                  *fasthttp.HostClient
		transportGetAnagrams GetAnagramsTransport
		options              map[interface{}]Option
	}
	type args struct {
		ctx        context.Context
		searchWord *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantResponse v1.GetAnagramsResponse

		wantErr bool
	}{
		{
			"test GetAnagrams",
			fields{hostClient, transportGetAnagrams, opts},
			args{context.Background(), searchWord},
			response,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                  tt.fields.cli,
				transportGetAnagrams: tt.fields.transportGetAnagrams,
				options:              tt.fields.options,
			}
			gotResponse, err := s.GetAnagrams(tt.args.ctx, tt.args.searchWord)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetAnagrams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("client.response() = %v, want %v", gotResponse, tt.wantResponse)
			}

		})
	}
}

func Test_client_LoadWords(t *testing.T) {

	var loadWords *[]string
	_ = faker.FakeData(&loadWords)

	var response v1.LoadWordsResponse
	_ = faker.FakeData(&response)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := response

		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportLoadWords := NewLoadWordsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientLoadWords,
		httpMethodLoadWords,
	)

	type fields struct {
		cli                *fasthttp.HostClient
		transportLoadWords LoadWordsTransport
		options            map[interface{}]Option
	}
	type args struct {
		ctx       context.Context
		loadWords *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantResponse v1.LoadWordsResponse

		wantErr bool
	}{
		{
			"test LoadWords",
			fields{hostClient, transportLoadWords, opts},
			args{context.Background(), loadWords},
			response,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                tt.fields.cli,
				transportLoadWords: tt.fields.transportLoadWords,
				options:            tt.fields.options,
			}
			gotResponse, err := s.LoadWords(tt.args.ctx, tt.args.loadWords)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LoadWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("client.response() = %v, want %v", gotResponse, tt.wantResponse)
			}

		})
	}
}

func TestNewClient(t *testing.T) {
	serverURL := fmt.Sprintf("https://%v.com", time.Now().UnixNano())
	parsedServerURL, _ := url.Parse(serverURL)
	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: rand.Int(),
	}
	opts := map[interface{}]Option{}

	transportGetAnagrams := NewGetAnagramsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetAnagrams,
		httpMethodGetAnagrams,
	)

	transportLoadWords := NewLoadWordsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientLoadWords,
		httpMethodLoadWords,
	)

	cl := &client{
		hostClient,
		transportGetAnagrams,
		transportLoadWords,
		opts,
	}

	type args struct {
		cli *fasthttp.HostClient

		transportGetAnagrams GetAnagramsTransport

		transportLoadWords LoadWordsTransport

		options map[interface{}]Option
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"test new client", args{hostClient, transportGetAnagrams, transportLoadWords, opts}, cl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.cli, tt.args.transportGetAnagrams, tt.args.transportLoadWords, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
