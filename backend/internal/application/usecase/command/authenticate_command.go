package command

type AuthenticateCommand struct {
	Email string
}

func NewAuthenticateCommand(e string) *AuthenticateCommand {
	return &AuthenticateCommand{
		Email: e,
	}
}
