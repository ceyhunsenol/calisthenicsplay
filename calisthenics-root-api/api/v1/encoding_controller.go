package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EncodingController struct {
	encodingService service.IEncodingService
}

func NewEncodingController(encodingService service.IEncodingService) *EncodingController {
	return &EncodingController{encodingService: encodingService}
}

func (e *EncodingController) InitEncodingRoutes(g *echo.Group) {
	g.POST("", e.SaveEncoding)
	g.PUT("/:id", e.UpdateEncoding)
	g.GET("", e.GetEncodings)
	g.GET("/:id", e.GetEncoding)
	g.DELETE("/:id", e.DeleteEncoding)
}

func (e *EncodingController) SaveEncoding(c echo.Context) error {
	encodingDTO := EncodingDTO{}
	if err := c.Bind(&encodingDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format"})
	}
	if err := c.Validate(&encodingDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	encoding := data.Encoding{
		LicenseKey: encodingDTO.LicenseKey,
		MediaID:    encodingDTO.MediaID,
	}

	_, err := e.encodingService.Save(encoding)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encoding could not be saved"})
	}
	return c.JSON(http.StatusCreated, &MessageResource{Message: "Created."})
}

func (e *EncodingController) UpdateEncoding(c echo.Context) error {
	encodingDTO := EncodingDTO{}
	if err := c.Bind(&encodingDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&encodingDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	id := c.Param("id")
	encoding, err := e.encodingService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Encoding not found"})
	}
	encoding.LicenseKey = encodingDTO.LicenseKey
	encoding.MediaID = encodingDTO.MediaID
	_, err = e.encodingService.Update(*encoding)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encoding could not be updated"})
	}
	return c.JSON(http.StatusOK, &MessageResource{Message: "Updated"})
}

func (e *EncodingController) GetEncodings(c echo.Context) error {
	encodings, err := e.encodingService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encodings could not be got"})
	}

	encodingResources := make([]EncodingResource, 0)
	for _, encoding := range encodings {
		encodingResources = append(encodingResources, EncodingResource{
			ID:         encoding.ID,
			LicenseKey: encoding.LicenseKey,
			MediaID:    encoding.MediaID,
		})
	}

	return c.JSON(http.StatusOK, encodingResources)
}

func (e *EncodingController) GetEncoding(c echo.Context) error {
	id := c.Param("id")
	encoding, err := e.encodingService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Encoding not found"})
	}

	encodingResource := EncodingResource{
		ID:         encoding.ID,
		LicenseKey: encoding.LicenseKey,
		MediaID:    encoding.MediaID,
	}

	return c.JSON(http.StatusOK, encodingResource)
}

func (e *EncodingController) DeleteEncoding(c echo.Context) error {
	id := c.Param("id")
	err := e.encodingService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encoding could not be deleted"})
	}
	return c.JSON(http.StatusNoContent, nil)
}
