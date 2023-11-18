package api

import (
	"calisthenics-content-api/middleware"
	"calisthenics-content-api/model"
	"calisthenics-content-api/service"
	"encoding/hex"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HLSMediaController struct {
	hlsMediaService service.IHLSMediaService
}

func NewHLSMediaController(hlsMediaService service.IHLSMediaService) *HLSMediaController {
	return &HLSMediaController{
		hlsMediaService: hlsMediaService,
	}
}

func (u *HLSMediaController) InitHLSMediaRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/hls-media")
	v1.GET("/master.m3u8", u.GetMasterPlaylist)
	v1.GET("/:id", u.GetMediaPlaylist)
	v1.GET("/license.key", u.GetLicenseKey)
	v1.GET("/media-url/:mediaID", u.GetMediaURL)
}

func (u *HLSMediaController) GetMasterPlaylist(c echo.Context) error {
	token := c.QueryParam("t")
	if token == "" {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Token required"})
	}
	playlist, err := u.hlsMediaService.GetMasterPlaylist(token)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	return c.Blob(http.StatusOK, "application/vnd.apple.mpegurl", []byte(playlist))
}

func (u *HLSMediaController) GetMediaPlaylist(c echo.Context) error {
	id := c.Param("id")
	token := c.QueryParam("t")
	request := model.VideoPlaylistRequest{
		Resolution: id,
		Token:      token,
	}
	playlist, err := u.hlsMediaService.GetMediaPlaylist(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	return c.Blob(http.StatusOK, "application/vnd.apple.mpegurl", []byte(playlist))
}

func (u *HLSMediaController) GetLicenseKey(c echo.Context) error {
	token := c.QueryParam("t")
	licenseKey, serviceError := u.hlsMediaService.GetLicenseKey(token)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	licenseKeyBytes, err := hex.DecodeString(licenseKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Unknown error"})
	}

	return c.Blob(http.StatusOK, "application/vnd.apple.mpegurl", licenseKeyBytes)
}

func (u *HLSMediaController) GetMediaURL(c echo.Context) error {
	mediaID := c.Param("mediaID")
	context, _ := middleware.GetServiceContextByEchoContext(c)
	request := model.VideoURLRequest{
		MediaID:      mediaID,
		Token:        context.Authorization,
		UserAgent:    context.UserAgent,
		Host:         context.Host,
		CallerIP:     context.CallerIP,
		PlatformType: context.PlatformType,
		LangCode:     context.LangCode,
	}
	url, err := u.hlsMediaService.GetMediaURL(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}

	resolutions := make([]VideoURLResolutionResource, 0)
	for _, resolution := range url.Resolutions {
		resolutionModel := VideoURLResolutionResource{
			Height:       resolution.Height,
			Bandwidth:    resolution.Bandwidth,
			AvgBandwidth: resolution.AvgBandwidth,
		}
		resolutions = append(resolutions, resolutionModel)
	}

	resource := VideoURLResource{
		VideoURL:    url.VideoURL,
		Resolutions: resolutions,
	}
	return c.JSON(http.StatusOK, resource)
}
