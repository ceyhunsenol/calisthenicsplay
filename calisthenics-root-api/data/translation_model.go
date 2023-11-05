package data

type Translation struct {
	BaseModel
	Code      string `json:"code"`
	LangCode  string `json:"lang_code"`
	Translate string `json:"translate"`
	Active    bool   `json:"active"`
	Domain    string `json:"domain"`
}

type ContentTranslation struct {
	BaseModel
	Code      string `json:"code"`
	LangCode  string `json:"lang_code"`
	Translate string `json:"translate"`
	Active    bool   `json:"active"`
}
