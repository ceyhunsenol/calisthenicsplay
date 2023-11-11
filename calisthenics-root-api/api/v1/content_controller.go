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
	cacheRequestService          service.ICacheRequestService
}

func NewContentController(contentService service.IContentService,
	helperContentOperations service.IHelperContentOperations,
	requirementContentOperations service.IRequirementContentOperations,
	contentTranslationOperations service.IContentTranslationOperations,
	DB *gorm.DB,
	cacheRequestService service.ICacheRequestService,
) *ContentController {
	return &ContentController{
		contentService:               contentService,
		helperContentOperations:      helperContentOperations,
		requirementContentOperations: requirementContentOperations,
		contentTranslationOperations: contentTranslationOperations,
		DB:                           DB,
		cacheRequestService:          cacheRequestService,
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
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&contentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	exists, err := u.contentService.ExistsByCode(contentDTO.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Content could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Content already exists in this code."})
	}

	content := data.Content{
		Code:            contentDTO.Code,
		DescriptionCode: contentDTO.Description,
		Active:          contentDTO.Active,
	}

	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Content could not be saved."})
	}
	_, err = u.contentService.Save(tx, content)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Content could not be saved."})
	}
	request := model.ContentTranslationRequest{
		ContentID:    content.ID,
		Translations: make([]model.ContentTranslationModel, 0),
	}
	for _, translation := range contentDTO.Translations {
		translationModel := model.ContentTranslationModel{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
		}
		request.Translations = append(request.Translations, translationModel)
	}
	serviceError := u.contentTranslationOperations.SaveContentTranslations(tx, request)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError = u.cacheRequestService.ContentRefreshRequest(content.ID)
	if serviceError != nil && serviceError.Message != "Request error" {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	tx.Commit()
	return c.JSON(http.StatusCreated, &MessageResource{Message: "Created."})
}

func (u *ContentController) UpdateContent(c echo.Context) error {
	contentDTO := ContentDTO{}
	if err := c.Bind(&contentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&contentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	id := c.Param("id")
	exists, err := u.contentService.GetByCode(contentDTO.Code)
	if err == nil && exists.ID != id {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Content already exists in this code."})
	}
	content, err := u.contentService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Content not found."})
	}

	content.Code = contentDTO.Code
	content.DescriptionCode = contentDTO.Description
	content.Active = contentDTO.Active

	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Content could not be updated."})
	}
	_, err = u.contentService.Update(tx, *content)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Content could not be updated."})
	}
	request := model.ContentTranslationRequest{
		ContentID:    content.ID,
		Translations: make([]model.ContentTranslationModel, 0),
	}
	for _, translation := range contentDTO.Translations {
		translationModel := model.ContentTranslationModel{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
		}
		request.Translations = append(request.Translations, translationModel)
	}
	serviceError := u.contentTranslationOperations.SaveContentTranslations(tx, request)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	// content apiye cache icin request atiliyor
	if !content.Active {
		serviceError = u.cacheRequestService.ContentWithMediasRefreshRequest(content.ID)
	} else {
		serviceError = u.cacheRequestService.ContentRefreshRequest(content.ID)
	}
	if serviceError != nil && serviceError.Message != "Request error" {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	tx.Commit()
	return c.JSON(http.StatusOK, &MessageResource{Message: "Updated."})
}

func (u *ContentController) GetContents(c echo.Context) error {
	contents, err := u.contentService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Contents could not be got."})
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
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Content not found."})
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
	tx := u.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Content could not be deleted."})
	}
	err := u.contentService.Delete(tx, id)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Content could not be deleted."})
	}
	serviceError := u.contentTranslationOperations.DeleteAllContentTranslations(tx, id)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError = u.cacheRequestService.ContentWithMediasRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	tx.Commit()
	return c.JSON(http.StatusNoContent, nil)
}

func (u *ContentController) AddHelperContent(c echo.Context) error {
	helperContentDTO := HelperContentDTO{}
	if err := c.Bind(&helperContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&helperContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	id := c.Param("id")
	request := model.HelperContentRequest{
		ContentID:       id,
		HelperContentID: helperContentDTO.HelperContentID,
	}
	err := u.helperContentOperations.AddHelperContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError := u.cacheRequestService.ContentRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	return c.JSON(http.StatusCreated, &MessageResource{Message: "Helper content added."})
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
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError := u.cacheRequestService.ContentRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (u *ContentController) AddRequirementContent(c echo.Context) error {
	requirementContentDTO := RequirementContentDTO{}
	if err := c.Bind(&requirementContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&requirementContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	id := c.Param("id")
	request := model.RequirementContentRequest{
		ContentID:            id,
		RequirementContentID: requirementContentDTO.RequirementContentID,
	}
	err := u.requirementContentOperations.AddRequirementContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError := u.cacheRequestService.ContentRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	return c.JSON(http.StatusCreated, &MessageResource{Message: "Requirement content added."})
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
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError := u.cacheRequestService.ContentRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	return c.JSON(http.StatusNoContent, nil)
}
