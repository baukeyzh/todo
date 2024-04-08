package errorHandler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func HandleValidationErrors(err error) map[string]string {
	var errorMessages = make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		var field = err.Field()
		var tag = err.Tag()

		switch tag {
		case "required":
			errorMessages[field] = "The field is required"
		case "max":
			errorMessages[field] = "The maximum length is " + err.Param()
		case "min":
			errorMessages[field] = "The minimum length is " + err.Param()
		}
	}
	return errorMessages
}

func CheckError(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		c.JSON(http.StatusBadRequest, gin.H{"error": HandleValidationErrors(errs)})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Validation error"})
}
