//go:build wireinject
// +build wireinject

package firebase

import (
	"firebase-playground/internal/interface/handler"

	"github.com/google/wire"
)

//go:generate wire
func InitializeFirebaseRouter() IRouter {
	wire.Build(
		NewRouter,
		handler.NewFirebaseHandler,
	)

	return nil
}
