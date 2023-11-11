package calisthenics

type RefreshRequest struct {
	CacheType string
	ID        string
}

type ErrorResponse struct {
	Message string `json:"message"`
}
