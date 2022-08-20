package goweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func FormUpload(w http.ResponseWriter, r *http.Request){
	t.ExecuteTemplate(w, "form_upload.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request){
	// request.ParseMultipartForm(32 << 20)
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./recources/images/" + fileHeader.Filename) 
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	
	t.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name" : name,
		"File" : "/static/" + fileHeader.Filename,
	})
}

func TestFormUpload(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", FormUpload)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./recources/images/"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed recources/images/success.svg
var testUploadFile []byte 

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Muhammad Yazid Baihaqi")
	file, _ := writer.CreateFormFile("file", "TestFileUpload.svg")
	file.Write(testUploadFile)
	writer.Close()

	request := httptest.NewRequest("POST", "locahost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
