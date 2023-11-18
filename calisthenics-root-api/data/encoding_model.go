package data

type Encoding struct {
	BaseModel
	LicenseKey    string         `json:"license_key"`
	MediaID       string         `gorm:"foreignKey:MediaID"  json:"media_id"`
	EncodingFiles []EncodingFile `gorm:"foreignKey:EncodingID"`
}

type EncodingFile struct {
	BaseModel
	FileName   string  `json:"file_name"`
	EncodingID string  `gorm:"foreignKey:EncodingID;references:ID" json:"encoding_id"`
	IV         string  `json:"iv"`
	Ext        float64 `json:"ext"`
}
