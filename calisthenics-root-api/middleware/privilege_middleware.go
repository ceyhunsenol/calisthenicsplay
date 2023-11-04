package middleware

import (
	"calisthenics-root-api/cache"
	"calisthenics-root-api/pkg"
	"calisthenics-root-api/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"regexp"
	"strings"
)

type PrivilegeMiddleware struct {
	userService          service.IUserService
	accessTokenSecretKey string
}

func NewPrivilegeMiddleware(userService service.IUserService) *PrivilegeMiddleware {
	return &PrivilegeMiddleware{
		userService:          userService,
		accessTokenSecretKey: viper.GetString("security.jwt.access-token-secret-key"),
	}
}

func (p *PrivilegeMiddleware) isAllowed(path string, allowedPatterns []string) bool {
	for _, pattern := range allowedPatterns {
		regexPattern := p.convertToRegex(pattern)
		matched, _ := regexp.MatchString(regexPattern, path)
		matchString, _ := regexp.MatchString(regexPattern, path+"/")
		if matched || matchString {
			return true
		}
	}
	return false
}

func (p *PrivilegeMiddleware) convertToRegex(pattern string) string {
	escaped := regexp.QuoteMeta(pattern)
	regexPattern := strings.Replace(escaped, "\\*\\*", ".*", -1)
	return "^" + regexPattern + "$"
}

func (p *PrivilegeMiddleware) MiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		context, _ := GetServiceContextByEchoContext(c)
		token, _ := pkg.ClearToken(context.Authorization)
		userID, err := pkg.GetUserIDFromToken(token, p.accessTokenSecretKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		user, err := p.userService.GetById(userID)
		if err != nil {
			return err
		}
		var endPoints []string
		for _, role := range user.Roles {
			for _, privilege := range role.Privileges {
				split := strings.Split(privilege.EndpointsJoin, ":")
				endPoints = append(endPoints, split...)
			}
		}
		path := c.Request().URL.Path
		allowed := p.isAllowed(path, endPoints)
		if !allowed {
			return echo.NewHTTPError(http.StatusForbidden, "Forbidden.")
		}
		cache.SetRequestCache(fmt.Sprintf("user_%s", userID), user, c)
		return next(c)
	}
}
