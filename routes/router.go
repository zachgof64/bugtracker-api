package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type NoDataResponse struct {
	StatusCode uint16 `json:"statusCode"`
	Message string `json:"message"`
}

type Header struct {
	Key string
	Value string
}


var router *mux.Router = mux.NewRouter()
var globalHeaders = []Header{
	{
		Key: "Content-Type",
		Value: "application/json",
	},
	{
		Key: "Accept",
		Value: "application/json",
	},
	{
		Key: "Access-Control-Allow-Origin",
		Value: "*",
	},
}


func globalHeaderHandler() mux.MiddlewareFunc {
	return mux.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("Authorization")
			var res NoDataResponse
			for _,h := range globalHeaders {
				w.Header().Add(h.Key, h.Value)
			}
			if apiKey == "038479357927598739" {
				next.ServeHTTP(w,r)
			} else {
				res.StatusCode = 401
				res.Message = "Not Authorized"
				w.WriteHeader(int(res.StatusCode))
				json.NewEncoder(w).Encode(res)
			}
			
		})
	})
}

// AddGlobalHeaders - adds headers to all requests
func AddGlobalHeaders(headers ...Header) {
		globalHeaders = append(globalHeaders, headers...)
}

func setupRoutes() {
	router.HandleFunc("/bugs", GetAllBugsHandler)
	router.HandleFunc("/bugs/{id}", GetBugHandler)
}

// SetupRouter - Setups mux router with global headers
func SetupRouter() *mux.Router {
	router.Use(globalHeaderHandler())
	setupRoutes()
	fmt.Println("API Started")
	return router
}