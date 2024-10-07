package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		UserName:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		UserName:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		UserName:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "Jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

// for mockDB to conform to our database interface

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
