package cache

import (
	"fmt"
)

type ContentCacheService struct {
	cacheService ICacheService
	key          string
}

func NewContentCacheService(cacheService ICacheService) *ContentCacheService {
	key := "CONTENT"
	cacheService.CreateCache(key)
	return &ContentCacheService{
		cacheService: cacheService,
		key:          key,
	}
}

func (c *ContentCacheService) Save(cache ContentCache) {
	c.Remove(cache.ID)
	if !cache.Active {
		return
	}
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ID), cache)
}

func (c *ContentCacheService) SaveAllSlice(caches []ContentCache) {
	for _, ca := range caches {
		c.Remove(ca.ID)
	}
	activeCaches := make([]ContentCache, 0)
	// ileride bağlantılı nesneler gelebilir diye sadece active içerikler alınıyor.
	for _, value := range caches {
		if value.Active {
			activeCaches = append(activeCaches, value)
		}
	}
	for _, value := range activeCaches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.ID), value)
	}
}

func (c *ContentCacheService) SaveAll(caches ...ContentCache) {
	c.SaveAllSlice(caches)
}

func (c *ContentCacheService) GetByID(ID string) (ContentCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return ContentCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(ContentCache)
	return cacheValue, nil
}

func (c *ContentCacheService) GetAllByIDsInSlice(IDs []string) []ContentCache {
	keys := make([]string, 0)
	for _, ID := range IDs {
		keys = append(keys, fmt.Sprintf("%s:_", ID))
	}
	values := c.cacheService.GetAllByIDIn(c.key, keys)

	cacheValues := make([]ContentCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(ContentCache))
	}
	return cacheValues
}

func (c *ContentCacheService) GetAllByIDsIn(IDs ...string) []ContentCache {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *ContentCacheService) GetAll() []ContentCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]ContentCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(ContentCache))
	}
	return cacheValues
}

func (c *ContentCacheService) Remove(ID string) {
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
}

func (c *ContentCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.CreateCache(c.key)
}
