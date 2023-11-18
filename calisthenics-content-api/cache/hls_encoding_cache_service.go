package cache

import (
	"fmt"
)

type IHLSEncodingCacheService interface {
	Save(cache HLSEncodingCache)
	SaveAllSlice(caches []HLSEncodingCache)
	SaveAll(caches ...HLSEncodingCache)
	GetByID(ID string) (HLSEncodingCache, error)
	GetAllByIDsInSlice(IDs []string) []HLSEncodingCache
	GetAllByIDsIn(IDs ...string) []HLSEncodingCache
	GetAll() []HLSEncodingCache
	Remove(ID string)
	RemoveAll()
}

type hlSEncodingCacheService struct {
	cacheService ICacheService
	key          string
}

func NewHLSEncodingCacheService(cacheService ICacheService) IHLSEncodingCacheService {
	key := "HLS_ENCODING"
	cacheService.CreateCache(key)
	return &hlSEncodingCacheService{
		cacheService: cacheService,
		key:          key,
	}
}

func (h *hlSEncodingCacheService) Save(cache HLSEncodingCache) {
	h.Remove(cache.ID)
	h.cacheService.Set(h.key, fmt.Sprintf("%s:_", cache.ID), cache)
}

func (h *hlSEncodingCacheService) GetAllByIDsInSlice(IDs []string) []HLSEncodingCache {
	keys := make([]string, 0)
	for _, ID := range IDs {
		keys = append(keys, fmt.Sprintf("%s:_", ID))
	}
	values := h.cacheService.GetAllByIDIn(h.key, keys)

	cacheValues := make([]HLSEncodingCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(HLSEncodingCache))
	}
	return cacheValues
}

func (h *hlSEncodingCacheService) GetAllByIDsIn(IDs ...string) []HLSEncodingCache {
	return h.GetAllByIDsInSlice(IDs)
}

func (h *hlSEncodingCacheService) GetByID(ID string) (HLSEncodingCache, error) {
	value, b := h.cacheService.Get(h.key, fmt.Sprintf("%s:_", ID))
	if !b {
		return HLSEncodingCache{}, fmt.Errorf("value not found with key")
	}

	cacheValue := value.(HLSEncodingCache)
	return cacheValue, nil
}

func (h *hlSEncodingCacheService) GetAll() []HLSEncodingCache {
	values := h.cacheService.GetAll(h.key)
	cacheValues := make([]HLSEncodingCache, 0)
	for _, value := range values {
		cacheValues = append(cacheValues, value.(HLSEncodingCache))
	}
	return cacheValues
}

func (h *hlSEncodingCacheService) Remove(ID string) {
	h.cacheService.Delete(h.key, fmt.Sprintf("%s:_", ID))
}

func (h *hlSEncodingCacheService) RemoveAll() {
	h.cacheService.DeleteKey(h.key)
	h.cacheService.CreateCache(h.key)
}

func (h *hlSEncodingCacheService) SaveAllSlice(caches []HLSEncodingCache) {
	for _, value := range caches {
		h.Remove(value.ID)
	}

	for _, value := range caches {
		h.cacheService.Set(h.key, fmt.Sprintf("%s:_", value.ID), value)
	}
}

func (h *hlSEncodingCacheService) SaveAll(caches ...HLSEncodingCache) {
	h.SaveAllSlice(caches)
}
