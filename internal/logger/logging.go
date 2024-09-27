package logger

import (
	"net/http"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func LogHandler(request *http.Request) *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&nested.Formatter{})

	return logger.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": request.Method,
		"uri":    request.URL.String(),
		"ip":     request.RemoteAddr,
	})
}

func Logging() *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&nested.Formatter{
		ShowFullLevel: true,
	})

	return logger.WithFields(logrus.Fields{
		"at": time.Now().Format("2006-01-02 15:04:05"),
	})
}
