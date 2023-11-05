package model

// ContentTranslationRequest Request
type ContentTranslationRequest struct {
	Code      string `json:"code"`
	LangCode  string `json:"langCode"`
	Translate string `json:"translate"`
	Active    bool   `json:"active"`
	ContentID string `json:"contentID"`
}
