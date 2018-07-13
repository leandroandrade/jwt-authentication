package control

import (
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"github.com/leandroandrade/jwt-authentication/rsa"
)

func extractToken(req *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			return rsa.PublicKey, nil
		}
	})

	return token, err
}
