package handler

import (
	"net/http"

	"github.com/rs/cors"
)

func GetCorsHandler(handler http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Content-Length"},
		MaxAge:           60,
	}).Handler(handler)
}
