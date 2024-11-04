package exception

import "ecommerce-cloning-app/internal/logger"

type InternalServerError struct {
	Error string
}

func NewInternalServerError(error string) InternalServerError {
	return InternalServerError{Error: error}
}

func PanicInternalServerError(err error, msg string) {
	if err != nil {
		logger.Logging().Error(msg)
		panic(NewInternalServerError(msg))
	}
}
