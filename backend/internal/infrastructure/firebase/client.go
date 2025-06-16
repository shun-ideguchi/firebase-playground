package firebase

import (
	"context"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func NewFirebaseAuthClient() (*auth.Client, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile(
		filepath.Join(
			currentDir,
			"internal/interface/middleware/react-golang-playground-firebase-adminsdk-fbsvc-f6c46c6afb.json",
		),
	)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return app.Auth(context.Background())
}
