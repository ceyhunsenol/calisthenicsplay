package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type ContentController struct {
	contentService               service.IContentService
	helperContentOperations      service.IHelperContentOperations
	requirementContentOperations service.IRequirementContentOperations
	contentTranslationOperations service.IContentTranslationOperations
	DB                           *gorm.DB
}

func NewContentController(contentService service.IContentService,
	helperContentOperations service.IHelperContentOperations,
	requirementContentOperations service.IRequirementContentOperations,
	contentTranslationOperations service.IContentTranslationOperations,
	DB *gorm.DB,
) *ContentController {
	return &ContentController{
		contentService:               contentService,
		helperContentOperations:      helperContentOperations,
		requirementContentOperations: requirementContentOperations,
		contentTranslationOperations: contentTranslationOperations,
		DB:                           DB,
	}
}

func (u *ContentController) InitContentRoutes(e *echo.Group) {
	e.POST("", u.SaveContent)
	e.PUT("/:id", u.UpdateContent)
	e.GET("", u.GetContents)
	e.GET("/:id", u.GetContent)
	e.DELETE("/:id", u.DeleteContent)

	e.POST("/:id/helpers", u.AddHelperContent)
	e.DELETE("/:id/helpers/:helperID", u.RemoveHelperContent)

	e.POST("/:id/requirements", u.AddRequirementContent)
	e.DELETE("/:id/requirements/:requirementID", u.RemoveRequirementContent)
}

func (u *ContentController) SaveContent(c echo.Context) error {
	contentDTO := ContentDTO{}
	if err := c.Bind(&contentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&contentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	exists, err := u.contentService.ExistsByCode(contentDTO.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Content already exists in this code."})
	}

	content := data.Content{
		Code:            contentDTO.Code,
		DescriptionCode: contentDTO.Description,
		Active:          contentDTO.Active,
	}

	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content could not be saved."})
	}
	_, err = u.contentService.Save(content)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content could not be saved."})
	}
	requests := make([]model.ContentTranslationRequest, 0)
	for _, translation := range contentDTO.Translations {
		request := model.ContentTranslationRequest{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
			ContentID: content.ID,
		}
		requests = append(requests, request)
	}
	serviceError := u.contentTranslationOperations.SaveContentTranslations(requests)
	if serviceError != nil {
		tx.Rollback()
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	tx.Commit()
	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Created."})
}

func (u *ContentController) UpdateContent(c echo.Context) error {
	contentDTO := ContentDTO{}
	if err := c.Bind(&contentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&contentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id := c.Param("id")
	exists, err := u.contentService.GetByCode(contentDTO.Code)
	if err == nil && exists.ID != id {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Content already exists in this code."})
	}
	content, err := u.contentService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Content not found."})
	}

	content.Code = contentDTO.Code
	content.DescriptionCode = contentDTO.Description
	content.Active = contentDTO.Active

	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content could not be updated."})
	}
	_, err = u.contentService.Update(*content)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content could not be updated."})
	}
	requests := make([]model.ContentTranslationRequest, 0)
	for _, translation := range contentDTO.Translations {
		request := model.ContentTranslationRequest{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
			ContentID: content.ID,
		}
		requests = append(requests, request)
	}
	serviceError := u.contentTranslationOperations.SaveContentTranslations(requests)
	if serviceError != nil {
		tx.Rollback()
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &MessageResource{Code: http.StatusOK, Message: "Updated."})
}

func (u *ContentController) GetContents(c echo.Context) error {
	contents, err := u.contentService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Contents could not be got."})
	}

	contentResources := make([]ContentResource, 0)
	for _, content := range contents {
		mediaResources := make([]ContentMediaResource, 0)
		for _, media := range content.Medias {
			mediaResources = append(mediaResources, ContentMediaResource{
				ID:              media.ID,
				DescriptionCode: media.DescriptionCode,
				URL:             media.URL,
				Type:            media.Type,
			})
		}
		contentResources = append(contentResources, ContentResource{
			ID:          content.ID,
			Code:        content.Code,
			Description: content.DescriptionCode,
			Active:      content.Active,
			Medias:      mediaResources,
		})
	}

	return c.JSON(http.StatusOK, contentResources)
}

func (u *ContentController) GetContent(c echo.Context) error {
	id := c.Param("id")
	content, err := u.contentService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Content not found."})
	}

	mediaResources := make([]ContentMediaResource, 0)
	for _, media := range content.Medias {
		mediaResources = append(mediaResources, ContentMediaResource{
			ID:              media.ID,
			DescriptionCode: media.DescriptionCode,
			URL:             media.URL,
			Type:            media.Type,
		})
	}
	contentResource := ContentResource{
		ID:          content.ID,
		Code:        content.Code,
		Description: content.DescriptionCode,
		Active:      content.Active,
		Medias:      mediaResources,
	}

	return c.JSON(http.StatusOK, contentResource)
}

func (u *ContentController) DeleteContent(c echo.Context) error {
	id := c.Param("id")
	err := u.contentService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content could not be deleted."})
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (u *ContentController) AddHelperContent(c echo.Context) error {
	helperContentDTO := HelperContentDTO{}
	if err := c.Bind(&helperContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&helperContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id := c.Param("id")
	request := model.HelperContentRequest{
		ContentID:       id,
		HelperContentID: helperContentDTO.HelperContentID,
	}
	err := u.helperContentOperations.AddHelperContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}

	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Helper content added."})
}

func (u *ContentController) RemoveHelperContent(c echo.Context) error {
	id := c.Param("id")
	helperID := c.Param("helperID")
	request := model.HelperContentRequest{
		ContentID:       id,
		HelperContentID: helperID,
	}
	err := u.helperContentOperations.RemoveHelperContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (u *ContentController) AddRequirementContent(c echo.Context) error {
	requirementContentDTO := RequirementContentDTO{}
	if err := c.Bind(&requirementContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&requirementContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id := c.Param("id")
	request := model.RequirementContentRequest{
		ContentID:            id,
		RequirementContentID: requirementContentDTO.RequirementContentID,
	}
	err := u.requirementContentOperations.AddRequirementContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}

	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Requirement content added."})
}

func (u *ContentController) RemoveRequirementContent(c echo.Context) error {
	id := c.Param("id")
	requirementID := c.Param("requirementID")
	request := model.RequirementContentRequest{
		ContentID:            id,
		RequirementContentID: requirementID,
	}
	err := u.requirementContentOperations.RemoveRequirementContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}

	return c.JSON(http.StatusNoContent, nil)
}
