package api

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/utils"
	"net/http"
)

type CacheController struct {
	mediaCacheOperations   service.IMediaCacheOperations
	contentCacheOperations service.IContentCacheOperations
	genreCacheOperations   service.IGenreCacheOperations
}

func NewCacheController(mediaCacheOperations service.IMediaCacheOperations,
	contentCacheOperations service.IContentCacheOperations,
	genreCacheOperations service.IGenreCacheOperations,
) *CacheController {
	return &CacheController{
		mediaCacheOperations:   mediaCacheOperations,
		contentCacheOperations: contentCacheOperations,
		genreCacheOperations:   genreCacheOperations,
	}
}

func (u *CacheController) InitCacheRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/cache")

	v1.GET("/refresh-all", u.RefreshAll)
	v1.GET("/refresh/:cacheType/:id", u.Refresh)
}

func (u *CacheController) RefreshAll(c echo.Context) error {
	cacheTypes := c.QueryParams()["cacheType"]
	if utils.Contains(cacheTypes, string(cache.Genre)) {
		serviceError := u.genreCacheOperations.SaveCacheGenres()
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
	}
	if utils.Contains(cacheTypes, string(cache.Content)) {
		serviceError := u.contentCacheOperations.SaveCacheContents()
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
	}
	if utils.Contains(cacheTypes, string(cache.Media)) {
		serviceError := u.mediaCacheOperations.SaveCacheMedias()
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
	}
	return c.JSON(http.StatusOK, &MessageResource{Message: "Cached all contents"})
}

func (u *CacheController) Refresh(c echo.Context) error {
	cacheType := c.Param("cacheType")
	id := c.Param("id")

	switch cacheType {
	case string(cache.Genre):
		cac, serviceError := u.genreCacheOperations.SaveCacheGenre(id)
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
		return c.JSON(http.StatusOK, cac)

	case string(cache.Content):
		cac, serviceError := u.contentCacheOperations.SaveCacheContent(id)
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
		return c.JSON(http.StatusOK, cac)

	case string(cache.Media):
		cac, serviceError := u.mediaCacheOperations.SaveCacheMedia(id)
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
		return c.JSON(http.StatusOK, cac)

	default:
		return c.JSON(http.StatusNotImplemented, &MessageResource{Message: "Not implemented"})
	}
}
