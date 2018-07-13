package control

import (
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(user model.User, session *mgo.Session) (*model.User, error) {
	var userMongo model.User
	if err := findUser(user, session, &userMongo); err != nil {
		return nil, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	ok := user.Username == userMongo.Username &&
		bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userMongo.Password)) == nil
	if !ok {
		return nil, fmt.Errorf("username or password invalid")
	}
	return &userMongo, nil
}

func findUser(user model.User, session *mgo.Session, userMongo *model.User) error {
	defer session.Close()

	collection := session.DB("jwt_authentication").C("user")
	if err := collection.Find(bson.M{"username": strings.ToLower(user.Username)}).
		One(userMongo); err != nil {
		return fmt.Errorf("username not found")
	}
	return nil

}
