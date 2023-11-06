package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type MediaController struct {
	mediaService                 service.IMediaService
	contentTranslationOperations service.IContentTranslationOperations
	DB                           *gorm.DB
}

func NewMediaController(contentTranslationOperations service.IContentTranslationOperations, mediaService service.IMediaService, DB *gorm.DB) *MediaController {
	return &MediaController{
		mediaService:                 mediaService,
		contentTranslationOperations: contentTranslationOperations,
		DB:                           DB,
	}
}

func (u *MediaController) InitMediaRoutes(e *echo.Group) {
	e.POST("", u.SaveMedia)
	e.PUT("/:id", u.UpdateMedia)
	e.GET("", u.GetMedias)
	e.GET("/:id", u.GetMedia)
	e.DELETE("/:id", u.DeleteMedia)
}

func (u *MediaController) SaveMedia(c echo.Context) error {
	mediaDTO := MediaDTO{}
	if err := c.Bind(&mediaDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&mediaDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	media := data.Media{
		DescriptionCode: mediaDTO.DescriptionCode,
		URL:             mediaDTO.URL,
		Type:            mediaDTO.Type,
		ContentID:       mediaDTO.ContentID,
		Active:          mediaDTO.Active,
	}

	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Media could not be saved."})
	}
	savedMedia, err := u.mediaService.Save(media)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Media could not be saved."})
	}
	request := model.ContentTranslationRequest{
		ContentID:    savedMedia.ID,
		Translations: make([]model.ContentTranslationModel, 0),
	}
	for _, translation := range mediaDTO.Translations {
		translationModel := model.ContentTranslationModel{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
		}
		request.Translations = append(request.Translations, translationModel)
	}
	serviceError := u.contentTranslationOperations.SaveContentTranslations(request)
	if serviceError != nil {
		tx.Rollback()
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	tx.Commit()
	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Created."})
}

func (u *MediaController) UpdateMedia(c echo.Context) error {
	mediaDTO := MediaDTO{}
	if err := c.Bind(&mediaDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&mediaDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id := c.Param("id")
	media, err := u.mediaService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Media not found."})
	}
	media.DescriptionCode = mediaDTO.DescriptionCode
	media.URL = mediaDTO.URL
	media.Type = mediaDTO.Type
	media.ContentID = mediaDTO.ContentID
	media.Active = mediaDTO.Active
	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Media could not be updated."})
	}
	_, err = u.mediaService.Update(*media)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Media could not be updated."})
	}
	request := model.ContentTranslationRequest{
		ContentID:    media.ID,
		Translations: make([]model.ContentTranslationModel, 0),
	}
	for _, translation := range mediaDTO.Translations {
		translationModel := model.ContentTranslationModel{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
		}
		request.Translations = append(request.Translations, translationModel)
	}
	serviceError := u.contentTranslationOperations.SaveContentTranslations(request)
	if serviceError != nil {
		tx.Rollback()
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &MessageResource{Code: http.StatusOK, Message: "Updated."})
}

func (u *MediaController) GetMedias(c echo.Context) error {
	medias, err := u.mediaService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Medias could not be got."})
	}

	mediaResources := make([]MediaResource, 0)
	for _, media := range medias {
		mediaResources = append(mediaResources, MediaResource{
			ID:              media.ID,
			DescriptionCode: media.DescriptionCode,
			URL:             media.URL,
			Type:            media.Type,
			ContentID:       media.ContentID,
			Active:          media.Active,
		})
	}

	return c.JSON(http.StatusOK, mediaResources)
}

func (u *MediaController) GetMedia(c echo.Context) error {
	id := c.Param("id")
	media, err := u.mediaService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Media not found."})
	}

	mediaResource := MediaResource{
		ID:              media.ID,
		DescriptionCode: media.DescriptionCode,
		URL:             media.URL,
		Type:            media.Type,
		ContentID:       media.ContentID,
		Active:          media.Active,
	}

	return c.JSON(http.StatusOK, mediaResource)
}

func (u *MediaController) DeleteMedia(c echo.Context) error {
	id := c.Param("id")
	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Media could not be deleted."})
	}
	err := u.mediaService.Delete(id)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Media could not be deleted."})
	}
	serviceError := u.contentTranslationOperations.DeleteAllContentTranslations(id)
	if serviceError != nil {
		tx.Rollback()
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	tx.Commit()
	return c.JSON(http.StatusNoContent, nil)
}
