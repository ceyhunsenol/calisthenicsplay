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
}

func NewGenreController(genreContentOperations service.IGenreContentOperations,
	genreService service.IGenreService,
	contentTranslationOperations service.IContentTranslationOperations,
	DB *gorm.DB,
) *GenreController {
	return &GenreController{
		genreService:                 genreService,
		genreContentOperations:       genreContentOperations,
		contentTranslationOperations: contentTranslationOperations,
		DB:                           DB,
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
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&genreDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	exists, err := g.genreService.ExistsByCode(genreDTO.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Genre could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Genre already exists in this code."})
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
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Genre could not be saved."})
	}
	savedGenre, err := g.genreService.Save(genre)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Genre could not be saved."})
	}
	requests := make([]model.ContentTranslationRequest, 0)
	for _, translation := range genreDTO.Translations {
		request := model.ContentTranslationRequest{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
			ContentID: savedGenre.ID,
		}
		requests = append(requests, request)
	}
	serviceError := g.contentTranslationOperations.SaveContentTranslations(requests)
	if err != nil {
		tx.Rollback()
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	tx.Commit()
	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Created."})
}

func (g *GenreController) UpdateGenre(c echo.Context) error {
	genreDTO := GenreDTO{}
	if err := c.Bind(&genreDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&genreDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id := c.Param("id")
	exists, err := g.genreService.GetByCode(genreDTO.Code)
	if err == nil && exists.ID != id {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Genre already exists in this code."})
	}
	genre, err := g.genreService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Genre not found."})
	}
	genre.Type = genreDTO.Type
	genre.Code = genreDTO.Code
	genre.Description = genreDTO.Description
	genre.Section = genreDTO.Section
	genre.Active = genreDTO.Active
	tx := g.DB.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Genre could not be saved."})
	}
	_, err = g.genreService.Update(*genre)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Genre could not be updated."})
	}
	requests := make([]model.ContentTranslationRequest, 0)
	for _, translation := range genreDTO.Translations {
		request := model.ContentTranslationRequest{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
			ContentID: genre.ID,
		}
		requests = append(requests, request)
	}
	serviceError := g.contentTranslationOperations.SaveContentTranslations(requests)
	if err != nil {
		tx.Rollback()
		return c.JSON(serviceError.Code, &MessageResource{Code: serviceError.Code, Message: serviceError.Message})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &MessageResource{Code: http.StatusOK, Message: "Updated."})
}

func (g *GenreController) GetGenres(c echo.Context) error {
	genres, err := g.genreService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Genres could not be got."})
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
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "Genre not found."})
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
	err := g.genreService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "Genre could not be deleted."})
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (g *GenreController) AddContent(c echo.Context) error {
	genreContentDTO := GenreContentDTO{}
	if err := c.Bind(&genreContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&genreContentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id := c.Param("id")
	request := model.GenreContentRequest{
		GenreID:   id,
		ContentID: genreContentDTO.ContentID,
	}
	err := g.genreContentOperations.AddContent(request)
	if err != nil {
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}

	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Content added to the genre."})
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
		return c.JSON(err.Code, &MessageResource{Code: err.Code, Message: err.Message})
	}
	return c.JSON(http.StatusNoContent, nil)
}
