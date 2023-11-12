package api

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/utils"
	"net/http"
)

type CacheController struct {
	mediaCacheOperations         service.IMediaCacheOperations
	contentCacheOperations       service.IContentCacheOperations
	genreCacheOperations         service.IGenreCacheOperations
	generalInfoCacheOperations   service.IGeneralInfoCacheOperations
	contentAccessCacheOperations service.IContentAccessCacheOperations
	mediaAccessCacheOperations   service.IMediaAccessCacheOperations
}

func NewCacheController(mediaCacheOperations service.IMediaCacheOperations,
	contentCacheOperations service.IContentCacheOperations,
	genreCacheOperations service.IGenreCacheOperations,
	generalInfoCacheOperations service.IGeneralInfoCacheOperations,
	contentAccessCacheOperations service.IContentAccessCacheOperations,
	mediaAccessCacheOperations service.IMediaAccessCacheOperations,
) *CacheController {
	return &CacheController{
		mediaCacheOperations:         mediaCacheOperations,
		contentCacheOperations:       contentCacheOperations,
		genreCacheOperations:         genreCacheOperations,
		generalInfoCacheOperations:   generalInfoCacheOperations,
		contentAccessCacheOperations: contentAccessCacheOperations,
		mediaAccessCacheOperations:   mediaAccessCacheOperations,
	}
}

func (u *CacheController) InitCacheRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/cache")

	v1.GET("/refresh-all", u.RefreshAll)
	v1.GET("/refresh/:cacheType/:id", u.Refresh)
	v1.GET("/refresh/content-with-medias/:id", u.RefreshContentWithMedias)
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
	if utils.Contains(cacheTypes, string(cache.GeneralInfo)) {
		serviceError := u.generalInfoCacheOperations.SaveCacheGeneralInfos()
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
	}
	if utils.Contains(cacheTypes, string(cache.ContentAccess)) {
		serviceError := u.contentAccessCacheOperations.SaveCacheContentAccessList()
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
	}
	if utils.Contains(cacheTypes, string(cache.MediaAccess)) {
		serviceError := u.mediaAccessCacheOperations.SaveCacheMediaAccessList()
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
		cac := u.genreCacheOperations.SaveCacheGenre(id)
		return c.JSON(http.StatusOK, cac)

	case string(cache.Content):
		cac := u.contentCacheOperations.SaveCacheContent(id)
		return c.JSON(http.StatusOK, cac)

	case string(cache.Media):
		cac := u.mediaCacheOperations.SaveCacheMedia(id)
		return c.JSON(http.StatusOK, cac)

	case string(cache.GeneralInfo):
		cac := u.generalInfoCacheOperations.SaveCacheGeneralInfo(id)
		return c.JSON(http.StatusOK, cac)

	case string(cache.ContentAccess):
		cac := u.contentAccessCacheOperations.SaveCacheContentAccess(id)
		return c.JSON(http.StatusOK, cac)

	case string(cache.MediaAccess):
		cac := u.mediaAccessCacheOperations.SaveCacheMediaAccess(id)
		return c.JSON(http.StatusOK, cac)

	default:
		return c.JSON(http.StatusNotImplemented, &MessageResource{Message: "Not implemented"})
	}
}

func (u *CacheController) RefreshContentWithMedias(c echo.Context) error {
	id := c.Param("id")
	serviceError := u.contentCacheOperations.SaveCacheContentWithMedias(id)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	return c.JSON(http.StatusOK, &MessageResource{Message: "Cached"})
}
