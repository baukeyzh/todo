package controller

import (
	"github.com/baukeyzh/todo/errorHandler"
	"github.com/baukeyzh/todo/models"
	"github.com/baukeyzh/todo/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type TaskController struct {
	service  service.TaskService
	validate *validator.Validate
}

func NewTaskController(service service.TaskService, validate *validator.Validate) *TaskController {
	return &TaskController{
		service:  service,
		validate: validate,
	}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	status := c.Param("status")
	if len(status) == 0 {
		status = "active"
	} else {
		var statusValidate = models.StatusValidate{Status: status}
		if err := tc.validate.Struct(statusValidate); err != nil {
			errorHandler.CheckError(c, err)
			return
		}
	}

	tasks, err := tc.service.ReadTasks(c, status)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	var idValidate = models.IdValidate{ID: id}
	if err := tc.validate.Struct(idValidate); err != nil {
		errorHandler.CheckError(c, err)
		return
	}

	task, err := tc.service.ReadTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) PostTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.validate.Struct(task); err != nil {
		errorHandler.CheckError(c, err)
		return
	}
	res, err := tc.service.CreateTask(c, task)
	if mongo.IsDuplicateKeyError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task already exist"})
		return
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": res + " Task created"})
}

func (tc *TaskController) PutTask(c *gin.Context) {
	id := c.Param("id")
	var taskForm = models.TaskForm{ID: id}

	if err := c.BindJSON(&taskForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.validate.Struct(taskForm); err != nil {
		errorHandler.CheckError(c, err)
		return
	}

	err := tc.service.UpdateTask(c, taskForm)
	if mongo.IsDuplicateKeyError(err) {
		c.JSON(http.StatusConflict, gin.H{"error": "Task already exist"})
		return
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": id + " task edited"})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var idValidate = models.IdValidate{ID: id}

	if err := tc.validate.Struct(idValidate); err != nil {
		errorHandler.CheckError(c, err)
		return
	}
	err := tc.service.DeleteTask(c, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": id + "task deleted"})
}

func (tc *TaskController) MarkTaskDone(c *gin.Context) {
	id := c.Param("id")
	var idValidate = models.IdValidate{ID: id}

	if err := tc.validate.Struct(idValidate); err != nil {
		errorHandler.CheckError(c, err)
		return
	}
	err := tc.service.MarkTaskDone(c, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": id + "task deleted"})
}
