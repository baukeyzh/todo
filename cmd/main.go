package main

import (
	"github.com/baukeyzh/todo"
	"github.com/baukeyzh/todo/controller"
	"github.com/baukeyzh/todo/db"
	"github.com/baukeyzh/todo/repository"
	"github.com/baukeyzh/todo/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

func main() {
	mongoDb, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	var validate = validator.New()
	taskRepo := repository.NewMongoTaskRepository(mongoDb)
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService, validate)

	router := gin.Default()

	todo.RegisterRoutes(router, taskController)

	log.Fatal(router.Run(":8080"))
}
