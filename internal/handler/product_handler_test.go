package handler_test

import (
	"bytes"
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
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/magiconair/properties/assert"
)

func truncateProducts(db *sql.DB) {
	db.Exec("DELETE FROM products")
}

func TestAddProductSuccess(t *testing.T) {
	db := setUpDB()
	truncateProducts(db)
	truncateStores(db)
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.UserRepository{}
	validate := validator.New()
	userService := service.UserService{DB: db, UserRepository: &userRepository, Validate: validate}
	storeRepository := repository.StoreRepository{}
	storeService := service.StoreService{StoreRepository: &storeRepository, UserRepository: &userRepository, DB: db, Validate: validate}
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
	userRepository.Insert(context.Background(), tx, user)
	tx.Commit()
	serviceResponse := userService.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})
	storeService.Create(context.Background(), dto.StoreCreateRequest{Name: "Riski Store"}, serviceResponse.Token)
	router := setupRouter(db)

	product := dto.ProductCreateUpdateRequest{
		PhotoProduct:      "test_foto.jpg",
		Name:              "Mouse",
		Category:          "Elektronik",
		Description:       "goood",
		DangeriousProduct: "no danger",
		Price:             100000,
		Stock:             10,
		Wholesaler:        "ga tau",
		ShippingCost:      2000,
		ShippingInsurance: 10,
		Conditions:        "new",
		PreOrder:          "no",
		Status:            "ready",
	}
	data, _ := json.Marshal(product)
	requestBody := bytes.NewReader(data)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/stores/products", requestBody)
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

func TestGetAllProductSuccess(t *testing.T) {
	db := setUpDB()
	truncateProducts(db)
	truncateStores(db)
	truncateUser(db)
	tx, _ := db.Begin()

	userRepository := repository.UserRepository{}
	validate := validator.New()
	userService := service.UserService{DB: db, UserRepository: &userRepository, Validate: validate}
	storeRepository := repository.StoreRepository{}
	storeService := service.StoreService{StoreRepository: &storeRepository, UserRepository: &userRepository, DB: db, Validate: validate}
	productRepository := repository.ProductRepository{}
	productService := service.ProductService{ProductRepository: &productRepository, StoreRepository: &storeRepository, UserRepository: &userRepository, DB: db, Validate: validate}

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
		TokenExpiredAt:      0,
	}
	userRepository.Insert(context.Background(), tx, user)
	tx.Commit()
	serviceResponse := userService.Login(context.Background(), dto.UserCreateRequest{NoTelepon: user.NoTelepon, Password: "rahasia"})
	storeService.Create(context.Background(), dto.StoreCreateRequest{Name: "Riski Store"}, serviceResponse.Token)

	product := dto.ProductCreateUpdateRequest{
		PhotoProduct:      "test_foto.jpg",
		Name:              "Mouse",
		Category:          "Elektronik",
		Description:       "goood",
		DangeriousProduct: "no danger",
		Price:             100000,
		Stock:             10,
		Wholesaler:        "ga tau",
		ShippingCost:      2000,
		ShippingInsurance: 10,
		Conditions:        "new",
		PreOrder:          "no",
		Status:            "ready",
	}
	product2 := dto.ProductCreateUpdateRequest{
		PhotoProduct:      "test_foto_aja.jpg",
		Name:              "keyboard",
		Category:          "computer",
		Description:       "bad",
		DangeriousProduct: "danger",
		Price:             500000,
		Stock:             100,
		Wholesaler:        "riski",
		ShippingCost:      5000,
		ShippingInsurance: 100,
		Conditions:        "used",
		PreOrder:          "yes",
		Status:            "ready banyak",
	}
	productService.Create(context.Background(), product, serviceResponse.Token)
	productService.Create(context.Background(), product2, serviceResponse.Token)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/stores/products", nil)
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
