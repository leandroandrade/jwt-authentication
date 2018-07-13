package control

import (
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func Refresh(req *http.Request, session *mgo.Session) ([]byte, int) {
	tokenUser, err := extractToken(req)
	if err != nil {
		return []byte(fmt.Sprintf(`{"error" : "%v"}`, err.Error())), http.StatusUnauthorized
	}

	if err = invalidateToken(tokenUser.Raw, session); err != nil {
		return []byte(fmt.Sprintf(`{"error" : "%v"}`, err.Error())), http.StatusInternalServerError
	}

	uuid := tokenUser.Claims.(jwt.MapClaims)["sub"].(string)
	token, err := generateToken(uuid)
	if err != nil {
		return []byte(fmt.Sprintf(`{"error" : "%v"}`, err.Error())), http.StatusInternalServerError
	}

	response, err := json.Marshal(model.Token{Token: token})
	if err != nil {
		return []byte(fmt.Sprintf(`{"error" : "%v"}`, err.Error())), http.StatusInternalServerError
	}

	return response, http.StatusOK
}
