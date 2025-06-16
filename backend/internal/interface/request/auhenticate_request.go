package request

type AuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
