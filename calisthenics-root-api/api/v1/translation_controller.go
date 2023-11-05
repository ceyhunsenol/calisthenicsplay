package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TranslationController struct {
	translationService service.ITranslationService
}

func NewTranslationController(translationService service.ITranslationService) *TranslationController {
	return &TranslationController{translationService: translationService}
}

func (t *TranslationController) InitTranslationRoutes(e *echo.Group) {
	e.POST("", t.SaveTranslation)
	e.PUT("/:id", t.UpdateTranslation)
	e.GET("", t.GetTranslations)
	e.GET("/:id", t.GetTranslation)
	e.DELETE("/:id", t.DeleteTranslation)
}

func (t *TranslationController) SaveTranslation(c echo.Context) error {
	translationDTO := TranslationDTO{}
	if err := c.Bind(&translationDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&translationDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	exists, err := t.translationService.ExistsByCodeAndLangCode(translationDTO.Code, translationDTO.LangCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Translation could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Translation already exists with this code and language code."})
	}

	translation := data.Translation{
		Code:      translationDTO.Code,
		LangCode:  translationDTO.LangCode,
		Translate: translationDTO.Translate,
		Active:    translationDTO.Active,
		Domain:    translationDTO.Domain,
	}

	_, err = t.translationService.Save(translation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusBadRequest, Message: "Translation could not be saved."})
	}
	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Created."})
}

func (t *TranslationController) UpdateTranslation(c echo.Context) error {
	translationDTO := TranslationDTO{}
	if err := c.Bind(&translationDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&translationDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id := c.Param("id")
	exists, err := t.translationService.GetByCodeAndLangCode(translationDTO.Code, translationDTO.LangCode)
	if err == nil && exists.ID != id {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Translation already exists with this code and language code."})
	}
	translation, err := t.translationService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Translation not found."})
	}
	translation.Code = translationDTO.Code
	translation.LangCode = translationDTO.LangCode
	translation.Translate = translationDTO.Translate
	translation.Active = translationDTO.Active
	translation.Domain = translationDTO.Domain
	_, err = t.translationService.Update(*translation)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &MessageResource{Code: http.StatusOK, Message: "Updated."})
}

func (t *TranslationController) GetTranslations(c echo.Context) error {
	translations, err := t.translationService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Translations could not be retrieved."})
	}

	translationResources := make([]TranslationResource, 0)
	for _, translation := range translations {
		translationResources = append(translationResources, TranslationResource{
			ID:        translation.ID,
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
			Domain:    translation.Domain,
		})
	}

	return c.JSON(http.StatusOK, translationResources)
}

func (t *TranslationController) GetTranslation(c echo.Context) error {
	id := c.Param("id")
	translation, err := t.translationService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Translation not found."})
	}

	translationResource := TranslationResource{
		ID:        translation.ID,
		Code:      translation.Code,
		LangCode:  translation.LangCode,
		Translate: translation.Translate,
		Active:    translation.Active,
		Domain:    translation.Domain,
	}

	return c.JSON(http.StatusOK, translationResource)
}

func (t *TranslationController) DeleteTranslation(c echo.Context) error {
	id := c.Param("id")
	err := t.translationService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Translation could not be deleted."})
	}
	return c.JSON(http.StatusNoContent, nil)
}
