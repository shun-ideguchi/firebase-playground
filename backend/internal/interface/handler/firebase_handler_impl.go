package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type firebaseHandler struct {
}

func NewFirebaseHandler() IFirebaseHandler {
	return &firebaseHandler{}
}

func (h *firebaseHandler) Authenticate(c *gin.Context) {
	fmt.Println("Authenticate!")
	c.JSON(http.StatusOK, gin.H{"message": "Authenticate"})
}
