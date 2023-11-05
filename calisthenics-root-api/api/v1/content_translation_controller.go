package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContentTranslationController struct {
	contentTranslationService service.IContentTranslationService
}

func NewContentTranslationController(contentTranslationService service.IContentTranslationService) *ContentTranslationController {
	return &ContentTranslationController{contentTranslationService: contentTranslationService}
}

func (c *ContentTranslationController) InitContentTranslationRoutes(e *echo.Group) {
	e.POST("", c.SaveContentTranslation)
	e.PUT("/:id", c.UpdateContentTranslation)
	e.GET("", c.GetContentTranslations)
	e.GET("/:id", c.GetContentTranslation)
	e.DELETE("/:id", c.DeleteContentTranslation)
}

func (c *ContentTranslationController) SaveContentTranslation(ctx echo.Context) error {
	contentTranslationDTO := ContentTranslationDTO{}
	if err := ctx.Bind(&contentTranslationDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := ctx.Validate(&contentTranslationDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	exists, err := c.contentTranslationService.ExistsByCodeAndLangCode(contentTranslationDTO.Code, contentTranslationDTO.LangCode)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content translation could not be saved."})
	}
	if exists {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Content translation already exists with this code and language code."})
	}

	contentTranslation := data.ContentTranslation{
		Code:      contentTranslationDTO.Code,
		LangCode:  contentTranslationDTO.LangCode,
		Translate: contentTranslationDTO.Translate,
		Active:    contentTranslationDTO.Active,
		ContentID: contentTranslationDTO.ContentID,
	}

	_, err = c.contentTranslationService.Save(contentTranslation)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusBadRequest, Message: "Content translation could not be saved."})
	}
	return ctx.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Created."})
}

func (c *ContentTranslationController) UpdateContentTranslation(ctx echo.Context) error {
	contentTranslationDTO := ContentTranslationDTO{}
	if err := ctx.Bind(&contentTranslationDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := ctx.Validate(&contentTranslationDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id := ctx.Param("id")
	exists, err := c.contentTranslationService.GetByCodeAndLangCode(contentTranslationDTO.Code, contentTranslationDTO.LangCode)
	if err == nil && exists.ID != id {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Content translation already exists with this code and language code."})
	}
	contentTranslation, err := c.contentTranslationService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Content translation not found."})
	}
	contentTranslation.Code = contentTranslationDTO.Code
	contentTranslation.LangCode = contentTranslationDTO.LangCode
	contentTranslation.Translate = contentTranslationDTO.Translate
	contentTranslation.Active = contentTranslationDTO.Active
	contentTranslation.ContentID = contentTranslationDTO.ContentID
	_, err = c.contentTranslationService.Update(*contentTranslation)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Translation content could not be updated."})
	}
	return ctx.JSON(http.StatusOK, &MessageResource{Code: http.StatusOK, Message: "Updated."})
}

func (c *ContentTranslationController) GetContentTranslations(ctx echo.Context) error {
	contentTranslations, err := c.contentTranslationService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content translations could not be retrieved."})
	}

	contentTranslationResources := make([]ContentTranslationResource, 0)
	for _, contentTranslation := range contentTranslations {
		contentTranslationResources = append(contentTranslationResources, ContentTranslationResource{
			ID:        contentTranslation.ID,
			Code:      contentTranslation.Code,
			LangCode:  contentTranslation.LangCode,
			Translate: contentTranslation.Translate,
			Active:    contentTranslation.Active,
			ContentID: contentTranslation.ContentID,
		})
	}

	return ctx.JSON(http.StatusOK, contentTranslationResources)
}

func (c *ContentTranslationController) GetContentTranslation(ctx echo.Context) error {
	id := ctx.Param("id")
	contentTranslation, err := c.contentTranslationService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Content translation not found."})
	}

	contentTranslationResource := ContentTranslationResource{
		ID:        contentTranslation.ID,
		Code:      contentTranslation.Code,
		LangCode:  contentTranslation.LangCode,
		Translate: contentTranslation.Translate,
		Active:    contentTranslation.Active,
		ContentID: contentTranslation.ContentID,
	}

	return ctx.JSON(http.StatusOK, contentTranslationResource)
}

func (c *ContentTranslationController) DeleteContentTranslation(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.contentTranslationService.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Content translation could not be deleted."})
	}
	return ctx.JSON(http.StatusNoContent, nil)
}
