package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

func TestRouter(t *testing.T) {

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

func TestStaticFileServer(t *testing.T) {

	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}
