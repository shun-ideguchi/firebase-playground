package usecase

import (
	"context"
	"firebase-playground/internal/application/usecase/command"
	"firebase-playground/internal/application/usecase/dto"
	"firebase-playground/pkg/jwt"

	"firebase.google.com/go/v4/auth"
)

type authenticateUsecase struct {
}

func NewAuthenticateUsecase() IAuthenticateUsecase {
	return &authenticateUsecase{}
}

func (u *authenticateUsecase) Execute(
	ctx context.Context, cmd *command.AuthenticateCommand,
) (*dto.AuthenticateDto, error) {
	firebaseToken := ctx.Value("firebase_token").(*auth.Token)
	token, err := jwt.GenerateJwt(firebaseToken.UID)
	if err != nil {
		return nil, err
	}

	return dto.NewAuthenticateDto(token), nil
}
