package control

import (
	"gopkg.in/mgo.v2"
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"time"
	"log"
)

func invalidateToken(tokenUser string, session *mgo.Session) error {
	defer session.Close()

	collection := session.DB("jwt_authentication").C("expired")

	token := model.Token{Token: tokenUser, CreatedAt: time.Now()}
	if err := collection.Insert(token); err != nil {
		log.Printf("ERR: %v", err.Error())
		return err
	}
	return nil
}
