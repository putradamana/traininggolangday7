package routes

import (
	c "go-training-restful/controllers"
	"labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	//user routing
	e.GET("/users", c.GetUserController)

	return e
}
