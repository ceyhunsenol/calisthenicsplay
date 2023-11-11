package data

type ContentAccess struct {
	BaseModel
	ContentID string `json:"content_id"`
	Audience  string `json:"audience"`
}

type MediaAccess struct {
	BaseModel
	MediaID  string `json:"media_id"`
	Audience string `json:"audience"`
}
