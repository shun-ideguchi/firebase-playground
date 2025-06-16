package router

import (
	"firebase-playground/config"
	"firebase-playground/internal/interface/router/user"
	"log"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	_, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	router := gin.Default()
	userRouter := user.InitializeUserRouter()
	userRouter.CreateRouter(router)

	return router
}
