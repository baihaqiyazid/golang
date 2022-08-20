package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)



//go:embed templates/*.gohtml
var templatesNew embed.FS

var t = template.Must(template.ParseFS(templatesNew, "templates/*.gohtml"))

type Profile struct {
	Title string
	Name string
}

//GLOBAL FUNCTIONS
func (profile Profile) Upper(value string) string {
	return strings.ToUpper(value)
}

func (profile Profile) SayHello(name string) string {
	return "Hello " + profile.Upper(name) + " My Name is "+ profile.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "function.gohtml", Profile{
		Title: "Profile",
		Name : "Yazid",
	})
}

//TEST
func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}