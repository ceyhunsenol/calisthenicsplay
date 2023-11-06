package model

// ContentTranslationRequest Request
type ContentTranslationRequest struct {
	ContentID    string `json:"contentID"`
	Translations []ContentTranslationModel
}

type ContentTranslationModel struct {
	Code      string `json:"code"`
	LangCode  string `json:"langCode"`
	Translate string `json:"translate"`
	Active    bool   `json:"active"`
}
