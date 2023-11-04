package v1

import (
	"calisthenics-root-api/model"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (u *AuthController) InitAuthRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.POST("/login", u.Login)
	v1.POST("/refresh-token", u.RefreshToken)
}

func (u *AuthController) Login(c echo.Context) error {
	loginDTO := LoginDTO{}
	if err := c.Bind(&loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if loginDTO.Email == "" && loginDTO.Username == "" {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "email or username must be provided."})
	}
	request := model.LoginRequest{
		Username: loginDTO.Username,
		Email:    loginDTO.Email,
		Password: loginDTO.Password,
	}
	tokenModel, err := u.authService.Login(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}
	return c.JSON(http.StatusCreated, NewTokenResource(tokenModel.Username, tokenModel.AccessToken, tokenModel.RefreshToken))
}

func (u *AuthController) RefreshToken(c echo.Context) error {
	refreshTokenDTO := RefreshTokenDTO{}
	if err := c.Bind(&refreshTokenDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&refreshTokenDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	tokenModel, err := u.authService.RefreshToken(model.RefreshTokenRequest{RefreshToken: refreshTokenDTO.RefreshToken})
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}
	return c.JSON(http.StatusCreated, NewTokenResource(tokenModel.Username, tokenModel.AccessToken, tokenModel.RefreshToken))
}
