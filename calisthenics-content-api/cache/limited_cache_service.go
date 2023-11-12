package cache

import (
	"fmt"
)

type ILimitedCacheService interface {
	Save(cache LimitedCache)
	GetByID(ID string) (interface{}, error)
	GetAll() []LimitedCache
	Remove(ID string)
}

type limitedCacheService struct {
	cacheService ICacheService
	key          string
}

func NewLimitedCacheService(cacheService ICacheService) ILimitedCacheService {
	key := "LIMITED_CACHE"
	cacheService.CreateCache(key)
	return &limitedCacheService{
		cacheService: cacheService,
		key:          key,
	}
}

func (c *limitedCacheService) Save(cache LimitedCache) {
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.Key), cache)
}

func (c *limitedCacheService) GetByID(ID string) (interface{}, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return LimitedCache{}.Value, fmt.Errorf("value not found with key")
	}

	v := value.(LimitedCache).Value
	return v, nil
}

func (c *limitedCacheService) GetAll() []LimitedCache {
	values := c.cacheService.GetAll(c.key)
	cacheValues := make([]LimitedCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(LimitedCache))
	}
	return cacheValues
}

func (c *limitedCacheService) Remove(ID string) {
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
}
