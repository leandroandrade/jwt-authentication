package boundary

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"github.com/leandroandrade/jwt-authentication/authentication/control"
	"gopkg.in/mgo.v2"
)

type Handler struct {
	mongo *mgo.Session
}

func NewHandler(m *mgo.Session) *Handler {
	return &Handler{
		mongo: m,
	}
}

func (h Handler) Login(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		log.Println("ERR: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	token, responseStatus := control.Login(user, h.mongo.Copy())
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(responseStatus)
	writer.Write(token)

}

func (h Handler) RefreshToken(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	var user model.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		log.Println("ERR: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(control.Refresh(user))
}

func (h Handler) Logout(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	err := control.Logout(request, h.mongo.Copy())
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusNoContent)
	}
}

func (h Handler) Hello(writer http.ResponseWriter, _ *http.Request, _ http.HandlerFunc) {
	writer.Write([]byte("Hello, World!"))
}
