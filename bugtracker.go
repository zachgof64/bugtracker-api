package main

import (
	"log"
	"net/http"

	"github.com/zeuce/bugtracker-api/routes"
)

func serve() {
	log.Fatal(http.ListenAndServe(":9999", routes.SetupRouter())) 
}

func main() {
	serve()
}
