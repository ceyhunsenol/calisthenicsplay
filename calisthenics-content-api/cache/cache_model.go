package cache

// GenreCache cache
type GenreCache struct {
	ID                   string
	Type                 string
	CodeMultiLang        *MultiLangCache
	DescriptionMultiLang *MultiLangCache
	Active               bool
	ContentIDs           []string
}

// ContentCache cache
type ContentCache struct {
	ID                    string
	CodeMultiLang         *MultiLangCache
	Description           string
	Active                bool
	HelperContentIDs      []string
	RequirementContentIDs []string
}

// MediaCache cache
type MediaCache struct {
	ID                   string
	DescriptionMultiLang *MultiLangCache
	URL                  string
	Type                 string
	Active               bool
	ContentID            string
}

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
