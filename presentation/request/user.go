package request

type CreateUserAccountRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Pass     string `json:"password"`
}
