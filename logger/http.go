package logger

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LogMethodNotAllowed(r *http.Request) {
	Logger.WithFields(logrus.Fields{
		"Method": r.Method,
		"URI":    r.RequestURI,
		"Addr":   r.RemoteAddr,
		"Status": http.StatusMethodNotAllowed,
	}).Error()
}

func LogInternalServerError(r *http.Request, err error) {
	Logger.WithFields(logrus.Fields{
		"Method": r.Method,
		"URI":    r.RequestURI,
		"Addr":   r.RemoteAddr,
		"Status": http.StatusInternalServerError,
		"Error":  err,
	}).Error()
}
