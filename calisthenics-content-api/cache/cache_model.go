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
	DescriptionMultiLang  *MultiLangCache
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
