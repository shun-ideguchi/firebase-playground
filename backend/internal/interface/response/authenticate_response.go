package response

import "firebase-playground/internal/application/usecase/dto"

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func NewAuthenticateResponse(dto *dto.AuthenticateDto) *AuthenticateResponse {
	return &AuthenticateResponse{
		Token: dto.Token,
	}
}
