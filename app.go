package main

import (
	"log"

	"main/src/api/login"
	"main/src/api/pdf"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/login", login.Main)
	apiRouter.HandleFunc("/pdf", pdf.Main)

	log.Println("Servico iniciado - porta 2000")
	http.ListenAndServe(":2000", r)
}
