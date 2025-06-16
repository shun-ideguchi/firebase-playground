package usecase

import (
	"context"
	"firebase-playground/internal/application/usecase/command"
	"firebase-playground/internal/application/usecase/dto"
)

type IAuthenticateUsecase interface {
	Execute(ctx context.Context, cmd *command.AuthenticateCommand) (*dto.AuthenticateDto, error)
}
