package utils

import (
	"log"
	"net/http"
)

// Check - checks for errors
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckAndSendErrorResponse(err error, w http.ResponseWriter) {
	if err != nil {
		SendResponse(500, err.Error(), nil, w)
	}
}