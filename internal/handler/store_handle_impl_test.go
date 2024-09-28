package handler_test

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/repository"
	"ecommerce-cloning-app/internal/service"
	"encoding/json"
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

	requestBody := strings.NewReader(`{"name" : "riski store"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/stores", requestBody)
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
}
