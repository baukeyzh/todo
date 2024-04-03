package todo

import (
	"github.com/baukeyzh/todo/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, taskController *controller.TaskController) {
	router.GET("/api/todo-list/tasks", taskController.GetTasks)
	router.POST("/api/todo-list/tasks", taskController.CreateTask)
	// Определите остальные маршруты, используя taskController
}
