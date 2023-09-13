package handler

import (
	"ethereum-explorer/logger"
	"fmt"
	"net/http"
	"time"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	body           string
	statusCode     int
}

func (rw *ResponseWriterWrapper) Write(data []byte) (int, error) {
	rw.body = string(data[:])
	return rw.ResponseWriter.Write(data)
}

func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func GetLoggerHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestLog := fmt.Sprintf("REQ > [%s] %s", r.Method, r.URL.Path)
		logger.Logger.Info(requestLog)
		rw := &ResponseWriterWrapper{
			ResponseWriter: w,
		}

		handler.ServeHTTP(rw, r)

		duration := time.Since(start)

		responseLog := fmt.Sprintf("RES > %d [%s] %s %d", rw.statusCode, r.Method, r.URL.Path, duration.Microseconds())
		logger.Logger.Info(responseLog)
	})
}
