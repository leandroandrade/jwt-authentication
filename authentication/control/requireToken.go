package control

import (
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/leandroandrade/jwt-authentication/mongo"
	"errors"
)

func RequireToken(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	token, err := extractToken(req)
	if err != nil || !token.Valid {
		writeError(rw, err)
		return
	}

	if err = isLogout(mongo.New().Get(), token.Raw); err != nil {
		writeError(rw, err)
		return
	}

	next(rw, req)
}

func writeError(rw http.ResponseWriter, err error) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusUnauthorized)
	rw.Write([]byte(fmt.Sprintf(`{"error" : "%v"}`, err.Error())))
}

func isLogout(session *mgo.Session, token string) error {
	defer session.Close()

	collection := session.DB("jwt_authentication").C("expired")
	result, _ := collection.Find(bson.M{"token": token}).Count()
	if result > 0 {
		return errors.New("Token is expired")
	}
	return nil
}
