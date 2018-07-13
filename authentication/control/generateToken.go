package control

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/leandroandrade/jwt-authentication/rsa"
)

func generateToken(userUUID string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}

	tokenString, err := token.SignedString(rsa.PrivateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}
