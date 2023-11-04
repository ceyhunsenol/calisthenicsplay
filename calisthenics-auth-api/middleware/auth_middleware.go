package middleware

import (
	"calisthenics-auth-api/model"
	"calisthenics-auth-api/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

type AuthenticationMiddleware struct {
	accessTokenSecretKey string
}

func NewAuthenticationMiddleware() *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		accessTokenSecretKey: viper.GetString("security.jwt.access-token-secret-key"),
	}
}

func (m *AuthenticationMiddleware) AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		serviceCtx, ok := c.Request().Context().Value("ServiceContextKey").(*model.ServiceContext)
		if !ok || serviceCtx.Authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "No authorization provided.")
		}
		token, err := pkg.ClearToken(serviceCtx.Authorization)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		_, err = pkg.GetUserIDFromToken(token, m.accessTokenSecretKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token.")
		}
		return next(c)
	}
}
