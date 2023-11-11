package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContentAccessController struct {
	contentAccessService service.IContentAccessService
}

func NewContentAccessController(contentAccessService service.IContentAccessService) *ContentAccessController {
	return &ContentAccessController{contentAccessService: contentAccessService}
}

func (c *ContentAccessController) InitContentAccessRoutes(e *echo.Group) {
	e.POST("", c.SaveContentAccess)
	e.PUT("/:id", c.UpdateContentAccess)
	e.GET("", c.GetContentAccessList)
	e.GET("/:id", c.GetContentAccess)
	e.DELETE("/:id", c.DeleteContentAccess)
}

func (c *ContentAccessController) SaveContentAccess(ctx echo.Context) error {
	contentAccessDTO := ContentAccessDTO{}
	if err := ctx.Bind(&contentAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := ctx.Validate(&contentAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	exists, err := c.contentAccessService.ExistsByContentID(contentAccessDTO.ContentID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "ContentAccess could not be saved"})
	}
	if exists {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "ContentAccess already exists in this contentID"})
	}

	contentAccess := data.ContentAccess{
		ContentID: contentAccessDTO.ContentID,
		Audience:  contentAccessDTO.Audience,
	}

	_, err = c.contentAccessService.Save(contentAccess)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "ContentAccess could not be saved."})
	}

	return ctx.JSON(http.StatusCreated, &MessageResource{Message: "Created."})
}

func (c *ContentAccessController) UpdateContentAccess(ctx echo.Context) error {
	contentAccessDTO := ContentAccessDTO{}
	if err := ctx.Bind(&contentAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := ctx.Validate(&contentAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	id := ctx.Param("id")
	exists, err := c.contentAccessService.GetByContentID(contentAccessDTO.ContentID)
	if err == nil && exists.ID != id {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "ContentAccess already exists in this contentID."})
	}
	contentAccess, err := c.contentAccessService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &MessageResource{Message: "ContentAccess not found."})
	}

	contentAccess.ContentID = contentAccessDTO.ContentID
	contentAccess.Audience = contentAccessDTO.Audience

	_, err = c.contentAccessService.Update(*contentAccess)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "ContentAccess could not be updated."})
	}

	return ctx.JSON(http.StatusOK, &MessageResource{Message: "Updated."})
}

func (c *ContentAccessController) GetContentAccessList(ctx echo.Context) error {
	contentAccessList, err := c.contentAccessService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "ContentAccess list could not be retrieved."})
	}

	contentAccessResources := make([]ContentAccessResource, 0)
	for _, contentAccess := range contentAccessList {
		contentAccessResources = append(contentAccessResources, ContentAccessResource{
			ID:        contentAccess.ID,
			ContentID: contentAccess.ContentID,
			Audience:  contentAccess.Audience,
		})
	}

	return ctx.JSON(http.StatusOK, contentAccessResources)
}

func (c *ContentAccessController) GetContentAccess(ctx echo.Context) error {
	id := ctx.Param("id")
	contentAccess, err := c.contentAccessService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &MessageResource{Message: "ContentAccess not found."})
	}

	contentAccessResource := ContentAccessResource{
		ID:        contentAccess.ID,
		ContentID: contentAccess.ContentID,
		Audience:  contentAccess.Audience,
	}

	return ctx.JSON(http.StatusOK, contentAccessResource)
}

func (c *ContentAccessController) DeleteContentAccess(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.contentAccessService.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "ContentAccess could not be deleted."})
	}
	return ctx.JSON(http.StatusNoContent, nil)
}
