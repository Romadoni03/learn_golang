package exception

import "ecommerce-cloning-app/internal/logger"

type UnauthorizedError struct {
	Error string
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{Error: error}
}

func PanicUnauthorizedError(err error, msg string) {
	if err != nil {
		logger.Logging().Error(err)
		panic(NewUnauthorizedError(msg))
	}
}
