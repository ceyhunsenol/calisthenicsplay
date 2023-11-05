package model

// GenreRequest request
type GenreRequest struct {
	Type    string
	Section string
}

// GenreModel model
type GenreModel struct {
	ID          string
	Type        string
	Code        string
	Description string
	Section     string
	Active      bool
	Contents    []string
}
