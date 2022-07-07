package main

import (
	"net/http"
	"test-efishery/db"
	"test-efishery/handler"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDatabase()
	db.Migration()

	r := mux.NewRouter()
	r.HandleFunc("/", handler.CreateForm).Methods("POST")
	r.HandleFunc("/", handler.GetForm).Methods("GET")
	r.HandleFunc("/{id}", handler.UpdateForm).Methods("PUT")
	r.HandleFunc("/{id}", handler.DeleteForm).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
