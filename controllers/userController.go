package controllers

import (
	"go-training-restful/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return echo.JSON(http.StatusOK, user)
}
