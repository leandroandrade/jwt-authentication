package main

import (
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/codegangsta/negroni"
	"github.com/leandroandrade/jwt-authentication/authentication/control"
	"github.com/leandroandrade/jwt-authentication/authentication/boundary"
	"github.com/leandroandrade/jwt-authentication/mongo"
)

func main() {
	handler := boundary.NewHandler(mongo.New())

	router := mux.NewRouter().StrictSlash(true)

	resources := router.PathPrefix("/resources").Subrouter()

	login := resources.PathPrefix("/login").Subrouter()
	login.Methods("POST").HandlerFunc(handler.Login)

	resources.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(control.RequireToken),
			negroni.HandlerFunc(handler.RefreshToken),
		)).Methods("GET")

	resources.Handle("/hello",
		negroni.New(
			negroni.HandlerFunc(control.RequireToken),
			negroni.HandlerFunc(handler.Hello),
		)).Methods("GET")

	resources.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(control.RequireToken),
			negroni.HandlerFunc(handler.Logout),
		)).Methods("GET")

	negr := negroni.Classic()
	negr.Use(gzip.Gzip(gzip.BestSpeed))

	negr.UseHandler(router)
	negr.Run(":3000")
}
