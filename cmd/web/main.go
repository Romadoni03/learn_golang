package main

import (
	"ecommerce-cloning-app/internal/config"
	"ecommerce-cloning-app/internal/handler"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/middleware"
	"ecommerce-cloning-app/internal/repository"
	"ecommerce-cloning-app/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := config.NewDB()
	validate := validator.New()
	//user
	userRepository := repository.UserRepositoryImpl{}
	userService := service.UserServiceImpl{UserRepository: &userRepository, DB: db, Validate: validate}
	userHandler := handler.UserHandlerImpl{UserService: &userService}
	//store
	storeRepository := repository.StoreRepositoryImpl{}
	storeService := service.StoreServiceImpl{StoreRepository: &storeRepository, UserRepository: &userRepository, DB: db, Validate: validate}
	storeHandler := handler.StoreHandlerImpl{StoreService: &storeService}
	//middleware
	middleware := middleware.AuthMiddleware{UserRepository: &userRepository, DB: db}

	router := config.NewRouter(&userHandler, &storeHandler, middleware)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	logger.Logging().Info("Server Is Running")
	err := server.ListenAndServe()
	helper.IfPanicError(err)
}
