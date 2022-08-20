package main

import (
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, i)
		w.Write([]byte("500 - Something bad happened!"))
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// fmt.Fprint(w, "Hello World!!")
		panic("Error ")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
