package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed recources/*.txt
var recources embed.FS

func TestRouterServeFile(t *testing.T) {
	router := httprouter.New()
	
	dir, _ := fs.Sub(recources, "recources")
	
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Router", string(body))
	fmt.Println(string(body))
}

func TestRouterServeFile2(t *testing.T) {
	router := httprouter.New()
	
	dir, _ := fs.Sub(recources, "recources")
	
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Good Bye router", string(body))
	fmt.Println(string(body))
}