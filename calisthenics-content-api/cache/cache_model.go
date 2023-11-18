package cache

import (
	"time"
)

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
	EncodingID           string
}

// GeneralInfoCache cache
type GeneralInfoCache struct {
	Key   string
	Value string
}

// ContentAccessCache cache
type ContentAccessCache struct {
	ContentID string
	Audience  string
}

// MediaAccessCache cache
type MediaAccessCache struct {
	MediaID  string
	Audience string
}

// HLSEncodingCache cache
type HLSEncodingCache struct {
	ID         string
	LicenseKey string
	MediaID    string
	Files      []HLSEncodingFileCache
}

// HLSEncodingFileCache cache
type HLSEncodingFileCache struct {
	FileName   string
	EncodingID string
	IV         string
	Ext        float64
}

// LimitedCache cache
type LimitedCache struct {
	Key     string
	Value   interface{}
	endDate time.Time
}

func NewLimitedCache(key string, value interface{}) LimitedCache {
	return LimitedCache{
		Key:     key,
		Value:   value,
		endDate: time.Now().Add(140 * time.Hour),
	}
}

func (l *LimitedCache) GetLimitedEndDate() time.Time {
	return l.endDate
}
