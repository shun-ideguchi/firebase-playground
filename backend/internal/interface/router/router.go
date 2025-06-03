package router

import (
	"firebase-playground/internal/interface/router/firebase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	firebaseRouter := firebase.InitializeFirebaseRouter()
	firebaseRouter.CreateRouter(router)

	return router
}
