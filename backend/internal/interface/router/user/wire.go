//go:build wireinject
// +build wireinject

package user

import (
	"firebase-playground/internal/application/usecase"
	"firebase-playground/internal/infrastructure/firebase"
	"firebase-playground/internal/interface/handler"
	"firebase-playground/internal/interface/middleware"

	"github.com/google/wire"
)

//go:generate wire
func InitializeUserRouter() IRouter {
	wire.Build(
		NewRouter,
		handler.NewUserHandler,
		handler.NewAuthenticateHandler,
		usecase.NewAuthenticateUsecase,
	)

	return nil
}

//go:generate wire
func InitializeFirebaseGoogleAuthClient() (*middleware.FirebaseGoogleMiddleware, error) {
	wire.Build(
		firebase.NewFirebaseAuthClient,
		middleware.NewFirebaseGoogleMiddleware,
	)

	return nil, nil
}
