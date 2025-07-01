package handler

import (
	"net/http"

	"firebase-playground/internal/application/usecase"
	"firebase-playground/internal/application/usecase/command"
	"firebase-playground/internal/interface/request"
	"firebase-playground/internal/interface/response"

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

	var req request.AuthenticateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := command.NewAuthenticateCommand(req.Email)
	dto, err := h.authenticateUsecase.Execute(ctx, cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := response.NewAuthenticateResponse(dto)

	c.JSON(http.StatusOK, res)
}
