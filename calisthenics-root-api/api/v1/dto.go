package v1

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

// RoleDTO DTO
type RoleDTO struct {
	Code string `json:"code" validate:"required"`
}

// PrivilegeDTO DTO
type PrivilegeDTO struct {
	Code          string `json:"code" validate:"required"`
	Description   string `json:"description"`
	EndpointsJoin string `json:"endpointsJoin" validate:"required"`
}

// MediaDTO DTO
type MediaDTO struct {
	DescriptionCode string           `json:"descriptionCode"`
	URL             string           `json:"url" validate:"required"`
	Type            string           `json:"type" validate:"required"`
	ContentID       string           `json:"contentID" validate:"required"`
	Active          bool             `json:"active"`
	Translations    []TranslationDTO `json:"translations"`
}

// ContentDTO DTO
type ContentDTO struct {
	Code         string           `json:"code" validate:"required"`
	Description  string           `json:"description"`
	Active       bool             `json:"active"`
	Translations []TranslationDTO `json:"translations"`
}

// HelperContentDTO DTO
type HelperContentDTO struct {
	HelperContentID string `json:"helperContentID" validate:"required"`
}

// RequirementContentDTO DTO
type RequirementContentDTO struct {
	RequirementContentID string `json:"requirementContentID" validate:"required"`
}

// GenreContentDTO DTO
type GenreContentDTO struct {
	ContentID string `json:"contentID" validate:"required"`
}

// GenreTypeDTO DTO
type GenreTypeDTO struct {
	Code string `json:"code" validate:"required"`
}

// GenreDTO DTO
type GenreDTO struct {
	Type         string                  `json:"type" validate:"required"`
	Code         string                  `json:"code" validate:"required"`
	Description  string                  `json:"description"`
	Section      string                  `json:"section"`
	Active       bool                    `json:"active"`
	Translations []ContentTranslationDTO `json:"translations"`
}

// TranslationDTO DTO
type TranslationDTO struct {
	Code      string `json:"code" validate:"required"`
	LangCode  string `json:"langCode" validate:"required"`
	Translate string `json:"translate" validate:"required"`
	Active    bool   `json:"active"`
	Domain    string `json:"domain" validate:"required"`
}

// ContentTranslationDTO DTO
type ContentTranslationDTO struct {
	Code      string `json:"code"`
	LangCode  string `json:"langCode"`
	Translate string `json:"translate"`
	Active    bool   `json:"active"`
	ContentID string `json:"contentID"`
}

// ContentAccessDTO DTO
type ContentAccessDTO struct {
	ContentID string `json:"contentID" validate:"required"`
	Audience  string `json:"audience" validate:"required"`
}

// MediaAccessDTO DTO
type MediaAccessDTO struct {
	MediaID  string `json:"mediaID" validate:"required"`
	Audience string `json:"audience" validate:"required"`
}
