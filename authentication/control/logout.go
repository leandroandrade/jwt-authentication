package control

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/leandroandrade/jwt-authentication/rsa"
	"gopkg.in/mgo.v2"
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"time"
	"log"
)

func Logout(req *http.Request, session *mgo.Session) error {
	defer session.Close()

	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return rsa.PublicKey, nil
	})
	if err != nil {
		return err
	}

	collection := session.DB("jwt_authentication").C("expired")

	token := model.Token{Token: tokenRequest.Raw, CreatedAt: time.Now()}
	if err := collection.Insert(token); err != nil {
		log.Printf("ERR: %v", err.Error())
		return err
	}
	return nil
}
