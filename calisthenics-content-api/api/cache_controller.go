package api

import (
	"calisthenics-content-api/cache"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/utils"
)

type CacheController struct {
	contentCacheService cache.IContentCacheService
}

func NewCacheController(contentCacheService cache.IContentCacheService) *CacheController {
	return &CacheController{
		contentCacheService: contentCacheService,
	}
}

func (u *CacheController) InitCacheRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/cache")

	v1.GET("/refresh-all", u.RefreshAll)
}

func (u *CacheController) RefreshAll(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	cacheTypes := c.QueryParams()["cacheType"]
	if utils.Contains(cacheTypes, string(cache.Genre)) {
	}
	if utils.Contains(cacheTypes, string(cache.Content)) {
	}
	if utils.Contains(cacheTypes, string(cache.Media)) {
	}
	return nil
}
