package handler

import "github.com/gin-gonic/gin"

type IAuthenticateHandler interface {
	Login(c *gin.Context)
}
