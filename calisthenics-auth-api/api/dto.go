package api

// LoginDTO DTO
type LoginDTO struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RefreshTokenDTO DTO
type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// RegisterDTO DTO
type RegisterDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserProfileDTO DTO
type UserProfileDTO struct {
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
	AvatarURL   string `json:"avatarURL" validate:"required"`
	Bio         string `json:"bio" validate:"required"`
}
