package cache

import (
	"calisthenics-content-api/pkg"
	"fmt"
	"strings"
)

type IMediaCacheService interface {
	Save(cache MediaCache)
	SaveAllSlice(caches []MediaCache)
	SaveAll(caches ...MediaCache)
	GetByID(ID string) (MediaCache, error)
	GetAllByIDsInSlice(IDs []string) []MediaCache
	GetAllByIDsIn(IDs ...string) []MediaCache
	GetAll() []MediaCache
	Remove(ID string)
	RemoveAll()
	GetAllIDsByContentID(contentID, videoType string) []string
	GetAllByContentID(contentID, videoType string) []MediaCache
}

type mediaCacheService struct {
	cacheService ICacheService
	key          string
	keyContent   string
}

func NewMediaCacheService(cacheService ICacheService) IMediaCacheService {
	key := "MEDIA"
	keyContent := key + ":CONTENT_ID"
	cacheService.CreateCache(key)
	cacheService.CreateCache(keyContent)
	return &mediaCacheService{
		cacheService: cacheService,
		key:          key,
		keyContent:   keyContent,
	}
}

func (c *mediaCacheService) Save(cache MediaCache) {
	c.Remove(cache.ID)
	if !cache.Active {
		return
	}
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ID), cache)
	IDs := c.GetAllIDsByContentID(cache.ContentID, cache.Type)
	pkg.AddIfNotExists(&IDs, cache.ID)
	c.cacheService.Set(c.keyContent, fmt.Sprintf(":%s:%s", cache.ContentID, cache.Type), IDs)
}

func (c *mediaCacheService) SaveAllSlice(caches []MediaCache) {
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
		IDs := c.GetAllIDsByContentID(split[0], split[1])
		for _, ca := range value {
			pkg.AddIfNotExists(&IDs, ca.ID)
		}
		c.cacheService.Set(c.keyContent, fmt.Sprintf(":%s:%s", split[0], split[1]), IDs)
	}
}

func (c *mediaCacheService) SaveAll(caches ...MediaCache) {
	c.SaveAllSlice(caches)
}

func (c *mediaCacheService) GetByID(ID string) (MediaCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return MediaCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(MediaCache)
	return cacheValue, nil
}

func (c *mediaCacheService) GetAllByIDsInSlice(IDs []string) []MediaCache {
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

func (c *mediaCacheService) GetAllByIDsIn(IDs ...string) []MediaCache {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *mediaCacheService) GetAll() []MediaCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]MediaCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(MediaCache))
	}
	return cacheValues
}

func (c *mediaCacheService) Remove(ID string) {
	value, err := c.GetByID(ID)
	if err != nil {
		return
	}
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
	IDs := c.GetAllIDsByContentID(value.ContentID, value.Type)
	pkg.RemoveIfExists(&IDs, ID)
	c.cacheService.Set(c.keyContent, fmt.Sprintf(":%s:%s", value.ContentID, value.Type), IDs)
}

func (c *mediaCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.DeleteKey(c.keyContent)
	c.cacheService.CreateCache(c.key)
	c.cacheService.CreateCache(c.keyContent)
}

func (c *mediaCacheService) GetAllIDsByContentID(contentID, videoType string) []string {
	value, b := c.cacheService.Get(c.keyContent, fmt.Sprintf(":%s:%s", contentID, videoType))
	if !b {
		return make([]string, 0)
	}
	return value.([]string)
}

func (c *mediaCacheService) GetAllByContentID(contentID, videoType string) []MediaCache {
	IDs := c.GetAllIDsByContentID(contentID, videoType)
	return c.GetAllByIDsInSlice(IDs)
}
