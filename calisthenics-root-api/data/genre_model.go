package data

type GenreType struct {
	BaseModel
	Code string `json:"code"`
}

type Genre struct {
	BaseModel
	Type        string    `json:"type"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Contents    []Content `json:"contents" gorm:"many2many:genre_content;"`
}

type GenreContent struct {
	GenreID   string `json:"genre_id"`
	ContentID string `json:"content_id"`
}
