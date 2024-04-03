package controller

import (
	"github.com/baukeyzh/todo/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service service.TaskService
}

func NewTaskController(service service.TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

// CreateTask Пример метода контроллера
func (tc *TaskController) CreateTask(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

// GetTasks Пример метода контроллера
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.service.GetTasks(c, time.Now())
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// Остальные методы контроллера...
