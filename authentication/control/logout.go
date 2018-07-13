package control

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/leandroandrade/jwt-authentication/rsa"
	"gopkg.in/mgo.v2"
)

func Logout(req *http.Request, session *mgo.Session) error {
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return rsa.PublicKey, nil
	})
	if err != nil {
		return err
	}
	return invalidateToken(tokenRequest.Raw, session)
}
