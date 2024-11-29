package tasks

import (
	"net/http"
	"taskgo/database"
	"taskgo/models"

	"github.com/labstack/echo/v4"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type Data struct {
	Task *models.Tasks
}

func CreateTask(c echo.Context) error {
	db := database.GetDB()
	task := new(models.Tasks)

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  "error",
			Message: "invalid request " + err.Error(),
		})
	}

	if task.Title == "" || task.Nivel < 0 {
		return c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  "error",
			Message: "cant have empty fields",
		})
	}

	if err := db.Create(&task); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Status:  "error",
			Message: "error creating task",
		})
	}

	data := Data{
		Task: task,
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
