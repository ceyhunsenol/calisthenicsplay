package cache

type CacheType string

const (
	Genre         CacheType = "genre"
	Content       CacheType = "content"
	Media         CacheType = "media"
	GeneralInfo   CacheType = "general_info"
	ContentAccess CacheType = "content_access"
	MediaAccess   CacheType = "media_access"
)
