package api

import (
	"calisthenics-content-api/model"
	"calisthenics-content-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GenreController struct {
	genreOperations service.IGenreOperations
}

func NewGenreController(genreOperations service.IGenreOperations) *GenreController {
	return &GenreController{
		genreOperations: genreOperations,
	}
}

func (u *GenreController) InitGenreRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/genres")
	v1.GET("", u.GetGenres)
}

func (u *GenreController) GetGenres(c echo.Context) error {
	genreType := c.QueryParam("type")
	section := c.QueryParam("section")

	if genreType == "" {
		return c.JSON(http.StatusBadRequest, MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format"})
	}

	request := model.GenreRequest{
		Type:    genreType,
		Section: section,
	}
	genres := u.genreOperations.GetGenres(request)
	return c.JSON(http.StatusOK, genres)
}
