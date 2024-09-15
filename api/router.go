package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Middleware Alive!")
}

func InitRouter() {
	router := httprouter.New()

	router.GET("/", Hello)

	log.Fatal(http.ListenAndServe(":6969", router))
}
