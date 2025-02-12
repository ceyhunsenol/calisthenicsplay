package cache

type ICacheService interface {
	CreateCache(key string)
	GetCache(key string) map[string]interface{}
	Set(key, cacheKey string, value interface{})
	Get(key, cacheKey string) (interface{}, bool)
	GetAll(key string) []interface{}
	GetAllByIDIn(key string, ids []string) []interface{}
	Delete(key, cacheKey string)
	DeleteKey(key string)
}
