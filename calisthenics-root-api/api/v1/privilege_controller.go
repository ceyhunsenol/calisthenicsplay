package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PrivilegeController struct {
	privilegeService service.IPrivilegeService
}

func NewPrivilegeController(privilegeService service.IPrivilegeService) *PrivilegeController {
	return &PrivilegeController{privilegeService: privilegeService}
}

func (u *PrivilegeController) InitPrivilegeRoutes(e *echo.Group) {
	e.POST("", u.SavePrivilege)
	e.PUT("/:id", u.UpdatePrivilege)
	e.GET("", u.GetPrivileges)
	e.GET("/:id", u.GetPrivilege)
	e.DELETE("/:id", u.DeletePrivilege)
}

func (u *PrivilegeController) SavePrivilege(c echo.Context) error {
	privilegeDTO := PrivilegeDTO{}
	if err := c.Bind(&privilegeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&privilegeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	exists, err := u.privilegeService.ExistsByCode(privilegeDTO.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Privilege could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Privilege already exists in this code."})
	}

	privilege := data.Privilege{
		Code:          privilegeDTO.Code,
		Description:   privilegeDTO.Description,
		EndpointsJoin: privilegeDTO.EndpointsJoin,
	}

	_, err = u.privilegeService.Save(privilege)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Privilege could not be saved."})
	}
	return c.JSON(http.StatusCreated, &MessageResource{Message: "Created."})
}

func (u *PrivilegeController) UpdatePrivilege(c echo.Context) error {
	roleDTO := PrivilegeDTO{}
	if err := c.Bind(&roleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&roleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	id := c.Param("id")
	exists, err := u.privilegeService.GetByCode(roleDTO.Code)
	if err == nil && exists.ID != id {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Privilege already exists in this code."})
	}
	role, err := u.privilegeService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Privilege not found."})
	}
	role.Code = roleDTO.Code
	role.Description = roleDTO.Description
	role.EndpointsJoin = roleDTO.EndpointsJoin
	_, err = u.privilegeService.Update(*role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Privilege could not be updated."})
	}
	return c.JSON(http.StatusOK, &MessageResource{Message: "Updated."})
}

func (u *PrivilegeController) GetPrivileges(c echo.Context) error {
	privileges, err := u.privilegeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Privileges could not be got."})
	}

	privilegeResources := make([]PrivilegeResource, 0)
	for _, privilege := range privileges {
		privilegeResources = append(privilegeResources, PrivilegeResource{
			ID:            privilege.ID,
			Code:          privilege.Code,
			Description:   privilege.Description,
			EndpointsJoin: privilege.EndpointsJoin,
		})
	}

	return c.JSON(http.StatusOK, privilegeResources)
}

func (u *PrivilegeController) GetPrivilege(c echo.Context) error {
	id := c.Param("id")
	privilege, err := u.privilegeService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Privilege not found."})
	}

	privilegeResource := PrivilegeResource{
		ID:            privilege.ID,
		Code:          privilege.Code,
		Description:   privilege.Description,
		EndpointsJoin: privilege.EndpointsJoin,
	}

	return c.JSON(http.StatusOK, privilegeResource)
}

func (u *PrivilegeController) DeletePrivilege(c echo.Context) error {
	id := c.Param("id")
	err := u.privilegeService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Privilege could not be deleted."})
	}
	return c.JSON(http.StatusNoContent, nil)
}
