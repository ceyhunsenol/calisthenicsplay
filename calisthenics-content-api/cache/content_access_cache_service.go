package cache

import (
	"fmt"
)

type IContentAccessCacheService interface {
	Save(cache ContentAccessCache)
	SaveAllSlice(caches []ContentAccessCache)
	SaveAll(caches ...ContentAccessCache)
	GetByID(ID string) (ContentAccessCache, error)
	GetAllByIDsInSlice(IDs []string) []ContentAccessCache
	GetAllByIDsIn(IDs ...string) []ContentAccessCache
	GetAll() []ContentAccessCache
	Remove(ID string)
	RemoveAll()
}

type contentAccessCacheService struct {
	cacheService ICacheService
	key          string
}

func NewContentAccessCacheService(cacheService ICacheService) IContentAccessCacheService {
	key := "CONTENT_ACCESS"
	cacheService.CreateCache(key)
	return &contentAccessCacheService{
		cacheService: cacheService,
		key:          key,
	}
}

func (c *contentAccessCacheService) Save(cache ContentAccessCache) {
	c.Remove(cache.ContentID)
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ContentID), cache)
}

func (c *contentAccessCacheService) SaveAllSlice(caches []ContentAccessCache) {
	for _, value := range caches {
		c.Remove(value.ContentID)
	}

	for _, value := range caches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.ContentID), value)
	}
}

func (c *contentAccessCacheService) SaveAll(caches ...ContentAccessCache) {
	c.SaveAllSlice(caches)
}

func (c *contentAccessCacheService) GetByID(ID string) (ContentAccessCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return ContentAccessCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(ContentAccessCache)
	return cacheValue, nil
}

func (c *contentAccessCacheService) GetAllByIDsInSlice(IDs []string) []ContentAccessCache {
	keys := make([]string, 0)
	for _, ID := range IDs {
		keys = append(keys, fmt.Sprintf("%s:_", ID))
	}
	values := c.cacheService.GetAllByIDIn(c.key, keys)

	cacheValues := make([]ContentAccessCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(ContentAccessCache))
	}
	return cacheValues
}

func (c *contentAccessCacheService) GetAllByIDsIn(IDs ...string) []ContentAccessCache {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *contentAccessCacheService) GetAll() []ContentAccessCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]ContentAccessCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(ContentAccessCache))
	}
	return cacheValues
}

func (c *contentAccessCacheService) Remove(ID string) {
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
}

func (c *contentAccessCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.CreateCache(c.key)
}
