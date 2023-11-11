package cache

// GenreCache cache
type GenreCache struct {
	ID                   string
	Type                 string
	CodeMultiLang        *MultiLangCache
	DescriptionMultiLang *MultiLangCache
	Section              string
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

// GeneralInfoCache cache
type GeneralInfoCache struct {
	ID    string
	Value string
}

// ContentAccessCache cache
type ContentAccessCache struct {
	ID        string
	ContentID string
	Audience  string
}

// MediaAccessCache cache
type MediaAccessCache struct {
	ID       string
	MediaID  string
	Audience string
}
