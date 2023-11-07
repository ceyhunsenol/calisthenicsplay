package api

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/utils"
	"net/http"
)

type CacheController struct {
	mediaOperations   service.IMediaOperations
	contentOperations service.IContentOperations
	genreOperations   service.IGenreOperations
}

func NewCacheController(mediaOperations service.IMediaOperations,
	contentOperations service.IContentOperations,
	genreOperations service.IGenreOperations,
) *CacheController {
	return &CacheController{
		mediaOperations:   mediaOperations,
		contentOperations: contentOperations,
		genreOperations:   genreOperations,
	}
}

func (u *CacheController) InitCacheRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/cache")

	v1.GET("/refresh-all", u.RefreshAll)
}

func (u *CacheController) RefreshAll(c echo.Context) error {
	cacheTypes := c.QueryParams()["cacheType"]
	if utils.Contains(cacheTypes, string(cache.Genre)) {
		err := u.genreOperations.SaveCacheGenres()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Unknown error"})
		}
	}
	if utils.Contains(cacheTypes, string(cache.Content)) {
		err := u.contentOperations.SaveCacheContents()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Unknown error"})
		}
	}
	if utils.Contains(cacheTypes, string(cache.Media)) {
		err := u.mediaOperations.SaveCacheMedias()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Unknown error"})
		}
	}
	return c.JSON(http.StatusOK, &MessageResource{Code: http.StatusOK, Message: "Cached all contents"})
}
