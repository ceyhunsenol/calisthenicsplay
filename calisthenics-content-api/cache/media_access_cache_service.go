package cache

import (
	"fmt"
)

type IMediaAccessCacheService interface {
	Save(cache MediaAccessCache)
	SaveAllSlice(caches []MediaAccessCache)
	SaveAll(caches ...MediaAccessCache)
	GetByID(ID string) (MediaAccessCache, error)
	GetAllByIDsInSlice(IDs []string) []MediaAccessCache
	GetAllByIDsIn(IDs ...string) []MediaAccessCache
	GetAll() []MediaAccessCache
	Remove(ID string)
	RemoveAll()
}

type mediaAccessCacheService struct {
	cacheService ICacheService
	key          string
}

func NewMediaAccessCacheService(cacheService ICacheService) IMediaAccessCacheService {
	key := "MEDIA_ACCESS"
	cacheService.CreateCache(key)
	return &mediaAccessCacheService{
		cacheService: cacheService,
		key:          key,
	}
}

func (c *mediaAccessCacheService) Save(cache MediaAccessCache) {
	c.Remove(cache.ID)
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ID), cache)
}

func (c *mediaAccessCacheService) SaveAllSlice(caches []MediaAccessCache) {
	for _, value := range caches {
		c.Remove(value.ID)
	}

	for _, value := range caches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.ID), value)
	}
}

func (c *mediaAccessCacheService) SaveAll(caches ...MediaAccessCache) {
	c.SaveAllSlice(caches)
}

func (c *mediaAccessCacheService) GetByID(ID string) (MediaAccessCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return MediaAccessCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(MediaAccessCache)
	return cacheValue, nil
}

func (c *mediaAccessCacheService) GetAllByIDsInSlice(IDs []string) []MediaAccessCache {
	keys := make([]string, 0)
	for _, ID := range IDs {
		keys = append(keys, fmt.Sprintf("%s:_", ID))
	}
	values := c.cacheService.GetAllByIDIn(c.key, keys)

	cacheValues := make([]MediaAccessCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(MediaAccessCache))
	}
	return cacheValues
}

func (c *mediaAccessCacheService) GetAllByIDsIn(IDs ...string) []MediaAccessCache {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *mediaAccessCacheService) GetAll() []MediaAccessCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]MediaAccessCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(MediaAccessCache))
	}
	return cacheValues
}

func (c *mediaAccessCacheService) Remove(ID string) {
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
}

func (c *mediaAccessCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.CreateCache(c.key)
}
