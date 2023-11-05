package data

type Content struct {
	BaseModel
	Code            string  `json:"code"`
	DescriptionCode string  `json:"description"`
	Active          bool    `json:"active"`
	Medias          []Media `json:"medias" gorm:"foreignKey:ContentID"`
	Genres          []Genre `json:"genres" gorm:"many2many:genre_contents;"`
	// Bu hareketin yardımcı hareketleri
	HelperContents []Content `gorm:"many2many:helper_contents;foreignKey:ID;joinForeignKey:ContentID;References:ID;JoinReferences:HelperContentID"`
	// Bu hareketin gereksinim duyduğu diğer hareketler
	RequirementContents []Content `gorm:"many2many:requirement_contents;foreignKey:ID;joinForeignKey:ContentID;References:ID;JoinReferences:RequirementContentID"`
}

type HelperContent struct {
	ContentID       string `json:"content_id"`
	HelperContentID string `json:"helper_content_id"`
}

type RequirementContent struct {
	ContentID            string `json:"content_id"`
	RequirementContentID string `json:"requirement_content_id"`
}

type Media struct {
	BaseModel
	DescriptionCode string `json:"description_code"`
	URL             string `json:"url"`
	Type            string `json:"type"`
	ContentID       string `gorm:"foreignKey:ID" json:"content_id"`
	Active          bool   `json:"active"`
}
