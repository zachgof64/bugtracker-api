package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeuce/bugtracker-api/utils"
)

type Bug struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Resolved bool `json:"resolved"`
	// Comments string `json:"comments"`
}

type GetAllBugsResponse struct {
	StatusCode uint16 `json:"statusCode"`
	Message string `json:"message"`
	Data []Bug `json:"data"`
}

type GetBugResponse struct {
	StatusCode uint16 `json:"statusCode"`
	Message string `json:"message"`
	Data Bug `json:"data"`
}

// GetAllBugsHandler - handler for getting all bugs
func GetAllBugsHandler(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectUsingENV("MYSQL_CONN_STRING")
	defer db.Close()
	query, queryErr := db.Query("SELECT * FROM bugs")
	utils.Check(queryErr)
	defer query.Close()

	var bugs []Bug

	for	query.Next() {
		var bug Bug
		query.Scan(&bug.Id, &bug.Title, &bug.Resolved)
		bugs = append(bugs, bug)
	}

	d := GetAllBugsResponse {
		StatusCode: 200,
		Message: "got all bugs",
		Data: bugs,
	}
	w.WriteHeader(int(d.StatusCode))
	
}

// GetBugHandler - gets a bug with specified id
func GetBugHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string = vars["id"]
	db := utils.ConnectUsingENV("MYSQL_CONN_STRING")
	defer db.Close()
	var bug Bug
	queryErr := db.QueryRow("SELECT * FROM bugs WHERE id ="  + id).Scan(&bug.Id, &bug.Title, &bug.Resolved)
	utils.Check(queryErr)
	d := GetBugResponse {
		StatusCode: 200,
		Message: "got all bugs",
		Data: bug,
	}


	json.NewEncoder(w).Encode(d)
}	
	