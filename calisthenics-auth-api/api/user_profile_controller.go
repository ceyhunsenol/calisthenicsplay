package api

import (
	"calisthenics-auth-api/middleware"
	"calisthenics-auth-api/model"
	"calisthenics-auth-api/pkg"
	"calisthenics-auth-api/service"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

type UserProfileController struct {
	userProfileService   service.IUserProfileService
	accessTokenSecretKey string
}

func NewUserProfileController(userProfileService service.IUserProfileService) *UserProfileController {
	return &UserProfileController{
		userProfileService:   userProfileService,
		accessTokenSecretKey: viper.GetString("security.jwt.access-token-secret-key"),
	}
}

func (u *UserProfileController) InitUserProfileRoutes(e *echo.Group) {
	e.PUT("/profile", u.UserProfile)
}

func (u *UserProfileController) UserProfile(c echo.Context) error {
	userProfileDTO := UserProfileDTO{}
	if err := c.Bind(&userProfileDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&userProfileDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	context, _ := middleware.GetServiceContextByEchoContext(c)
	token, _ := pkg.ClearToken(context.Authorization)
	userID, err := pkg.GetUserIDFromToken(token, u.accessTokenSecretKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &MessageResource{Code: http.StatusUnauthorized, Message: err.Error()})
	}
	request := model.UserProfileRequest{
		UserID:      userID,
		DateOfBirth: userProfileDTO.DateOfBirth,
		AvatarURL:   userProfileDTO.AvatarURL,
		Bio:         userProfileDTO.Bio,
	}
	serviceError := u.userProfileService.UserProfile(request)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	return c.JSON(http.StatusNoContent, nil)
}
