package handler_test

import (
	"context"
	"database/sql"
	entity "ecommerce-cloning-app/entities"
	"ecommerce-cloning-app/internal/config"
	"ecommerce-cloning-app/internal/dto"
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
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

func setUpDB() *sql.DB {
	godotenv.Load("../../.env")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSource)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) *httprouter.Router {
	validate := validator.New()
	userRepository := repository.UserRepository{}
	userService := service.UserService{UserRepository: &userRepository, DB: db, Validate: validate}
	userHandler := handler.UserHandler{UserService: &userService}
	storeRepository := repository.StoreRepository{}
	storeService := service.StoreService{StoreRepository: &storeRepository, UserRepository: &userRepository, DB: db, Validate: validate}
	storeHandler := handler.StoreHandler{StoreService: &storeService}
	productRepository := repository.ProductRepository{}
	productService := service.ProductService{ProductRepository: &productRepository, StoreRepository: &storeRepository, UserRepository: &userRepository, DB: db, Validate: validate}
	productHandler := handler.ProductHandler{ProductService: &productService}
	middleware := middleware.AuthMiddleware{UserRepository: &userRepository, DB: db}

	router := config.NewRouter(&userHandler, &storeHandler, &productHandler, middleware)

	return router

}

func truncateUser(db *sql.DB) {
	db.Exec("DELETE FROM users")
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

func TestCreateUserFailed(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"no_telepon" : "", "password" : "Rahasia"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "BAD REQUEST", responseBody["status"])

	fmt.Println(responseBody["data"])
}

func TestLoginSuccess(t *testing.T) {
	db := setUpDB()
	truncateUser(db)

	tx, _ := db.Begin()
	repository := repository.UserRepository{}
	user := entity.User{
		NoTelepon:           "083156490686",
		Password:            helper.HashingPassword("rahasia"),
		Username:            helper.GeneratedUsername(),
		LastUpdatedUsername: helper.GeneratedTimeNow(),
		Name:                "",
		Email:               "",
		PhotoProfile:        "account_profile.png",
		Bio:                 "",
		Gender:              "",
		StatusMember:        "Basic",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"no_telepon" : "083156490686", "password" : "rahasia"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	cookie := response.Cookies()
	fmt.Println(cookie[0])
	fmt.Println(cookie[1])

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]any
	// json.Unmarshal(body, &responseBody)

	// assert.Equal(t, "OK", responseBody["status"])

	// fmt.Println(responseBody["data"])

}

func TestLoginFailed(t *testing.T) {
	db := setUpDB()
	truncateUser(db)

	tx, _ := db.Begin()
	repository := repository.UserRepository{}
	user := entity.User{
		NoTelepon:           "083156490686",
		Password:            helper.HashingPassword("rahasia"),
		Username:            helper.GeneratedUsername(),
		LastUpdatedUsername: helper.GeneratedTimeNow(),
		Name:                "",
		Email:               "",
		PhotoProfile:        "account_profile.png",
		Bio:                 "",
		Gender:              "",
		StatusMember:        "Basic",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"no_telepon" : "083156490686", "password" : "rahasia1"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

	fmt.Println(responseBody["data"])

}

func TestLogoutSuccess(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	tx, _ := db.Begin()
	repository := repository.UserRepository{}
	validate := validator.New()
	service := service.UserService{DB: db, UserRepository: &repository, Validate: validate}
	user := entity.User{
		NoTelepon:           "083156490686",
		Password:            helper.HashingPassword("rahasia"),
		Username:            helper.GeneratedUsername(),
		LastUpdatedUsername: helper.GeneratedTimeNow(),
		Name:                "",
		Email:               "",
		PhotoProfile:        "account_profile.png",
		Bio:                 "",
		Gender:              "",
		StatusMember:        "Basic",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()
	userLogin, token := service.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", userLogin.AccessToken)
	request.AddCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
	})

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(response.StatusCode)
	// cookie := response.Cookies()
	// fmt.Println(cookie[0])

}

func TestFindUser(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	tx, _ := db.Begin()
	repository := repository.UserRepository{}
	validate := validator.New()
	service := service.UserService{DB: db, UserRepository: &repository, Validate: validate}
	user := entity.User{
		NoTelepon:           "083156490686",
		Password:            helper.HashingPassword("rahasia"),
		Username:            helper.GeneratedUsername(),
		LastUpdatedUsername: helper.GeneratedTimeNow(),
		Name:                "",
		Email:               "",
		PhotoProfile:        "account_profile.png",
		Bio:                 "",
		Gender:              "",
		StatusMember:        "Basic",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()
	responseLogin, token := service.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/profile", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", responseLogin.AccessToken)
	request.AddCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
	})

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, user.Username, responseBody["data"].(map[string]interface{})["username"])

	logrus.Info(responseBody["data"])
}

func TestUpdateprofileFailed(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	tx, _ := db.Begin()
	repository := repository.UserRepository{}
	// validate := validator.New()
	// service := service.UserService{DB: db, UserRepository: &repository, Validate: validate}
	user := entity.User{
		NoTelepon:           "083156490686",
		Password:            helper.HashingPassword("rahasia"),
		Username:            helper.GeneratedUsername(),
		LastUpdatedUsername: helper.GeneratedTimeNow(),
		Name:                "",
		Email:               "",
		PhotoProfile:        "account_profile.png",
		Bio:                 "",
		Gender:              "",
		StatusMember:        "Basic",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()
	// serviceResponse := service.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"username" : "riskiTaka", "name" : "Riski Store"}`)
	request := httptest.NewRequest(http.MethodPatch, "http://localhost:3000/api/users/profile", requestBody)
	request.Header.Add("Content-Type", "application/json")
	// request.Header.Add("API-KEY", serviceResponse.Token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
