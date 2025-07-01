package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type FirebaseGoogleMiddleware struct {
	authClient *auth.Client
}

func NewFirebaseGoogleMiddleware(
	authClient *auth.Client,
) *FirebaseGoogleMiddleware {
	return &FirebaseGoogleMiddleware{
		authClient: authClient,
	}
}

func (f *FirebaseGoogleMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		token, err := f.authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Fatalf("error verifying ID token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// ctx := context.WithValue(c.Request.Context(), "user_id", token.Claims["user_id"])
		ctx := context.WithValue(c.Request.Context(), "firebase_token", token)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
