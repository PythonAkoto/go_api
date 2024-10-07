package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi" // go web dev framework
	"github.com/pythonakoto/goapi/internal/handlers"
	log "github.com/sirupsen/logrus" // log = alias, used for debugging
)

func main() {

	log.SetReportCaller(true) // when printing we get file and line number, pass true to turn on

	// create a new router instance uing the chi library
	// this returns a pointer to a mux type -  thsi will be a struct that we use to set up the API
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go API service...")

	fmt.Println(`Go API..`)

	err := http.ListenAndServe("localhost:8000", r)
	// handle any errors when trying to serve localhost
	if err != nil {
		log.Error(err)
	}
}
