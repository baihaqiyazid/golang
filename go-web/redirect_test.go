package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func Post(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Home")
}

func Pofile(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

func Google(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "https://www.google.com/", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/profile", Pofile)
	mux.HandleFunc("/home", Post)
	mux.HandleFunc("/link", Google)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}