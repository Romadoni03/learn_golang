package handler_test

import (
	"database/sql"
	"ecommerce-cloning-app/internal/config"
	"ecommerce-cloning-app/internal/handler"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/middleware"
	"ecommerce-cloning-app/internal/repository"
	"ecommerce-cloning-app/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/magiconair/properties/assert"
)

func setUpDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/portofolio_golang")
	helper.IfPanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) *httprouter.Router {
	validate := validator.New()
	userRepository := repository.UserRepositoryImpl{}
	userService := service.UserServiceImpl{UserRepository: &userRepository, DB: db, Validate: validate}
	userHandler := handler.UserHandlerImpl{UserService: &userService}
	middleware := middleware.AuthMiddleware{UserRepository: &userRepository, DB: db}

	router := config.NewRouter(&userHandler, middleware)

	return router

}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateUserSuccess(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"no_telepon" : "082332271835", "password" : "Rahasia"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody["data"])
}

func TestLoginSuccess(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"no_telepon" : "123123125", "password" : "rahasia"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody["data"])

}

func TestLogout(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-KEY", "e7e4551a-2f1e-4062-8f57-af012eea10c1")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	// assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(response.StatusCode)
	fmt.Println(request.Header.Get("API-KEY"))

}
