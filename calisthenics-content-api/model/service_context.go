package model

type ServiceContext struct {
	Authorization string
	LangCode      string
	PlatformType  string
	ClientIP      string
	CallerIP      string
	ClientIPList  []string
	UserAgent     string
	Host          string
}
