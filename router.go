package todo

import (
	"github.com/baukeyzh/todo/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, taskController *controller.TaskController) {
	router.GET("/api/todo-list/tasks", taskController.GetTasks)
	router.GET("/api/todo-list/tasks/:id", taskController.GetTask)
	router.POST("/api/todo-list/tasks", taskController.PostTask)
	router.PUT("/api/todo-list/tasks/:id", taskController.PutTask)
	router.DELETE("/api/todo-list/tasks/:id", taskController.DeleteTask)
	router.PATCH("/api/todo-list/tasks/:id/done", taskController.MarkTaskDone)
}
