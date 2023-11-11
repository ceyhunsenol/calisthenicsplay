package api

import "time"

// TokenResource resource
type TokenResource struct {
	Username     string `json:"username"`
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken"`
}

func NewTokenResource(username, accessToken, refreshToken string) *TokenResource {
	return &TokenResource{
		Username:     username,
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
	}
}

// UserResource resource
type UserResource struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Username    string               `json:"username"`
	Email       string               `json:"email"`
	ProfileInfo *UserProfileResource `json:"profileInfo"`
}

// UserProfileResource resource
type UserProfileResource struct {
	ID          string    `json:"id"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	AvatarURL   string    `json:"avatarURL"`
	Bio         string    `json:"bio"`
}

// MessageResource resource
type MessageResource struct {
	Message string `json:"message"`
}
