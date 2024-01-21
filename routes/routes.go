package routes

import (
	"github.com/abimo2020/book-store/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	// e.POST("/book")

	return e
}
