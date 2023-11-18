package cache

import (
	"fmt"
)

type ITranslationCacheService interface {
	Save(cache MultiLangCache)
	SaveAllSlice(caches []MultiLangCache)
	SaveAll(caches ...MultiLangCache)
	GetMultiLangByID(ID string) (MultiLangCache, error)
	GetByLang(ID, langCode string) string
	Remove(ID string)
	RemoveAll()
}

type translationService struct {
	cacheService ICacheService
	key          string
}

func NewTranslationCacheService(cacheService ICacheService) ITranslationCacheService {
	key := "GENRE"
	cacheService.CreateCache(key)
	return &translationService{
		cacheService: cacheService,
		key:          key,
	}
}

func (c *translationService) Save(cache MultiLangCache) {
	c.Remove(cache.Code)
	c.cacheService.Set(c.key, fmt.Sprintf("%s:_", cache.Code), cache)
}

func (c *translationService) SaveAllSlice(caches []MultiLangCache) {
	for _, value := range caches {
		c.Remove(value.Code)
	}

	for _, value := range caches {
		c.cacheService.Set(c.key, fmt.Sprintf("%s:_", value.Code), value)
	}
}

func (c *translationService) SaveAll(caches ...MultiLangCache) {
	c.SaveAllSlice(caches)
}

func (c *translationService) GetMultiLangByID(ID string) (MultiLangCache, error) {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return MultiLangCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(MultiLangCache)
	return cacheValue, nil
}

func (c *translationService) GetByLang(ID, langCode string) string {
	value, b := c.cacheService.Get(c.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return ""
	}
	cacheValue := value.(MultiLangCache)
	return cacheValue.GetByLang(langCode)
}

func (c *translationService) Remove(ID string) {
	c.cacheService.Delete(c.key, fmt.Sprintf("%s:_", ID))
}

func (c *translationService) RemoveAll() {
	c.cacheService.DeleteKey(c.key)
	c.cacheService.CreateCache(c.key)
}
