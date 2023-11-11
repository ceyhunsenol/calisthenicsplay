package cache

import (
	"fmt"
)

type IGeneralInfoCacheService interface {
	Save(cache GeneralInfoCache)
	SaveAllSlice(caches []GeneralInfoCache)
	SaveAll(caches ...GeneralInfoCache)
	GetByID(ID string) (string, error)
	GetAllByIDsInSlice(IDs []string) []string
	GetAllByIDsIn(IDs ...string) []string
	GetAll() []string
	Remove(ID string)
	RemoveAll()
}

type generalInfoCacheService struct {
	cacheService ICacheService
	key          string
}

func NewGeneralInfoCacheService(cacheService ICacheService) IGeneralInfoCacheService {
	key := "GENERAL_INFO"
	cacheService.CreateCache(key)
	return &generalInfoCacheService{
		cacheService: cacheService,
		key:          key,
	}
}

func (c *generalInfoCacheService) Save(cache GeneralInfoCache) {
	c.Remove(cache.Key)
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.Key), cache.Value)
}

func (c *generalInfoCacheService) SaveAllSlice(caches []GeneralInfoCache) {
	for _, value := range caches {
		c.Remove(value.Key)
	}

	for _, value := range caches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.Key), value.Value)
	}
}

func (c *generalInfoCacheService) SaveAll(caches ...GeneralInfoCache) {
	c.SaveAllSlice(caches)
}

func (c *generalInfoCacheService) GetByID(ID string) (string, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return "", fmt.Errorf("value not found with key")
	}

	cacheValue := value.(string)
	return cacheValue, nil
}

func (c *generalInfoCacheService) GetAllByIDsInSlice(IDs []string) []string {
	keys := make([]string, 0)
	for _, ID := range IDs {
		keys = append(keys, fmt.Sprintf("%s:_", ID))
	}
	values := c.cacheService.GetAllByIDIn(c.key, keys)

	cacheValues := make([]string, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(string))
	}
	return cacheValues
}

func (c *generalInfoCacheService) GetAllByIDsIn(IDs ...string) []string {
	return c.GetAllByIDsInSlice(IDs)
}

func (c *generalInfoCacheService) GetAll() []string {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]string, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(string))
	}
	return cacheValues
}

func (c *generalInfoCacheService) Remove(ID string) {
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
}

func (c *generalInfoCacheService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.CreateCache(c.key)
}
