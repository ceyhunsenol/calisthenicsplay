package api

import (
	"calisthenics-auth-api/model"
	"calisthenics-auth-api/service"
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

	v1.POST("/register", u.Register)
	v1.POST("/login", u.Login)
	v1.POST("/refresh-token", u.RefreshToken)
}

func (u *AuthController) Register(c echo.Context) error {
	registerDTO := RegisterDTO{}
	if err := c.Bind(&registerDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&registerDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	request := model.RegisterRequest{
		Username: registerDTO.Username,
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
	}
	tokenModel, err := u.authService.Register(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	return c.JSON(http.StatusCreated, NewTokenResource(tokenModel.Username, tokenModel.AccessToken, tokenModel.RefreshToken))
}

func (u *AuthController) Login(c echo.Context) error {
	loginDTO := LoginDTO{}
	if err := c.Bind(&loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	if loginDTO.Email == "" && loginDTO.Username == "" {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "email or username must be provided."})
	}
	request := model.LoginRequest{
		Username: loginDTO.Username,
		Email:    loginDTO.Email,
		Password: loginDTO.Password,
	}
	tokenModel, err := u.authService.Login(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	return c.JSON(http.StatusCreated, NewTokenResource(tokenModel.Username, tokenModel.AccessToken, tokenModel.RefreshToken))
}

func (u *AuthController) RefreshToken(c echo.Context) error {
	refreshTokenDTO := RefreshTokenDTO{}
	if err := c.Bind(&refreshTokenDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&refreshTokenDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	tokenModel, err := u.authService.RefreshToken(model.RefreshTokenRequest{RefreshToken: refreshTokenDTO.RefreshToken})
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	return c.JSON(http.StatusCreated, NewTokenResource(tokenModel.Username, tokenModel.AccessToken, tokenModel.RefreshToken))
}
