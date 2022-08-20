package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request)  {
	
	name := r.URL.Query().Get("name")

	if name == ""{
		fmt.Fprint(w, "Hello")
	}else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParams(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/home?name=Yazid", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParams(w http.ResponseWriter, r *http.Request)  {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestMultipleParams(t *testing.T) {

	request := httptest.NewRequest("GET", "http://localhost:8080/home?firstname=Yazid&lastname=Baihaqi", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request)  {
	
	query := r.URL.Query()
	names := query["name"]

	fmt.Fprint(w, "Hello " + strings.Join(names, " "))
}

func TestMultipleParamaterValue(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/home?name=Muhammad&name=Yazid&name=Baihaqi", nil)	
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body) 

	fmt.Println(string(body))
}