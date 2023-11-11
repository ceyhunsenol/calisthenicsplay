package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type GenreController struct {
	genreService                 service.IGenreService
	genreContentOperations       service.IGenreContentOperations
	contentTranslationOperations service.IContentTranslationOperations
	DB                           *gorm.DB
	cacheRequestService          service.ICacheRequestService
}

func NewGenreController(genreContentOperations service.IGenreContentOperations,
	genreService service.IGenreService,
	contentTranslationOperations service.IContentTranslationOperations,
	DB *gorm.DB,
	cacheRequestService service.ICacheRequestService,
) *GenreController {
	return &GenreController{
		genreService:                 genreService,
		genreContentOperations:       genreContentOperations,
		contentTranslationOperations: contentTranslationOperations,
		DB:                           DB,
		cacheRequestService:          cacheRequestService,
	}
}

func (g *GenreController) InitGenreRoutes(e *echo.Group) {
	e.POST("", g.SaveGenre)
	e.PUT("/:id", g.UpdateGenre)
	e.GET("", g.GetGenres)
	e.GET("/:id", g.GetGenre)
	e.DELETE("/:id", g.DeleteGenre)

	e.POST("/:id/contents", g.AddContent)
	e.DELETE("/:id/contents/:contentID", g.RemoveContent)
}

func (g *GenreController) SaveGenre(c echo.Context) error {
	genreDTO := GenreDTO{}
	if err := c.Bind(&genreDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&genreDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	exists, err := g.genreService.ExistsByCode(genreDTO.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genre could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Genre already exists in this code."})
	}

	genre := data.Genre{
		Type:        genreDTO.Type,
		Code:        genreDTO.Code,
		Description: genreDTO.Description,
		Section:     genreDTO.Section,
		Active:      genreDTO.Active,
	}

	tx := g.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genre could not be saved."})
	}
	savedGenre, err := g.genreService.Save(tx, genre)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genre could not be saved."})
	}
	request := model.ContentTranslationRequest{
		ContentID:    savedGenre.ID,
		Translations: make([]model.ContentTranslationModel, 0),
	}
	for _, translation := range genreDTO.Translations {
		translationModel := model.ContentTranslationModel{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
		}
		request.Translations = append(request.Translations, translationModel)
	}
	serviceError := g.contentTranslationOperations.SaveContentTranslations(tx, request)
	if err != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError = g.cacheRequestService.GenreRefreshRequest(savedGenre.ID)
	if serviceError != nil && serviceError.Message != "Request error" {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	tx.Commit()
	return c.JSON(http.StatusCreated, &MessageResource{Message: "Created."})
}

func (g *GenreController) UpdateGenre(c echo.Context) error {
	genreDTO := GenreDTO{}
	if err := c.Bind(&genreDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&genreDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	id := c.Param("id")
	exists, err := g.genreService.GetByCode(genreDTO.Code)
	if err == nil && exists.ID != id {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Genre already exists in this code."})
	}
	genre, err := g.genreService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Genre not found."})
	}
	genre.Type = genreDTO.Type
	genre.Code = genreDTO.Code
	genre.Description = genreDTO.Description
	genre.Section = genreDTO.Section
	genre.Active = genreDTO.Active
	tx := g.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genre could not be saved."})
	}
	_, err = g.genreService.Update(tx, *genre)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genre could not be updated."})
	}
	request := model.ContentTranslationRequest{
		ContentID:    genre.ID,
		Translations: make([]model.ContentTranslationModel, 0),
	}
	for _, translation := range genreDTO.Translations {
		translationModel := model.ContentTranslationModel{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
		}
		request.Translations = append(request.Translations, translationModel)
	}
	serviceError := g.contentTranslationOperations.SaveContentTranslations(tx, request)
	if err != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError = g.cacheRequestService.GenreRefreshRequest(genre.ID)
	if serviceError != nil && serviceError.Message != "Request error" {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	tx.Commit()
	return c.JSON(http.StatusOK, &MessageResource{Message: "Updated."})
}

func (g *GenreController) GetGenres(c echo.Context) error {
	genres, err := g.genreService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genres could not be got."})
	}

	genreResources := make([]GenreResource, 0)
	for _, genre := range genres {
		genreResources = append(genreResources, GenreResource{
			ID:          genre.ID,
			Type:        genre.Type,
			Code:        genre.Code,
			Description: genre.Description,
			Section:     genre.Section,
			Active:      genre.Active,
		})
	}

	return c.JSON(http.StatusOK, genreResources)
}

func (g *GenreController) GetGenre(c echo.Context) error {
	id := c.Param("id")
	genre, err := g.genreService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Genre not found."})
	}

	genreResource := GenreResource{
		ID:          genre.ID,
		Type:        genre.Type,
		Code:        genre.Code,
		Description: genre.Description,
		Section:     genre.Section,
		Active:      genre.Active,
	}

	return c.JSON(http.StatusOK, genreResource)
}

func (g *GenreController) DeleteGenre(c echo.Context) error {
	id := c.Param("id")
	tx := g.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genre could not be deleted."})
	}
	err := g.genreService.Delete(tx, id)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Genre could not be deleted."})
	}
	serviceError := g.contentTranslationOperations.DeleteAllContentTranslations(tx, id)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError = g.cacheRequestService.GenreRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	tx.Commit()
	return c.JSON(http.StatusNoContent, nil)
}

func (g *GenreController) AddContent(c echo.Context) error {
	genreContentDTO := GenreContentDTO{}
	if err := c.Bind(&genreContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&genreContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	id := c.Param("id")
	request := model.GenreContentRequest{
		GenreID:   id,
		ContentID: genreContentDTO.ContentID,
	}
	err := g.genreContentOperations.AddContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError := g.cacheRequestService.GenreRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}
	return c.JSON(http.StatusCreated, &MessageResource{Message: "Content added to the genre."})
}

func (g *GenreController) RemoveContent(c echo.Context) error {
	id := c.Param("id")
	requirementID := c.Param("contentID")
	request := model.GenreContentRequest{
		GenreID:   id,
		ContentID: requirementID,
	}
	err := g.genreContentOperations.RemoveContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Message: err.Message})
	}
	// content apiye cache icin request atiliyor
	serviceError := g.cacheRequestService.GenreRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	return c.JSON(http.StatusNoContent, nil)
}
