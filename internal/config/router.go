package config

import (
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/handler"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userHandler handler.UserHandler) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", userHandler.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}
