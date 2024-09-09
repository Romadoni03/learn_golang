package main

import (
	"ecommerce-cloning-app/internal/config"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/handler"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/middleware"
	"ecommerce-cloning-app/internal/repository"
	"ecommerce-cloning-app/internal/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db, _ := config.NewDB()
	validate := validator.New()
	userRepository := repository.UserRepositoryImpl{}
	userService := service.UserServiceImpl{UserRepository: &userRepository, DB: db, Validate: validate}
	userHandler := handler.UserHandlerImpl{UserService: &userService}
	middleware := middleware.AuthMiddleware{UserRepository: &userRepository, DB: db}

	router := httprouter.New()
	router.POST("/api/users", userHandler.Create)
	router.POST("/api/users/login", userHandler.Login)
	router.DELETE("/api/users/logout", middleware.AuthMiddleware(userHandler.Logout))

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	fmt.Println("server is running")
	helper.IfPanicError(err)
}
