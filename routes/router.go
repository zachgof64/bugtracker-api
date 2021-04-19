package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeuce/bugtracker-api/utils"
)


type Header struct {
	Key   string
	Value string
}

type RequestBody struct {
	Limit    int    `json:"limit,omitempty"`
	Title    string `json:"title,omitempty"`
	Resolved bool   `json:"resolved,omitempty"`
}

var router *mux.Router = mux.NewRouter()
var globalHeaders = []Header{
	{
		Key:   "Content-Type",
		Value: "application/json",
	},
	{
		Key:   "Accept",
		Value: "application/json",
	},
	{
		Key:   "Access-Control-Allow-Origin",
		Value: "*",
	},
}

func globalHeaderHandler() mux.MiddlewareFunc {
	return mux.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("Authorization")
			for _, h := range globalHeaders {
				w.Header().Add(h.Key, h.Value)
			}
			if apiKey == "038479357927598739" {
				next.ServeHTTP(w, r)
			} else {
				utils.SendResponse(401, "not authorized", nil, w)
			}

		})
	})
}

// AddGlobalHeaders - adds headers to all requests
func AddGlobalHeaders(headers ...Header) {
	globalHeaders = append(globalHeaders, headers...)
}

func setupRoutes() {
	router.HandleFunc("/bugs", BugsHandler).Methods("GET", "POST")
	router.HandleFunc("/bugs/{id}", BugHandler).Methods("GET",  "PATCH", "DELETE")
}

// SetupRouter - Setups mux router with global headers
func SetupRouter() *mux.Router {
	router.Use(globalHeaderHandler())
	setupRoutes()
	fmt.Println("API Started")
	return router
}
