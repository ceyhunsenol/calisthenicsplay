package middleware

import (
	"calisthenics-root-api/model"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

func ServiceContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("X-Authorization")
		langCode := c.Request().Header.Get("X-Lang-Code")

		if langCode == "" {
			langCode = "en"
		}

		ctx := c.Request().Context()
		serviceCtx := &model.ServiceContext{
			Authorization: authorization,
			LangCode:      langCode,
		}

		ctx = context.WithValue(ctx, "ServiceContextKey", serviceCtx)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func GetServiceContextByEchoContext(c echo.Context) (*model.ServiceContext, error) {
	serviceCtx, ok := c.Request().Context().Value("ServiceContextKey").(*model.ServiceContext)
	if !ok {
		return nil, fmt.Errorf("ServiceContext not found")
	}
	return serviceCtx, nil
}
