package cache

import (
	"fmt"
)

type GenreCacheService struct {
	cacheService ICacheService
	key          string
	keyType      string
}

func NewGenreCacheService(cacheService ICacheService) *GenreCacheService {
	key := "GENRE"
	keyType := key + ":TYPE"
	cacheService.CreateCache(key)
	cacheService.CreateCache(keyType)
	return &GenreCacheService{
		cacheService: cacheService,
		key:          key,
		keyType:      keyType,
	}
}

func (c *GenreCacheService) Save(cache GenreCache) {
	c.Remove(cache.ID)
	if !cache.Active {
		return
	}
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ID), cache)
	IDs := c.GetAllByType(cache.Type)
	AddIfNotExists(&IDs, cache.ID)
	c.cacheService.Set(c.keyType, fmt.Sprintf(":%s", cache.Type), IDs)
}

func (c *GenreCacheService) SaveAllSlice(caches []GenreCache) {
	for _, value := range caches {
		c.Remove(value.ID)
	}
	activeCaches := make([]GenreCache, 0)
	for _, value := range caches {
		if value.Active {
			activeCaches = append(activeCaches, value)
		}
	}

	for _, value := range activeCaches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.ID), value)
	}

	grouped := GroupByField(activeCaches, func(c GenreCache) string {
		return c.Type
	})

	for key, value := range grouped {
		IDs := c.GetAllByType(key)
		for _, ca := range value {
			AddIfNotExists(&IDs, ca.ID)
		}
		c.cacheService.Set(c.keyType, fmt.Sprintf(":%s", key), IDs)
	}
}

func (c *GenreCacheService) SaveAll(caches ...GenreCache) {
	c.SaveAllSlice(caches)
}

func (c *GenreCacheService) GetByID(ID string) (GenreCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return GenreCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(GenreCache)
	return cacheValue, nil
}

func (c *GenreCacheService) GetAllByIDsInSlice(IDs []string) []GenreCache {
	keys := make([]string, 0)
	for _, ID := range IDs {
		keys = append(keys, fmt.Sprintf("%s:_", ID))
	}
	values := c.cacheService.GetAllByIDIn(c.key, keys)

	cacheValues := make([]GenreCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(GenreCache))
	}
	return cacheValues
}

func (c *GenreCacheService) GetAllByIDsIn(IDs ...string) []GenreCache {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *GenreCacheService) GetAll() []GenreCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]GenreCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(GenreCache))
	}
	return cacheValues
}

func (c *GenreCacheService) Remove(ID string) {
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
}

func (c *GenreCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.CreateCache(c.key)
}

func (c *GenreCacheService) GetAllByType(genreType string) []string {
	value, b := c.cacheService.Get(c.keyType, fmt.Sprintf(":%s", genreType))
	if !b {
		return make([]string, 0)
	}
	return value.([]string)
}
