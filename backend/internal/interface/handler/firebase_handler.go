package handler

import "github.com/gin-gonic/gin"

type IFirebaseHandler interface {
	Authenticate(c *gin.Context)
}
