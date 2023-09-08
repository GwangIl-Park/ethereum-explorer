package handler

import (
	"ethereum-explorer/logger"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func GetLoggerHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		duration := time.Since(start)

		logger.Logger.WithFields(logrus.Fields{
			"Method":   r.Method,
			"URI":      r.RequestURI,
			"Addr":     r.RemoteAddr,
			"Duration": duration,
		}).Info("Response")
	})
}
