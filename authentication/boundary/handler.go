package boundary

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/leandroandrade/jwt-authentication/authentication/model"
	"github.com/leandroandrade/jwt-authentication/authentication/control"
	"github.com/leandroandrade/jwt-authentication/mongo"
)

type Handler struct {
	mongo *mongo.MongoDatabase
}

func NewHandler(db *mongo.MongoDatabase) *Handler {
	return &Handler{
		mongo: db,
	}
}

func (h Handler) Login(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		log.Println("ERR: ", err)
		return
	}

	token, responseStatus := control.Login(user, h.mongo.Get())
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(responseStatus)
	writer.Write(token)

}

func (h Handler) RefreshToken(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	token, responseStatus := control.Refresh(request, h.mongo.Get())
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(responseStatus)
	writer.Write(token)
}

func (h Handler) Logout(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	err := control.Logout(request, h.mongo.Get())
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
