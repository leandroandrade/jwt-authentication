package control

import (
	"net/http"
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"fmt"
)

func Login(requestUser model.User, session *mgo.Session) ([]byte, int) {
	user, err := Authenticate(requestUser, session)
	if err != nil {
		return []byte(fmt.Sprintf(`{"error" : "%v"}`, err.Error())), http.StatusUnauthorized
	}

	token, err := generateToken(user.UUID)
	if err != nil {
		return []byte(fmt.Sprintf(`{"error" : "%v"}`, err.Error())), http.StatusInternalServerError
	}

	response, _ := json.Marshal(model.Token{Token: token})
	return response, http.StatusOK
}
