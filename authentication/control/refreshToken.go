package control

import (
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"encoding/json"
)

func Refresh(user model.User) []byte {
	token, err := generateToken(user.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(model.Token{Token: token})
	if err != nil {
		panic(err)
	}
	return response
}
