package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server {
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandler(t *testing.T) {
	
	var handler http.HandlerFunc = func (w http.ResponseWriter, r *http.Request) {
		
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}

func TestServeMux(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home")
	})

	mux.HandleFunc("/post/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "All Post")
	})

	mux.HandleFunc("/post/learning-go", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Learning Golang")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHelloHandler(t *testing.T) {
	
	request := httptest.NewRequest("GET", "http:://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}