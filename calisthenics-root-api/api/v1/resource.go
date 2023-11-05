package v1

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

// MessageResource resource
type MessageResource struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// UserResource resource
type UserResource struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// RoleResource resource
type RoleResource struct {
	ID         uint                ` json:"id"`
	Code       string              `json:"code"`
	Privileges []PrivilegeResource `json:"privileges"`
}

// PrivilegeResource resource
type PrivilegeResource struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	Description   string `json:"description"`
	EndpointsJoin string `json:"endpointsJoin"`
}

// MediaResource resource
type MediaResource struct {
	ID              string `json:"id"`
	DescriptionCode string `json:"description"`
	URL             string `json:"url"`
	Type            string `json:"type"`
	Active          bool   `json:"active"`
	ContentID       string `json:"contentID"`
}

// ContentResource resource
type ContentResource struct {
	ID          string                 `json:"id"`
	Code        string                 `json:"code"`
	Description string                 `json:"description"`
	Active      bool                   `json:"active"`
	Medias      []ContentMediaResource `json:"medias"`
}

// ContentMediaResource resource
type ContentMediaResource struct {
	ID              string `json:"id"`
	DescriptionCode string `json:"description"`
	URL             string `json:"url"`
	Type            string `json:"type"`
}

// GenreTypeResource resource
type GenreTypeResource struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

// GenreResource resource
type GenreResource struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Section     string `json:"section"`
	Active      bool   `json:"active"`
}

// TranslationResource Resource
type TranslationResource struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	LangCode  string `json:"langCode"`
	Translate string `json:"translate"`
	Active    bool   `json:"active"`
	Domain    string `json:"domain"`
}

// ContentTranslationResource Resource
type ContentTranslationResource struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	LangCode  string `json:"langCode"`
	Translate string `json:"translate"`
	Active    bool   `json:"active"`
	ContentID string `json:"contentID"`
}
