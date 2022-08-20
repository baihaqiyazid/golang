package goweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	dir := http.Dir("./recources")
	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/home/", http.StripPrefix("/home", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed recources 
var recources embed.FS

func TestFileServerEmbed(t *testing.T) {

	dir, _ := fs.Sub(recources, "recources")
	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/home/", http.StripPrefix("/home", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
