package responses

import "dating-apps/api/cmd/models"

type Login struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}
