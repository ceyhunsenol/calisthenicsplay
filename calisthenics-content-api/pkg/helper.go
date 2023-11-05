package pkg

func AddIfNotExists(IDs *[]string, element string) {
	found := false
	for _, id := range *IDs {
		if id == element {
			found = true
			break
		}
	}
	if !found {
		*IDs = append(*IDs, element)
	}
}

func RemoveIfExists(IDs *[]string, element string) {
	for i, id := range *IDs {
		if id == element {
			*IDs = append((*IDs)[:i], (*IDs)[i+1:]...)
			return
		}
	}
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
