package webrequest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendSuccessfulRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"value":"fixed"}`))
	}))
	defer server.Close()

	value, err := SendRequest(server.URL)
	bytes := []byte(`{"value":"fixed"}`)

	if string(value) != string(bytes) {
		t.Fatalf(`SendRequest(%s) == %q, expected %q`, server.URL, value, bytes)
	}

	if err != nil {
		t.Fatalf(`Successful SendRequest(%s) returned an error %s`, server.URL, err)
	}
}

func TestSendBadRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()
	value, err := SendRequest(server.URL)

	if len(value) > 0 {
		t.Fatal(`Expected response of nil for bad request (400) response code.`)
	}

	if err == nil {
		t.Fatal("Error should be thrown")
	}
}
