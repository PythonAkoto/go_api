package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/pythonakoto/goapi/api"
	"github.com/pythonakoto/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

// craete a repomnse writer & HHTP response pointer for the handler function
// we assume this has passed the authorization middleware, so we just need to get the usernames from the parameters

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// get the params from the URL and set them to the values in the struct
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins((params.Username))
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
