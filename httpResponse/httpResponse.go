package httpResponse

import (
	"encoding/json"
	"ethereum-explorer/logger"
	"net/http"
)

type ResponseData interface {
	GetDTO() interface{}
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.LogInternalServerError(r, err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func SendResponse(w http.ResponseWriter, data ResponseData) {
	jsonData, _ := json.Marshal(data.GetDTO())

	w.Write(jsonData)
}