package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseData interface{}

type Response struct {
	StatusCode int       `json:"statusCode"`
	Message    string       `json:"message"`
	Data       ResponseData `json:"data,omitempty"`
}

func SendResponse(statusCode int, message string, data ResponseData,  w http.ResponseWriter){
	d := Response{
		StatusCode: statusCode,
		Message: message,
		Data: data,
	}

	json.NewEncoder(w).Encode(d)
}