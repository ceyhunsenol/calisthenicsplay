package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RoleController struct {
	roleService service.IRoleService
}

func NewRoleController(roleService service.IRoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

func (u *RoleController) InitRoleRoutes(e *echo.Group) {
	e.POST("", u.SaveRole)
	e.PUT("/:id", u.UpdateRole)
	e.GET("", u.GetRoles)
	e.GET("/:id", u.GetRole)
	e.DELETE("/:id", u.DeleteRole)
}

func (u *RoleController) SaveRole(c echo.Context) error {
	roleDTO := RoleDTO{}
	if err := c.Bind(&roleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err := c.Validate(&roleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}

	exists, err := u.roleService.ExistsByCode(roleDTO.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Role could not be saved."})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Role already exists in this code."})
	}

	role := data.Role{
		Code: roleDTO.Code,
	}

	_, err = u.roleService.Save(role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Role could not be saved."})
	}
	return c.JSON(http.StatusCreated, &MessageResource{Message: "Created."})
}

func (u *RoleController) UpdateRole(c echo.Context) error {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}
	roleDTO := RoleDTO{}
	if err = c.Bind(&roleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Invalid request format."})
	}
	if err = c.Validate(&roleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: err.Error()})
	}
	exists, err := u.roleService.GetByCode(roleDTO.Code)
	if err == nil && exists.ID != uint(idUint) {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Role already exists in this code."})
	}
	role, err := u.roleService.GetByID(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Role not found."})
	}
	role.Code = roleDTO.Code
	_, err = u.roleService.Update(*role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Role could not be updated."})
	}
	return c.JSON(http.StatusOK, &MessageResource{Message: "Updated."})
}

func (u *RoleController) GetRoles(c echo.Context) error {
	roles, err := u.roleService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Roles could not be got."})
	}

	var roleResources []RoleResource
	for _, role := range roles {
		roleResources = append(roleResources, RoleResource{
			ID:   role.ID,
			Code: role.Code,
		})
	}

	return c.JSON(http.StatusOK, roleResources)
}

func (u *RoleController) GetRole(c echo.Context) error {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}
	role, err := u.roleService.GetByID(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, &MessageResource{Message: "Role not found."})
	}

	privilegeResources := make([]PrivilegeResource, 0)
	for _, privilege := range role.Privileges {
		privilegeResources = append(privilegeResources, PrivilegeResource{
			ID:            privilege.ID,
			Code:          privilege.Code,
			Description:   privilege.Description,
			EndpointsJoin: privilege.EndpointsJoin,
		})
	}
	roleResource := RoleResource{
		ID:         role.ID,
		Code:       role.Code,
		Privileges: privilegeResources,
	}
	return c.JSON(http.StatusOK, roleResource)
}

func (u *RoleController) DeleteRole(c echo.Context) error {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}
	err = u.roleService.Delete(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Roles could not be deleted."})
	}
	return c.JSON(http.StatusNoContent, nil)
}
