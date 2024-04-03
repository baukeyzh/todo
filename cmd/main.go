package main

import (
	"github.com/baukeyzh/todo"
	"github.com/baukeyzh/todo/controller"
	"github.com/baukeyzh/todo/db"
	"github.com/baukeyzh/todo/repository"
	"github.com/baukeyzh/todo/service"
	"github.com/gin-gonic/gin"
	"log"
)

// main.go или в соответствующих файлах роутинга и контроллеров
func main() {
	mongoDb, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	taskRepo := repository.NewMongoTaskRepository(mongoDb)
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService)

	router := gin.Default()

	todo.RegisterRoutes(router, taskController)

	log.Fatal(router.Run(":8080"))
}
