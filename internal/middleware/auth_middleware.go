package middleware

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/auth"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/repository"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	UserRepository *repository.UserRepository
	DB             *sql.DB
}

func (middleware *AuthMiddleware) AuthMiddleware(handler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		logger.LogHandler(request).Info("Incoming Request")
		accessToken := request.Header.Get("token")
		if accessToken == "" {
			logger.LogHandler(request).Error("UNAUTHORIZED")
			panic(exception.NewUnauthorizedError("UNAUTHORIZED"))
		}

		claims, errValidate := auth.ValidateJWT(accessToken)
		logger.Logging().Info("Phone :", claims.Phone)
		logger.Logging().Info("Access Token :", accessToken)
		logger.Logging().Info("ExpiredAt :", claims.ExpiresAt.Time)
		if errValidate != nil {
			logger.LogHandler(request).Error("UNAUTHORIZED")
			panic(exception.NewUnauthorizedError("UNAUTHORIZED"))
		}

		request = request.WithContext(context.WithValue(request.Context(), "phone", claims.Phone))
		handler(writer, request, params)
	}
}
