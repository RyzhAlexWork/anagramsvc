package cache

import (
	"sync"

	"anagramsvc/pkg/api/v1"
)

//GenerateCache ...
type GenerateCache interface {
	Get() (data *[]string)
	Set(data []string) (message string)
}

type generateCache struct {
	sync.RWMutex
	data []string
}

// Позволяет получить данные и кэша
func (g *generateCache) Get() (data *[]string) {
	g.RLock()
	data = &g.data
	g.RUnlock()
	return
}

// Кладёт данные в кэш
func (g *generateCache) Set(data []string) (message string) {
	g.Lock()
	g.data = data
	g.Unlock()
	message = v1.LoadMessage
	return
}

//NewGenerateCache ...
func NewGenerateCache() (_ GenerateCache) {
	cache := &generateCache{
		RWMutex: sync.RWMutex{},
		data:    nil,
	}
	return cache
}
