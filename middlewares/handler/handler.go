package handler

import "net/http"

func GetHandler(router http.Handler) http.Handler {
	return GetCorsHandler(router)
}
