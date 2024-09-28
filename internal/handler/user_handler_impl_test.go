package handler_test

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/config"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/entity"
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
	"github.com/sirupsen/logrus"
)

func setUpDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/portofolio_golang?parseTime=true")
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
	repository := repository.UserRepositoryImpl{}
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
		TokenExpiredAt:      0,
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

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody["data"])

}

func TestLoginFailed(t *testing.T) {
	db := setUpDB()
	truncateUser(db)

	tx, _ := db.Begin()
	repository := repository.UserRepositoryImpl{}
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
		TokenExpiredAt:      0,
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
	repository := repository.UserRepositoryImpl{}
	validate := validator.New()
	service := service.UserServiceImpl{DB: db, UserRepository: &repository, Validate: validate}
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
		TokenExpiredAt:      0,
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()
	serviceResponse := service.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-KEY", serviceResponse.Token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(response.StatusCode)
	fmt.Println(request.Header.Get("API-KEY"))

}

func TestGetByToken(t *testing.T) {
	db := setUpDB()
	truncateUser(db)
	tx, _ := db.Begin()
	repository := repository.UserRepositoryImpl{}
	validate := validator.New()
	service := service.UserServiceImpl{DB: db, UserRepository: &repository, Validate: validate}
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
		TokenExpiredAt:      0,
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()
	serviceResponse := service.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/profile", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-KEY", serviceResponse.Token)

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
	repository := repository.UserRepositoryImpl{}
	validate := validator.New()
	service := service.UserServiceImpl{DB: db, UserRepository: &repository, Validate: validate}
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
		TokenExpiredAt:      0,
	}
	repository.Insert(context.Background(), tx, user)
	tx.Commit()
	serviceResponse := service.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"username" : "riskiTaka", "name" : "Riski Store"}`)
	request := httptest.NewRequest(http.MethodPatch, "http://localhost:3000/api/users/profile", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-KEY", serviceResponse.Token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
