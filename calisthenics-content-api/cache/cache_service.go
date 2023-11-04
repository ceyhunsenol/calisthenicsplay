package cache

type ICacheService interface {
	CreateCache(key string)
	GetCache(key string) map[string]interface{}
	Set(key, cacheKey string, value interface{})
	Get(key, cacheKey string) (interface{}, bool)
	GetAll(key string) map[string]interface{}
	GetAllByIDIn(key string, ids []string) map[string]interface{}
	Delete(key, cacheKey string)
	DeleteKey(key string)
}

func GroupByField[T any](items []T, getField func(T) string) map[string][]T {
	grouped := make(map[string][]T)

	for _, item := range items {
		key := getField(item)
		group, exists := grouped[key]
		if !exists {
			grouped[key] = []T{item}
		} else {
			grouped[key] = append(group, item)
		}
	}

	return grouped
}
