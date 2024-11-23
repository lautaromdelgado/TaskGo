package tasks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateTask(c echo.Context) error {

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
