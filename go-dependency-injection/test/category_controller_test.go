package test

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "fmt"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/model/entity"
	"go-restful-api/repository"
	"go-restful-api/services"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupDB() *sql.DB  {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go-restful-test")
	helper.Panic(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(20 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategories(db *sql.DB)  {
	db.Exec("TRUNCATE categories")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupDB()
	truncateCategories(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name" : "Laptop" }`)
	request := httptest.NewRequest("POST", "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "Laptop", responseBody["data"].(map[string]interface{})["name"])

	// fmt.Println(string(body))
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupDB()
	truncateCategories(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name" : "" }`)
	request := httptest.NewRequest("POST", "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

	// fmt.Println(string(body))
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupDB()
	truncateCategories(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Laptop",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name" : "Laptop" }`)
	request := httptest.NewRequest("PUT", "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Laptop", responseBody["data"].(map[string]interface{})["name"])

	// fmt.Println(string(body))
}


func TestUpdateCategoryFailed(t *testing.T) {
	db := setupDB()
	truncateCategories(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Laptop",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{ "name" : "" }`)
	request := httptest.NewRequest("PUT", "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

	// fmt.Println(string(body))
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupDB()
	truncateCategories(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Laptop",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest("GET", "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])

	// fmt.Println(string(body))
}

func TestGetCategoryFailed(t *testing.T) {

	db := setupDB()
	truncateCategories(db)

	router := setupRouter(db)

	request := httptest.NewRequest("GET", "http://localhost:3000/api/categories/10", nil)
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

	// fmt.Println(string(body))
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupDB()
	truncateCategories(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Laptop",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest("DELETE", "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])

	// fmt.Println(string(body))
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupDB()
	truncateCategories(db)
	router := setupRouter(db)

	request := httptest.NewRequest("DELETE", "http://localhost:3000/api/categories/10", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

	// fmt.Println(string(body))
}

func TestListCategorySuccess(t *testing.T) {
	db := setupDB()
	truncateCategories(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category1 := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Laptop",
	})
	category2 := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Mouse",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest("GET", "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])

	var categories = responseBody["data"].([]interface{})
	categories1 := categories[0].(map[string]interface{})
	categories2 := categories[1].(map[string]interface{})  

	assert.Equal(t, category1.Id, int(categories1["id"].(float64)))
	assert.Equal(t, category1.Name, categories1["name"])

	assert.Equal(t, category2.Id, int(categories2["id"].(float64)))
	assert.Equal(t, category2.Name, categories2["name"])

	// fmt.Println(string(body))
}

func TestUnauthorized(t *testing.T) {
	db := setupDB()
	truncateCategories(db)

	router := setupRouter(db)

	request := httptest.NewRequest("GET", "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

	// fmt.Println(string(body))
}