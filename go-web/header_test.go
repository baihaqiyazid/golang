package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request)  {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType) 
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("POST", "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("X-Powered-By", "Baihaqi")
	fmt.Fprint(w, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("POST", "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	fmt.Println(response.Header.Get("X-Powered-By"))

}