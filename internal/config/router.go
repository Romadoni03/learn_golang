package config

import (
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/handler"
	"ecommerce-cloning-app/internal/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userHandler *handler.UserHandler, storeHandler *handler.StoreHandler, productHandler *handler.ProductHandler, middleware middleware.AuthMiddleware) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", userHandler.Create)
	router.POST("/api/users/login", userHandler.Login)
	router.DELETE("/api/users/logout", middleware.AuthMiddleware(userHandler.Logout))
	router.GET("/api/users/profile", middleware.AuthMiddleware(userHandler.GetByToken))
	router.PATCH("/api/users/profile", middleware.AuthMiddleware(userHandler.Update))
	router.POST("/api/stores", middleware.AuthMiddleware(storeHandler.Create))
	router.DELETE("/api/stores", middleware.AuthMiddleware(storeHandler.Delete))
	router.GET("/api/stores", middleware.AuthMiddleware(storeHandler.FindByUser))
	router.POST("/api/stores/products", middleware.AuthMiddleware(productHandler.Create))
	router.GET("/api/stores/products", middleware.AuthMiddleware(productHandler.FindAll))

	router.PanicHandler = exception.ErrorHandler

	return router
}
