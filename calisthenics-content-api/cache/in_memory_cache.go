package cache

import (
	"sync"
)

type inMemoryCache struct {
	caches map[string]*cache
	mu     sync.RWMutex
}

type cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCacheManager() ICacheService {
	return &inMemoryCache{
		caches: make(map[string]*cache),
	}
}

func (cm *inMemoryCache) CreateCache(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.caches[key] = &cache{
		data: make(map[string]interface{}),
	}
}

func (cm *inMemoryCache) GetCache(key string) map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	cacheData, ok := cm.caches[key]
	if !ok {
		return nil
	}
	return cacheData.data
}

func (cm *inMemoryCache) Set(key, cacheKey string, value interface{}) {
	cm.mu.RLock()
	cacheData, ok := cm.caches[key]
	cm.mu.RUnlock()

	if !ok {
		cm.mu.Lock()
		cm.caches[key] = &cache{
			data: make(map[string]interface{}),
		}
		cm.mu.Unlock()
	}

	cacheData.mu.Lock()
	cacheData.data[cacheKey] = value
	cacheData.mu.Unlock()
}

func (cm *inMemoryCache) Get(key, cacheKey string) (interface{}, bool) {
	cm.mu.RLock()
	cacheData, ok := cm.caches[key]
	cm.mu.RUnlock()

	if !ok {
		return nil, false
	}

	cacheData.mu.RLock()
	value, ok := cacheData.data[cacheKey]
	cacheData.mu.RUnlock()

	return value, ok
}

func (cm *inMemoryCache) GetAll(key string) map[string]interface{} {
	cm.mu.RLock()
	cacheData, ok := cm.caches[key]
	cm.mu.RUnlock()

	if !ok {
		return make(map[string]interface{})
	}

	cacheData.mu.RLock()
	dataCopy := make(map[string]interface{})
	for k, v := range cacheData.data {
		dataCopy[k] = v
	}
	cacheData.mu.RUnlock()

	return dataCopy
}

func (cm *inMemoryCache) GetAllByIDIn(key string, ids []string) map[string]interface{} {
	cm.mu.RLock()
	cacheData, ok := cm.caches[key]
	cm.mu.RUnlock()

	if !ok {
		return make(map[string]interface{})
	}

	cacheData.mu.RLock()
	dataCopy := make(map[string]interface{})
	for _, id := range ids {
		if value, ok := cacheData.data[id]; ok {
			dataCopy[id] = value
		}
	}
	cacheData.mu.RUnlock()

	return dataCopy
}

func (cm *inMemoryCache) Delete(key, cacheKey string) {
	cm.mu.RLock()
	cacheData, ok := cm.caches[key]
	cm.mu.RUnlock()

	if ok {
		cacheData.mu.Lock()
		delete(cacheData.data, cacheKey)
		cacheData.mu.Unlock()
	}
}

func (cm *inMemoryCache) DeleteKey(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.caches, key)
}
