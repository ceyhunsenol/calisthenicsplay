package api

import (
	"calisthenics-content-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CacheController struct {
	mediaCacheOperations         service.IMediaCacheOperations
	contentCacheOperations       service.IContentCacheOperations
	genreCacheOperations         service.IGenreCacheOperations
	generalInfoCacheOperations   service.IGeneralInfoCacheOperations
	contentAccessCacheOperations service.IContentAccessCacheOperations
	mediaAccessCacheOperations   service.IMediaAccessCacheOperations
	hlsEncodingCacheOperations   service.IHLSEncodingCacheOperations
	initService                  service.IInitService
}

func NewCacheController(mediaCacheOperations service.IMediaCacheOperations,
	contentCacheOperations service.IContentCacheOperations,
	genreCacheOperations service.IGenreCacheOperations,
	generalInfoCacheOperations service.IGeneralInfoCacheOperations,
	contentAccessCacheOperations service.IContentAccessCacheOperations,
	mediaAccessCacheOperations service.IMediaAccessCacheOperations,
	hlsEncodingCacheOperations service.IHLSEncodingCacheOperations,
	initService service.IInitService,
) *CacheController {
	return &CacheController{
		mediaCacheOperations:         mediaCacheOperations,
		contentCacheOperations:       contentCacheOperations,
		genreCacheOperations:         genreCacheOperations,
		generalInfoCacheOperations:   generalInfoCacheOperations,
		contentAccessCacheOperations: contentAccessCacheOperations,
		mediaAccessCacheOperations:   mediaAccessCacheOperations,
		hlsEncodingCacheOperations:   hlsEncodingCacheOperations,
		initService:                  initService,
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

	for _, cacheType := range cacheTypes {
		serviceError := u.initService.CallFromCacheFuncAllByCacheType(cacheType)
		if serviceError != nil {
			return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
		}
	}

	return c.JSON(http.StatusOK, &MessageResource{Message: "Cached all contents"})
}

func (u *CacheController) Refresh(c echo.Context) error {
	cacheType := c.Param("cacheType")
	id := c.Param("id")
	return c.JSON(http.StatusOK, u.initService.CallFromCacheFuncByCacheType(cacheType, id))
}

func (u *CacheController) RefreshContentWithMedias(c echo.Context) error {
	id := c.Param("id")
	serviceError := u.contentCacheOperations.SaveCacheContentWithMedias(id)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	return c.JSON(http.StatusOK, &MessageResource{Message: "Cached"})
}
