package middleware

import (
	"calisthenics-content-api/model"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"strings"
)

func ServiceContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		authorization := req.Header.Get("X-Authorization")
		langCode := req.Header.Get("X-Lang-Code")
		platformType := req.Header.Get("X-Platform-Type")
		userAgent := req.Header.Get("x-user-agent")
		clientIP := c.RealIP()

		scheme := "http"
		if req.TLS != nil {
			scheme = "https"
		}
		host := req.Host

		// Ana URL'yi oluştur
		baseURL := fmt.Sprintf("%s://%s", scheme, host)

		if userAgent == "" {
			userAgent = c.Request().Header.Get("user-agent")
		}

		if langCode == "" {
			langCode = "en"
		}

		ctx := c.Request().Context()
		serviceCtx := &model.ServiceContext{
			Authorization: authorization,
			LangCode:      langCode,
			PlatformType:  platformType,
			ClientIP:      clientIP,
			CallerIP:      clientIP,
			ClientIPList:  strings.Split(clientIP, ","),
			UserAgent:     userAgent,
			Host:          baseURL,
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
