package main

import (
	"ecommerce-cloning-app/internal/config"
	"ecommerce-cloning-app/internal/handler"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/middleware"
	"ecommerce-cloning-app/internal/repository"
	"ecommerce-cloning-app/internal/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := config.NewDB()
	validate := validator.New()
	userRepository := repository.UserRepositoryImpl{}
	userService := service.UserServiceImpl{UserRepository: &userRepository, DB: db, Validate: validate}
	userHandler := handler.UserHandlerImpl{UserService: &userService}
	middleware := middleware.AuthMiddleware{UserRepository: &userRepository, DB: db}

	router := config.NewRouter(&userHandler, middleware)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	fmt.Println("server is running")
	helper.IfPanicError(err)
}
