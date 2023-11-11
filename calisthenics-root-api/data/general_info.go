package data

type GeneralInfo struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	InfoKey         string `json:"info_key"`
	InfoValue       string `json:"info_value"`
	InfoDescription string `json:"info_description"`
}
