package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/chucknorris/src/adapters/rest/dto"
	"github.com/oswaldom-code/chucknorris/src/services"
)

// Pong is a simple ping-pong controller
func Pong(c *gin.Context) {
	response := dto.MessageResponse{
		Status:  "success",
		Message: "pong",
	}
	c.JSON(http.StatusOK, response)
}

// Call is a controller that calls the service
func Call(c *gin.Context) {
	// get the service
	service := services.NewCallService()
	call, err := service.Call()
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := dto.ResponseWithData{
		Status:  "success",
		Message: fmt.Sprintf("Total data: %v", len(call)),
		Data:    call,
	}
	c.JSON(http.StatusOK, response)
}
