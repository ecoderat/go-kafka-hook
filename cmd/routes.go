package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/articles", app.Index).Methods("GET")
	r.HandleFunc("/articles/new", app.Create).Methods("POST")
	r.HandleFunc("/articles/update", app.Update).Methods("PUT")
	r.HandleFunc("/articles/delete", app.Delete).Methods("DELETE")

	return r
}
