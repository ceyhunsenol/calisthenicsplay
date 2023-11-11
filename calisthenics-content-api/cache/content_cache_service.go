package cache

import (
	"fmt"
)

type IContentCacheService interface {
	Save(cache ContentCache)
	SaveAllSlice(caches []ContentCache)
	SaveAll(caches ...ContentCache)
	GetByID(ID string) (ContentCache, error)
	GetAllByIDsInSlice(IDs []string) []ContentCache
	GetAllByIDsIn(IDs ...string) []ContentCache
	GetAll() []ContentCache
	Remove(ID string)
	RemoveAll()
	GetIdByCode(code string) (string, error)
	GetByCode(code string) (ContentCache, error)
}

type ContentCacheService struct {
	cacheService ICacheService
	key          string
	keyCode      string
}

func NewContentCacheService(cacheService ICacheService) IContentCacheService {
	key := "CONTENT"
	keyCode := key + ":CONTENT_CODE"
	cacheService.CreateCache(key)
	cacheService.CreateCache(keyCode)
	return &ContentCacheService{
		cacheService: cacheService,
		key:          key,
		keyCode:      keyCode,
	}
}

func (c *ContentCacheService) Save(cache ContentCache) {
	c.Remove(cache.ID)
	if !cache.Active {
		return
	}
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ID), cache)
	c.cacheService.Set(c.keyCode, fmt.Sprintf(":%s", cache.CodeMultiLang.Code), cache.ID)
}

func (c *ContentCacheService) SaveAllSlice(caches []ContentCache) {
	for _, ca := range caches {
		c.Remove(ca.ID)
	}
	activeCaches := make([]ContentCache, 0)
	for _, value := range caches {
		if value.Active {
			activeCaches = append(activeCaches, value)
		}
	}
	for _, value := range activeCaches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.ID), value)
		c.cacheService.Set(c.keyCode, fmt.Sprintf(":%s", value.CodeMultiLang.Code), value.ID)
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
	value, err := c.GetByID(ID)
	if err != nil {
		return
	}
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
	c.cacheService.Delete(c.keyCode, fmt.Sprintf(":%s", value.CodeMultiLang.Code))
}

func (c *ContentCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.DeleteKey(c.keyCode)
	c.cacheService.CreateCache(c.key)
	c.cacheService.CreateCache(c.keyCode)
}

func (c *ContentCacheService) GetIdByCode(code string) (string, error) {
	value, b := c.cacheService.Get(c.keyCode, fmt.Sprintf(":%s", code))
	if !b {
		return "", fmt.Errorf("value not found with key")
	}

	cacheValue := value.(string)
	return cacheValue, nil
}

func (c *ContentCacheService) GetByCode(code string) (ContentCache, error) {
	id, err := c.GetIdByCode(code)
	if err != nil {
		return ContentCache{}, err
	}
	cac, err := c.GetByID(id)
	if err != nil {
		return ContentCache{}, err
	}
	return cac, nil
}
