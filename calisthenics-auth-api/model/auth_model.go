package model

// LoginRequest request
type LoginRequest struct {
	Username string
	Email    string
	Password string
}

// TokenModel model
type TokenModel struct {
	Username     string
	AccessToken  string
	TokenType    string
	RefreshToken string
}

func NewTokenModel(username, accessToken, refreshToken string) *TokenModel {
	return &TokenModel{
		Username:     username,
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
	}
}

// RefreshTokenRequest request
type RefreshTokenRequest struct {
	RefreshToken string
}

// RegisterRequest request
type RegisterRequest struct {
	Username string
	Email    string
	Password string
}
