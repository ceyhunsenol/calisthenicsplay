package model

type VideoURLRequest struct {
	MediaID      string
	Token        string
	UserAgent    string
	Host         string
	CallerIP     string
	PlatformType string
	LangCode     string
}

type VideoURLResolutionModel struct {
	Height       int
	Bandwidth    int
	AvgBandwidth int
}

type VideoURLModel struct {
	VideoURL    string
	Resolutions []VideoURLResolutionModel
}

type VideoPlaylistRequest struct {
	Resolution string
	Token      string
}

type HLSMediaTokenModel struct {
	MediaID      string
	EncodingID   string
	UserAgent    string
	CallerIP     string
	PlatformType string
	Host         string
}

type TokenValidationRequest struct {
	UserAgent string
	CallerIP  string
}
