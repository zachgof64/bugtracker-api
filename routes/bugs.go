package routes

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/zeuce/bugtracker-api/utils"
)

type Bug struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Resolved bool   `json:"resolved"`
}

var db *sql.DB = utils.ConnectUsingENV("MYSQL_CONN_STRING")


// GetAllBugsHandler - handler for getting all bugs
func BugsHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	var limit string
	bodyDecErr := json.NewDecoder(r.Body).Decode(&body)
	if bodyDecErr != io.EOF {
		limit = strconv.Itoa(body.Limit)
	} else {
		limit = strconv.Itoa(10)
	}

	query, queryErr := db.Query("SELECT * FROM bugs LIMIT " + limit)
	utils.CheckAndSendErrorResponse(queryErr, w)
	defer query.Close()

	var bugs []Bug

	for query.Next() {
		var bug Bug
		query.Scan(&bug.Id, &bug.Title, &bug.Resolved)
		bugs = append(bugs, bug)
	}
	utils.SendResponse(200,"success", bugs, w)
}

func getBug(id string, w http.ResponseWriter) {
	var bug Bug
	queryErr := db.QueryRow("SELECT * FROM bugs WHERE id ="+id).Scan(&bug.Id, &bug.Title, &bug.Resolved)
	if queryErr != nil && strings.Contains(queryErr.Error(), "no rows in result") {
		utils.SendResponse(404, "not found", nil, w)
	} else {
		utils.SendResponse(200, "success", bug, w)
	}
}

// func postBugs(title string, w http.ResponseWriter) {
// 	db := utils.ConnectUsingENV("MYSQL_CONN_STRING")
// 	_, queryErr := db.Query("INSERT INTO bugs(title) VALUES('"+ title +"')")
// 	utils.CheckAndSendErrorResponse(queryErr, w)
// }

func deleteBug(id string, w http.ResponseWriter) {

	_, queryErr := db.Query("DELETE FROM bugs WHERE id =" + id)
	utils.CheckAndSendErrorResponse(queryErr,w )

	utils.SendResponse(200, "success", nil, w)
}

// GetBugHandler - gets a bug with specified id
func BugHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string = vars["id"]
	switch r.Method {
	case "GET":
		getBug(id, w)
	// case "POST":
	// 	// postBug()
	case "DELETE":
		deleteBug(id, w)
	}
}
