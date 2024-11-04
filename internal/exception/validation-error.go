package exception

import "ecommerce-cloning-app/internal/logger"

type ValidationError struct {
	Error string
}

func NewValidationError(error string) ValidationError {
	return ValidationError{Error: error}
}

func PanicValidationError(err error, msg string) {
	if err != nil {
		logger.Logging().Error(err)
		panic(NewValidationError(msg))
	}
}
