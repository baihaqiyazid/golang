package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/home.gohtml"))
	t.ExecuteTemplate(w, "home.gohtml", map[string]interface{}{
		"Title": "Home",
		"Name" : "Yazid",
		"Address" : map[string]interface{}{
			"Street" : "Jl. Potlot",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest("GET","http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

type Address struct{
	Street string
}

type Home struct{
	Title, Name string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/home.gohtml"))
	t.ExecuteTemplate(w, "home.gohtml", Home{
		Title: "Home",
		Name: "Yazid",
		Address : Address{
			Street : "Jl. Potlot",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET","http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}