package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request)  {
	
	if r.URL.Query().Get("name") != ""{
		http.ServeFile(w, r, "./recources/success.html")
	}else {
		http.ServeFile(w,r, "./recources/fail.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed recources/success.html
var recourcesSuccess string

//go:embed recources/fail.html
var recourcesFail string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request)  {
	
	if r.URL.Query().Get("name") != ""{
		fmt.Fprint(w, recourcesSuccess)
	}else {
		fmt.Fprint(w, recourcesFail)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}