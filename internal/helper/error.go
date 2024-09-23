package helper

import "github.com/sirupsen/logrus"

func IfPanicError(err error) {
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}

func PanicWithMessage(err error, message string) {
	if err != nil {
		logrus.Error(message)
		panic(message)
	}
}
