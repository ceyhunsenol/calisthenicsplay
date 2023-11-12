package calisthenics

import "time"

type ErrorResponse struct {
	Message string `json:"message"`
}

// UserResponse response
type UserResponse struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Username    string               `json:"username"`
	Email       string               `json:"email"`
	ProfileInfo *UserProfileResponse `json:"profileInfo"`
}

// UserProfileResponse response
type UserProfileResponse struct {
	ID          string    `json:"id"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	AvatarURL   string    `json:"avatarURL"`
	Bio         string    `json:"bio"`
}
