package main

import (
	"net/http"

	"github.com/zeuce/bugtracker-api/routes"
	"github.com/zeuce/bugtracker-api/utils"
)

func serve() {
	utils.Check(http.ListenAndServe(":9999", routes.SetupRouter()))
}

func main() {
	serve()

}
