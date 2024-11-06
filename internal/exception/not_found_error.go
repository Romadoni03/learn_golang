package exception

import "ecommerce-cloning-app/internal/logger"

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func PanicNotFoundError(err error, msg string) {
	if err != nil {
		logger.Logging().Error(err)
		panic(NewNotFoundError(msg))
	}
}
