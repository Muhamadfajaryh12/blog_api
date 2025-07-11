package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
)

type Error interface {
	Error() string
}

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Message string
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

func ErrorHandle(c *gin.Context, err Error) {
	var statusCode int
	var Message string

	switch  e:= err.(type) {
	case NotFoundError:
		statusCode = http.StatusNotFound
		Message = e.Message
	case BadRequestError:
		statusCode = http.StatusBadRequest	
		Message = e.Message
	case InternalServerError:
		statusCode = http.StatusInternalServerError	
		Message = e.Message
	case UnauthorizedError:
		statusCode = http.StatusUnauthorized	
		Message = e.Message
	default:
		statusCode = http.StatusInternalServerError
		Message = "Unexpected error occurred"
	}
	
	c.JSON(statusCode, dto.ResponseErrorDTO{
		Status:statusCode,
		Message: Message,
	})
}