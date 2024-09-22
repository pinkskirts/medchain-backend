package main

import (
	"log"
	"main/src/api/account"
	"main/src/api/login"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/Account", account.Account)
	apiRouter.HandleFunc("/login", login.Main)

	log.Println("Servico iniciado - porta 2000")
	http.ListenAndServe(":2000", r)
}
