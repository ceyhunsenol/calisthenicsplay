package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GenreTypeController struct {
	genreTypeService service.IGenreTypeService
}

func NewGenreTypeController(genreTypeService service.IGenreTypeService) *GenreTypeController {
	return &GenreTypeController{genreTypeService: genreTypeService}
}

func (g *GenreTypeController) InitGenreTypeRoutes(e *echo.Group) {
	e.POST("", g.SaveGenreType)
	e.PUT("/:id", g.UpdateGenreType)
	e.GET("", g.GetGenreTypes)
	e.GET("/:id", g.GetGenreType)
	e.DELETE("/:id", g.DeleteGenreType)
}

func (g *GenreTypeController) SaveGenreType(c echo.Context) error {
	genreTypeDTO := GenreTypeDTO{}
	if err := c.Bind(&genreTypeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&genreTypeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}

	exists, err := g.genreTypeService.ExistsByCode(genreTypeDTO.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "GenreType could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "GenreType already exists in this code."})
	}

	genreType := data.GenreType{
		Code: genreTypeDTO.Code,
	}

	_, err = g.genreTypeService.Save(genreType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "GenreType could not be saved."})
	}
	return c.JSON(http.StatusCreated, &MessageResource{Code: http.StatusCreated, Message: "Created."})
}

func (g *GenreTypeController) UpdateGenreType(c echo.Context) error {
	genreTypeDTO := GenreTypeDTO{}
	if err := c.Bind(&genreTypeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "Invalid request format."})
	}
	if err := c.Validate(&genreTypeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id := c.Param("id")
	exists, err := g.genreTypeService.GetByCode(genreTypeDTO.Code)
	if err == nil && exists.ID != id {
		return c.JSON(http.StatusBadRequest, &MessageResource{Code: http.StatusBadRequest, Message: "GenreType already exists in this code."})
	}
	genreType, err := g.genreTypeService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "GenreType not found."})
	}
	genreType.Code = genreTypeDTO.Code
	_, err = g.genreTypeService.Update(*genreType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "GenreTypes could not be updated."})
	}
	return c.JSON(http.StatusOK, &MessageResource{Code: http.StatusOK, Message: "Updated."})
}

func (g *GenreTypeController) GetGenreTypes(c echo.Context) error {
	genreTypes, err := g.genreTypeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "GenreTypes could not be got."})
	}

	genreTypeResources := make([]GenreTypeResource, 0)
	for _, genreType := range genreTypes {
		genreTypeResources = append(genreTypeResources, GenreTypeResource{
			ID:   genreType.ID,
			Code: genreType.Code,
		})
	}

	return c.JSON(http.StatusOK, genreTypeResources)
}

func (g *GenreTypeController) GetGenreType(c echo.Context) error {
	id := c.Param("id")
	genreType, err := g.genreTypeService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Code: http.StatusNotFound, Message: "GenreType not found."})
	}

	genreTypeResource := GenreTypeResource{
		ID:   genreType.ID,
		Code: genreType.Code,
	}

	return c.JSON(http.StatusOK, genreTypeResource)
}

func (g *GenreTypeController) DeleteGenreType(c echo.Context) error {
	id := c.Param("id")
	err := g.genreTypeService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Code: http.StatusInternalServerError, Message: "GenreType could not be deleted."})
	}
	return c.JSON(http.StatusNoContent, nil)
}
