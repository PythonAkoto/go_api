package tools

import (
	log "github.com/sirupsen/logrus"
)

// define the types of data the database will return

// Database collections
type LoginDetails struct {
	AuthToken string
	UserName  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

// create database interface that will have the functions that will ineteract with the API
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// returns the database interface
func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil

}
