package control

import (
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/leandroandrade/jwt-authentication/rsa"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/leandroandrade/jwt-authentication/mongo"
)

func RequireToken(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			return rsa.PublicKey, nil
		}
	})

	if err == nil && token.Valid && !isLogout(mongo.New().Copy(), token.Raw) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(err.Error()))
	}
}

func isLogout(session *mgo.Session, token string) bool {
	defer session.Close()

	collection := session.DB("jwt_authentication").C("expired")
	result, _ := collection.Find(bson.M{"token": token}).Count()
	return result > 0
}
