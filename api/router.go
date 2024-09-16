package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	utils "github.com/pinkskirts/medchain-backend/core/utils"
)

const (
	API_V1 = "/api/v1"
)

var (
	logger = utils.BuildStdoutLogger("Router")
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Middleware Alive!")
}

func InitRouter() {
	logger.Info("Initializing router")
	router := httprouter.New()

	router.GET("/", Index)

	logger.Fatal(http.ListenAndServe(":6969", router))
}
