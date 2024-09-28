package helper

import (
	"ecommerce-cloning-app/internal/logger"
)

func IfPanicError(err error) {
	if err != nil {
		logger.Logging().Error(err)
		panic(err)
	}
}

func PanicWithMessage(err error, message string) {
	if err != nil {
		logger.Logging().Error(message)
		panic(err)
	}
}
