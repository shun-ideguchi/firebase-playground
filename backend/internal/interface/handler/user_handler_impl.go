package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

func NewUserHandler() IUserHandler {
	return &userHandler{}
}

func (h *userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "John Doe"})
}
