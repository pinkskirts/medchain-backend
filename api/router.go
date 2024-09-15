package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	utils "github.com/pinkskirts/medchain-backend/core/utils"
)

var (
	logger = utils.BuildStdoutLogger("Router")
)

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Middleware Alive!")
}

func InitRouter() {
	logger.Info("Initializing router")
	router := httprouter.New()

	router.GET("/", Hello)

	logger.Fatal(http.ListenAndServe(":6969", router))
}
