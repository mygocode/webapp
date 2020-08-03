package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestGetCurrencysHandler(t *testing.T) {

	currencyList = nil
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	currencyHandler := http.HandlerFunc(getCurrencyHandler)
	currencyHandler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(currencyList) == 0 {
		logrus.Info("No data in currency list.")
	} else {

		expected := Currency{"PKR", "Pakistani Rupee"}
		// currencyArray := []Currency{}
		// err = json.NewDecoder(responseRecorder.Body).Decode(&currencyArray)

		// if err != nil {
		// 	t.Fatal(err)
		// }

		actual := currencyList[0]

		if actual != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
		}
	}
}

func TestCreateCurrencyHandler(t *testing.T) {

	currencyList = nil
	form := newCreateCurrencyForm()
	request, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createCurrencyHandler)
	hf.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Get the latest entry from currencyList object defined in currency_handler.go
	currencyData := currencyList[len(currencyList)-1]
	expected := Currency{"PKR", "Pakistani Rupee"}

	fmt.Println(currencyList)

	if currencyData != expected {
		t.Errorf("Handler returned wrong body: Got %v , Expected %v", currencyData, expected)
	}
}

func TestBlankCurrencyHandler(t *testing.T) {

	currencyList = nil
	form := newBlankCurrencyForm()
	request, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createCurrencyHandler)
	hf.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Get the latest entry from currencyList object defined in currency_handler.go

	currencyData := Currency{"", ""}
	if len(currencyList) > 0 {
		currencyData = currencyList[0]
	}

	expected := Currency{"", ""}

	if currencyData != expected {
		t.Errorf("Handler returned wrong body: Got %v , Expected %v", currencyData, expected)
	}
}

func newCreateCurrencyForm() *url.Values {
	form := url.Values{}
	form.Set("shortform", "PKR")
	form.Set("fullform", "Pakistani Rupee")
	return &form
}

func newBlankCurrencyForm() *url.Values {
	form := url.Values{}
	form.Set("shortform", "")
	form.Set("fullform", "")
	return &form
}
