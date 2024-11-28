package routes

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	a := e.Group("/api")    // Para uso con JWT
	b := e.Group("/public") // Para uso público

}
