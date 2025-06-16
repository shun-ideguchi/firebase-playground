package handler

import "github.com/gin-gonic/gin"

type IUserHandler interface {
	GetUser(c *gin.Context)
}
