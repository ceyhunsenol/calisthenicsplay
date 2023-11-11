package cache

import (
	"calisthenics-content-api/pkg"
	"fmt"
)

type IGenreCacheService interface {
	Save(cache GenreCache)
	SaveAllSlice(caches []GenreCache)
	SaveAll(caches ...GenreCache)
	GetByID(ID string) (GenreCache, error)
	GetAllByIDsInSlice(IDs []string) []GenreCache
	GetAllByIDsIn(IDs ...string) []GenreCache
	GetAll() []GenreCache
	Remove(ID string)
	RemoveAll()
	GetAllIDsByType(genreType string) []string
	GetAllByType(genreType string) []GenreCache
}

type genreCacheService struct {
	cacheService ICacheService
	key          string
	keyType      string
}

func NewGenreCacheService(cacheService ICacheService) IGenreCacheService {
	key := "GENRE"
	keyType := key + ":TYPE"
	cacheService.CreateCache(key)
	cacheService.CreateCache(keyType)
	return &genreCacheService{
		cacheService: cacheService,
		key:          key,
		keyType:      keyType,
	}
}

func (c *genreCacheService) Save(cache GenreCache) {
	c.Remove(cache.ID)
	if !cache.Active {
		return
	}
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.ID), cache)
	IDs := c.GetAllIDsByType(cache.Type)
	pkg.AddIfNotExists(&IDs, cache.ID)
	c.cacheService.Set(c.keyType, fmt.Sprintf(":%s", cache.Type), IDs)
}

func (c *genreCacheService) SaveAllSlice(caches []GenreCache) {
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

	grouped := pkg.GroupByField(activeCaches, func(c GenreCache) string {
		return c.Type
	})

	for key, value := range grouped {
		IDs := c.GetAllIDsByType(key)
		for _, ca := range value {
			pkg.AddIfNotExists(&IDs, ca.ID)
		}
		c.cacheService.Set(c.keyType, fmt.Sprintf(":%s", key), IDs)
	}
}

func (c *genreCacheService) SaveAll(caches ...GenreCache) {
	c.SaveAllSlice(caches)
}

func (c *genreCacheService) GetByID(ID string) (GenreCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return GenreCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(GenreCache)
	return cacheValue, nil
}

func (c *genreCacheService) GetAllByIDsInSlice(IDs []string) []GenreCache {
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

func (c *genreCacheService) GetAllByIDsIn(IDs ...string) []GenreCache {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *genreCacheService) GetAll() []GenreCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]GenreCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(GenreCache))
	}
	return cacheValues
}

func (c *genreCacheService) Remove(ID string) {
	value, err := c.GetByID(ID)
	if err != nil {
		return
	}
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
	IDs := c.GetAllIDsByType(value.Type)
	pkg.RemoveIfExists(&IDs, ID)
	c.cacheService.Set(c.keyType, fmt.Sprintf(":%s", value.Type), IDs)
}

func (c *genreCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.DeleteKey(c.keyType)
	c.cacheService.CreateCache(c.key)
	c.cacheService.CreateCache(c.keyType)
}

func (c *genreCacheService) GetAllIDsByType(genreType string) []string {
	value, b := c.cacheService.Get(c.keyType, fmt.Sprintf(":%s", genreType))
	if !b {
		return make([]string, 0)
	}
	return value.([]string)
}

func (c *genreCacheService) GetAllByType(genreType string) []GenreCache {
	IDs := c.GetAllIDsByType(genreType)
	return c.GetAllByIDsInSlice(IDs)
}
