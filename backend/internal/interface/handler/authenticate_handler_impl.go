package handler

import (
	"fmt"
	"net/http"

	"firebase-playground/internal/application/usecase"
	"firebase-playground/internal/application/usecase/command"
	"firebase-playground/internal/interface/request"

	"github.com/gin-gonic/gin"
)

type authenticateHandler struct {
	authenticateUsecase usecase.IAuthenticateUsecase
}

func NewAuthenticateHandler(
	authenticateUsecase usecase.IAuthenticateUsecase,
) IAuthenticateHandler {
	return &authenticateHandler{
		authenticateUsecase: authenticateUsecase,
	}
}

func (h *authenticateHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	fmt.Println(ctx)

	var req request.AuthenticateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: メールだけに修正
	cmd := command.NewAuthenticateCommand(req.Email, req.Password)
	token, err := h.authenticateUsecase.Execute(ctx, cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
