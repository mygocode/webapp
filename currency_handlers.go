package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Currency ...
type Currency struct {
	Shortform string `json:"shortform"`
	Fullform  string `json:"fullform"`
}

var currencyList []Currency

func getCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "currency" variable to json
	currencyListBytes, err := json.Marshal(currencyList)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON to the response
	w.Write(currencyListBytes)
}

func createCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Currency
	currency := Currency{}
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the currency from the form info
	currency.Shortform = r.Form.Get("shortform")
	currency.Fullform = r.Form.Get("fullform")

	//Check for blank data
	if currency.Shortform != "" && currency.Fullform != "" {
		currencyList = append(currencyList, currency)
	}

	//Finally, we redirect the user to the original HTML page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
