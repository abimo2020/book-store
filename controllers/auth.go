package controllers

import (
	"net/http"

	"github.com/abimo2020/book-store/middlewares"
	"github.com/abimo2020/book-store/models"
	"github.com/abimo2020/book-store/models/payload"
	usecase "github.com/abimo2020/book-store/usecases"
	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) (err error) {
	var req payload.UserRequest
	var resp *models.Users
	var token string

	c.Bind(&req)

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if resp, err = usecase.Register(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if token, err = middlewares.CreateToken(resp.Email, resp.IsAdmin, req.ID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
		"token":   token,
	})
}

func LoginController(c echo.Context) (err error) {
	var req payload.UserRequest
	var resp *models.Users
	var token string

	c.Bind(&req)

	if err = c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if resp, err = usecase.Login(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if token, err = middlewares.CreateToken(resp.Email, resp.IsAdmin, req.ID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to login",
		"token":   token,
	})
}
