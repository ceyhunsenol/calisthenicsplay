package v1

import (
	"calisthenics-root-api/service"
	"fmt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) InitUserRoutes(e *echo.Group) {
	e.POST("/test1", u.Test1)
	e.POST("/test2", u.Test2)
	e.POST("/test2/child", u.Test2Child)
	e.POST("/test2/child2", u.Test2Child2)
	e.POST("/test1/child", u.Test1Child)
	e.POST("/test1/child2", u.Test1Child2)
}

func (u *UserController) Test1(c echo.Context) error {

	fmt.Println("Test1.")
	return nil
}

func (u *UserController) Test2(c echo.Context) error {

	fmt.Println("Test2.")
	return nil
}

func (u *UserController) Test2Child(c echo.Context) error {

	fmt.Println("Test2Child.")
	return nil
}

func (u *UserController) Test2Child2(c echo.Context) error {

	fmt.Println("Test2Child2.")
	return nil
}

func (u *UserController) Test1Child(c echo.Context) error {

	fmt.Println("Test1Child.")
	return nil
}

func (u *UserController) Test1Child2(c echo.Context) error {

	fmt.Println("Test1Child2.")
	return nil
}
