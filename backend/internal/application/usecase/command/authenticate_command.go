package command

type AuthenticateCommand struct {
	Email    string
	Password string
}

func NewAuthenticateCommand(email, password string) *AuthenticateCommand {
	return &AuthenticateCommand{
		Email:    email,
		Password: password,
	}
}
