package api

import (
	"calisthenics-auth-api/middleware"
	"calisthenics-auth-api/pkg"
	"calisthenics-auth-api/service"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

type UserController struct {
	userService          service.IUserService
	accessTokenSecretKey string
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService:          userService,
		accessTokenSecretKey: viper.GetString("security.jwt.access-token-secret-key"),
	}
}

func (u *UserController) InitUserRoutes(e *echo.Group) {
	e.GET("/user-info", u.UserInfo)
}

func (u *UserController) UserInfo(c echo.Context) error {
	context, _ := middleware.GetServiceContextByEchoContext(c)
	token, _ := pkg.ClearToken(context.Authorization)
	userID, err := pkg.GetUserIDFromToken(token, u.accessTokenSecretKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &MessageResource{Code: http.StatusUnauthorized, Message: err.Error()})
	}
	user, err := u.userService.GetById(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &MessageResource{Code: http.StatusUnauthorized, Message: "Unauthorized."})
	}
	userResource := UserResource{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
	if user.Profile.ID != "" {
		userResource.ProfileInfo = &UserProfileResource{
			ID:          user.Profile.ID,
			DateOfBirth: user.Profile.DateOfBirth,
			AvatarURL:   user.Profile.AvatarURL,
			Bio:         user.Profile.Bio,
		}
	}
	return c.JSON(http.StatusOK, userResource)
}
