package api

import (
	"calisthenics-content-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContentController struct {
	contentOperations service.IContentOperations
}

func NewContentController(contentOperations service.IContentOperations) *ContentController {
	return &ContentController{
		contentOperations: contentOperations,
	}
}

func (u *ContentController) InitContentRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/contents")
	v1.GET("/content/code/:code", u.GetContentByCode)
}

func (u *ContentController) GetContentByCode(c echo.Context) error {
	code := c.Param("code")
	contentModel, serviceError := u.contentOperations.GetContentByCode(code)
	if serviceError != nil {
		return c.JSON(serviceError.Code, &MessageResource{Message: serviceError.Message})
	}
	contentResource := ContentResource{ID: contentModel.ID}
	return c.JSON(http.StatusOK, contentResource)
}
