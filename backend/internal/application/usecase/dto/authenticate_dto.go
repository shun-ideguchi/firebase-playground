package dto

type AuthenticateDto struct {
	Token string
}

func NewAuthenticateDto(token string) *AuthenticateDto {
	return &AuthenticateDto{
		Token: token,
	}
}
