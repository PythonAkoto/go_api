package middleware

import (
	"errors"
	"net/http"

	"github.com/pythonakoto/goapi/api"
	"github.com/pythonakoto/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

// this needs to take in and return a http handler interface
func Authorization(next http.Handler) http.Handler {
	// takes a response writer & a pointer to a http request
	// we use a reponse writer to construct a response to the caller
	// the request contains all the information that is coming from incoming request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the username & auth token from the request parameter
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			// we log the error to the console & write the unauthorized error
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// if we have the correct info we can get the data from out mock database
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// this calls the next middleware in line
		// or the handler function for the endpoint if there is no mmiddleware left
		next.ServeHTTP(w, r)

	})
}
