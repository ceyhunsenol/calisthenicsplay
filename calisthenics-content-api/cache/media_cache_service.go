package cache

import (
	"calisthenics-content-api/pkg"
	"fmt"
	"strings"
)

type MediaCacheService struct {
	cacheService ICacheService
	key          string
	keyContent   string
}

func NewMediaCacheService(cacheService ICacheService) *MediaCacheService {
	key := "MEDIA"
	keyContent := key + ":CONTENT_ID"
	cacheService.CreateCache(key)
	cacheService.CreateCache(keyContent)
	return &MediaCacheService{
		cacheService: cacheService,
		key:          key,
		keyContent:   keyContent,
	}
}

func (c *MediaCacheService) Save(cache MediaCache) {
	c.Remove(cache.ID)
	if !cache.Active {
		return
	}
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ID), cache)
	IDs := c.GetAllByContentID(cache.ContentID, cache.Type)
	pkg.AddIfNotExists(&IDs, cache.ID)
	c.cacheService.Set(c.keyContent, fmt.Sprintf(":%s:%s", cache.ContentID, cache.Type), IDs)
}

func (c *MediaCacheService) SaveAllSlice(caches []MediaCache) {
	for _, value := range caches {
		c.Remove(value.ID)
	}
	activeCaches := make([]MediaCache, 0)
	for _, value := range caches {
		if value.Active {
			activeCaches = append(activeCaches, value)
		}
	}

	for _, value := range activeCaches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.ID), value)
	}

	grouped := pkg.GroupByField(activeCaches, func(c MediaCache) string {
		return c.ContentID + ":" + c.Type
	})

	for s, value := range grouped {
		split := strings.Split(s, ":")
		IDs := c.GetAllByContentID(split[0], split[1])
		for _, ca := range value {
			pkg.AddIfNotExists(&IDs, ca.ID)
		}
		c.cacheService.Set(c.keyContent, fmt.Sprintf(":%s:%s", split[0], split[1]), IDs)
	}
}

func (c *MediaCacheService) SaveAll(caches ...MediaCache) {
	c.SaveAllSlice(caches)
}

func (c *MediaCacheService) GetByID(ID string) (MediaCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return MediaCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(MediaCache)
	return cacheValue, nil
}

func (c *MediaCacheService) GetAllByIDsInSlice(IDs []string) []MediaCache {
	keys := make([]string, 0)
	for _, ID := range IDs {
		keys = append(keys, fmt.Sprintf("%s:_", ID))
	}
	values := c.cacheService.GetAllByIDIn(c.key, keys)

	cacheValues := make([]MediaCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(MediaCache))
	}
	return cacheValues
}

func (c *MediaCacheService) GetAllByIDsIn(IDs ...string) []MediaCache {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *MediaCacheService) GetAll() []MediaCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]MediaCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(MediaCache))
	}
	return cacheValues
}

func (c *MediaCacheService) Remove(ID string) {
	value, err := c.GetByID(ID)
	if err != nil {
		return
	}
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
	IDs := c.GetAllByContentID(value.ContentID, value.Type)
	pkg.RemoveIfExists(&IDs, ID)
	c.cacheService.Set(c.keyContent, fmt.Sprintf(":%s:%s", value.ContentID, value.Type), IDs)
}

func (c *MediaCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.DeleteKey(c.keyContent)
	c.cacheService.CreateCache(c.key)
	c.cacheService.CreateCache(c.keyContent)
}

func (c *MediaCacheService) GetAllByContentID(contentID, videoType string) []string {
	value, b := c.cacheService.Get(c.keyContent, fmt.Sprintf(":%s:%s", contentID, videoType))
	if !b {
		return make([]string, 0)
	}
	return value.([]string)
}
