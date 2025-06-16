package user

import (
	"firebase-playground/internal/interface/handler"
	"log"

	"github.com/gin-gonic/gin"
)

type (
	IRouter interface {
		CreateRouter(router *gin.Engine)
	}

	router struct {
		userHandler handler.IUserHandler
		authHandler handler.IAuthenticateHandler
	}
)

func NewRouter(
	userHandler handler.IUserHandler,
	authHandler handler.IAuthenticateHandler) IRouter {
	return &router{
		userHandler: userHandler,
		authHandler: authHandler,
	}
}

func (r *router) CreateRouter(router *gin.Engine) {
	firebaseGoogleAuthClient, err := InitializeFirebaseGoogleAuthClient()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Google Auth client: %v", err)
	}

	{
		auth := router.Group("/auth", firebaseGoogleAuthClient.Authenticate())
		auth.POST("/login", r.authHandler.Login)
	}

	{
		user := router.Group("/user")
		user.GET("/", firebaseGoogleAuthClient.Authenticate(), r.userHandler.GetUser)
	}
}
