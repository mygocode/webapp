package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	router := newRouter()
	mockServer := httptest.NewServer(router)
	response, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Stataus returned is NOT OK")
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	responseString := string(body)
	expected := "Hello World!"

	if responseString != expected {
		t.Errorf("Expected response should be %s, got %s", expected, responseString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	response, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", response.StatusCode)
	}

	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	responseString := string(b)
	expected := ""

	if responseString != expected {
		t.Errorf("Response should be %s, got %s", expected, responseString)
	}

}
