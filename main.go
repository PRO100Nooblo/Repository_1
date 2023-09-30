package main

import (
	"Project1/pkg/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	DB := config.Init()
	h := config.New(DB)

	r := mux.NewRouter()
	r.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
