package usecase

import (
	"context"
	"firebase-playground/internal/application/usecase/command"
	"firebase-playground/internal/application/usecase/dto"
	"firebase-playground/pkg/jwt"
)

type authenticateUsecase struct {
}

func NewAuthenticateUsecase() IAuthenticateUsecase {
	return &authenticateUsecase{}
}

func (u *authenticateUsecase) Execute(ctx context.Context, cmd *command.AuthenticateCommand) (*dto.AuthenticateDto, error) {
	userID := ctx.Value("user_id").(string)
	token, err := jwt.GenerateJwt(userID)
	if err != nil {
		return nil, err
	}

	return dto.NewAuthenticateDto(token), nil
}
