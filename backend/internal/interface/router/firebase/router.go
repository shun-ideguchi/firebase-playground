package firebase

import (
	"firebase-playground/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

type (
	IRouter interface {
		CreateRouter(router *gin.Engine)
	}

	router struct {
		firebaseHandler handler.IFirebaseHandler
	}
)

func NewRouter(firebaseHandler handler.IFirebaseHandler) IRouter {
	return &router{firebaseHandler: firebaseHandler}
}

func (r *router) CreateRouter(router *gin.Engine) {
	{
		firebase := router.Group("/firebase")
		firebase.GET("/authenticate", r.firebaseHandler.Authenticate)
	}
}
