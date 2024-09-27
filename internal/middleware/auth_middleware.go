package middleware

import (
	"database/sql"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/repository"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func (middleware *AuthMiddleware) AuthMiddleware(handler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		logger.LogHandler(request).Info("Incoming Request")
		token := request.Header.Get("API-KEY")
		if token == "" {
			logger.LogHandler(request).Error("UNAUTHORIZED")
			panic(exception.NewUnauthorizedError("UNAUTHORIZED"))
		}
		tx, err := middleware.DB.Begin()
		helper.IfPanicError(err)
		defer tx.Commit()
		user, _ := middleware.UserRepository.FindFirstByToken(request.Context(), tx, token)

		if user.Token != token && user.TokenExpiredAt < time.Now().UnixMilli() {
			logger.LogHandler(request).Error("UNAUTHORIZED")
			panic(exception.NewUnauthorizedError("UNAUTHORIZED"))
		}

		handler(writer, request, params)
	}
}
