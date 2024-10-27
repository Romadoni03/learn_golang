package handler_test

import (
	"context"
	"database/sql"
	entity "ecommerce-cloning-app/entities"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/repository"
	"ecommerce-cloning-app/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/magiconair/properties/assert"
)

func truncateStores(db *sql.DB) {
	db.Exec("DELETE FROM stores")
}

func TestCreateStoreSuccess(t *testing.T) {
	db := setUpDB()
	truncateStores(db)
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
	serviceResponse, _ := service.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "riski store"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/stores", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-KEY", serviceResponse.Message)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteStoreSuccess(t *testing.T) {
	db := setUpDB()
	truncateStores(db)
	truncateUser(db)
	tx, _ := db.Begin()
	validate := validator.New()
	userRepository := repository.UserRepository{}
	storeRepository := repository.StoreRepository{}
	userService := service.UserService{DB: db, UserRepository: &userRepository, Validate: validate}
	storeService := service.StoreService{DB: db, UserRepository: &userRepository, StoreRepository: &storeRepository, Validate: validate}
	user := entity.User{
		NoTelepon:           "082332271835",
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
	userRepository.Insert(context.Background(), tx, user)
	tx.Commit()
	serviceResponse, _ := userService.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})
	storeService.Create(context.Background(), dto.StoreCreateRequest{Name: "riski_taka_store"}, serviceResponse.Message)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/stores", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-KEY", serviceResponse.Message)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Success Delete Store", responseBody["data"].(map[string]interface{})["message"])
	fmt.Println(responseBody["data"])
}

func TestFindByUser(t *testing.T) {
	db := setUpDB()
	truncateStores(db)
	truncateUser(db)
	tx, _ := db.Begin()
	validate := validator.New()
	userRepository := repository.UserRepository{}
	storeRepository := repository.StoreRepository{}
	userService := service.UserService{DB: db, UserRepository: &userRepository, Validate: validate}
	storeService := service.StoreService{DB: db, UserRepository: &userRepository, StoreRepository: &storeRepository, Validate: validate}
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
	userRepository.Insert(context.Background(), tx, user)
	tx.Commit()
	serviceResponse, _ := userService.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})
	storeService.Create(context.Background(), dto.StoreCreateRequest{Name: "riski_taka_store"}, serviceResponse.Message)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/stores", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-KEY", serviceResponse.Message)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "riski_taka_store", responseBody["data"].(map[string]interface{})["name"])
	fmt.Println(responseBody["data"])
}
