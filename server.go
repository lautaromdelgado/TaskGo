package main

import (
	"taskgo/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Inicialización de la base de datos.
	database.InitDataBase()

	// Inicialización de las rutas.

	e.Logger.Fatal(e.Start(":7000"))
}
