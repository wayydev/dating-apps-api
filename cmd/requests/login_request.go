package requests

type Login struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
