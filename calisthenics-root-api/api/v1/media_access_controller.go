package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MediaAccessController struct {
	mediaAccessService service.IMediaAccessService
}

func NewMediaAccessController(mediaAccessService service.IMediaAccessService) *MediaAccessController {
	return &MediaAccessController{mediaAccessService: mediaAccessService}
}

func (m *MediaAccessController) InitMediaAccessRoutes(e *echo.Group) {
	e.POST("", m.SaveMediaAccess)
	e.PUT("/:id", m.UpdateMediaAccess)
	e.GET("", m.GetMediaAccessList)
	e.GET("/:id", m.GetMediaAccess)
	e.DELETE("/:id", m.DeleteMediaAccess)
}

func (m *MediaAccessController) SaveMediaAccess(ctx echo.Context) error {
	mediaAccessDTO := MediaAccessDTO{}
	if err := ctx.Bind(&mediaAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := ctx.Validate(&mediaAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	exists, err := m.mediaAccessService.ExistsByMediaID(mediaAccessDTO.MediaID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "MediaAccess could not be saved"})
	}
	if exists {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "MediaAccess already exists in this mediaID"})
	}

	mediaAccess := data.MediaAccess{
		MediaID:  mediaAccessDTO.MediaID,
		Audience: mediaAccessDTO.Audience,
	}

	_, err = m.mediaAccessService.Save(mediaAccess)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "MediaAccess could not be saved."})
	}

	return ctx.JSON(http.StatusCreated, &MessageResource{Message: "Created."})
}

func (m *MediaAccessController) UpdateMediaAccess(ctx echo.Context) error {
	mediaAccessDTO := MediaAccessDTO{}
	if err := ctx.Bind(&mediaAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := ctx.Validate(&mediaAccessDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	id := ctx.Param("id")
	exists, err := m.mediaAccessService.GetByMediaID(mediaAccessDTO.MediaID)
	if err == nil && exists.ID != id {
		return ctx.JSON(http.StatusBadRequest, &MessageResource{Message: "MediaAccess already exists in this mediaID."})
	}
	mediaAccess, err := m.mediaAccessService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &MessageResource{Message: "MediaAccess not found."})
	}

	mediaAccess.MediaID = mediaAccessDTO.MediaID
	mediaAccess.Audience = mediaAccessDTO.Audience

	_, err = m.mediaAccessService.Update(*mediaAccess)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "MediaAccess could not be updated."})
	}

	return ctx.JSON(http.StatusOK, &MessageResource{Message: "Updated."})
}

func (m *MediaAccessController) GetMediaAccessList(ctx echo.Context) error {
	mediaAccessList, err := m.mediaAccessService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "MediaAccess list could not be retrieved."})
	}

	mediaAccessResources := make([]MediaAccessResource, 0)
	for _, mediaAccess := range mediaAccessList {
		mediaAccessResources = append(mediaAccessResources, MediaAccessResource{
			ID:       mediaAccess.ID,
			MediaID:  mediaAccess.MediaID,
			Audience: mediaAccess.Audience,
		})
	}

	return ctx.JSON(http.StatusOK, mediaAccessResources)
}

func (m *MediaAccessController) GetMediaAccess(ctx echo.Context) error {
	id := ctx.Param("id")
	mediaAccess, err := m.mediaAccessService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &MessageResource{Message: "MediaAccess not found."})
	}

	mediaAccessResource := MediaAccessResource{
		ID:       mediaAccess.ID,
		MediaID:  mediaAccess.MediaID,
		Audience: mediaAccess.Audience,
	}

	return ctx.JSON(http.StatusOK, mediaAccessResource)
}

func (m *MediaAccessController) DeleteMediaAccess(ctx echo.Context) error {
	id := ctx.Param("id")
	err := m.mediaAccessService.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &MessageResource{Message: "MediaAccess could not be deleted."})
	}
	return ctx.JSON(http.StatusNoContent, nil)
}
