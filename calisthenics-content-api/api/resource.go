package api

// MessageResource resource
type MessageResource struct {
	Message string `json:"message"`
}

// GenreResource resource
type GenreResource struct {
	ID          string
	Type        string
	Code        string
	Description string
	Section     string
	Active      bool
	Contents    []string
}

// ContentResource resource
type ContentResource struct {
	ID string
}

// VideoURLResolutionResource resource
type VideoURLResolutionResource struct {
	Height       int
	Bandwidth    int
	AvgBandwidth int
}

// VideoURLResource resource
type VideoURLResource struct {
	VideoURL    string                       `json:"videoURL"`
	Resolutions []VideoURLResolutionResource `json:"resolutions"`
}
