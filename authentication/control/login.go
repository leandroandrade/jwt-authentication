package control

import (
	"net/http"
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"gopkg.in/mgo.v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"encoding/json"
	"strings"
)

func Login(requestUser model.User, session *mgo.Session) ([]byte, int) {
	user, err := authenticate(requestUser, session)
	if err != nil {
		return []byte(err.Error()), http.StatusUnauthorized
	}

	token, err := generateToken(user.UUID)
	if err != nil {
		return []byte(""), http.StatusInternalServerError
	}

	response, _ := json.Marshal(model.Token{Token: token})
	return response, http.StatusOK
}

func authenticate(user model.User, session *mgo.Session) (*model.User, error) {
	defer session.Close()

	var userMongo model.User

	collection := session.DB("jwt_authentication").C("user")
	if err := collection.Find(bson.M{"username": strings.ToLower(user.Username)}).
		One(&userMongo); err != nil {
		return nil, fmt.Errorf("username not found")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	ok := user.Username == userMongo.Username &&
		bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userMongo.Password)) == nil
	if !ok {
		return nil, fmt.Errorf("username or password invalid")
	}
	return &userMongo, nil
}
