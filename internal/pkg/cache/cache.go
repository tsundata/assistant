package cache

import (
	"sync"
)

type InMemoryCache struct {
	mu   *sync.Mutex
	data map[string]interface{}
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{mu: &sync.Mutex{}, data: make(map[string]interface{})}
}

func (i *InMemoryCache) Set(key string, value interface{}) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.data[key] = value
}

func (i *InMemoryCache) Get(key string) (interface{}, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if value, ok := i.data[key]; ok {
		return value, true
	}
	return nil, false
}
