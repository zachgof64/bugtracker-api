package utils

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)


func ConnectUsingENV(varName string) *sql.DB {
	var connString string
	var env string = os.Getenv(varName)
	if len(env) > 0{
		connString = env
	}
	
	db, dbConnErr :=  sql.Open("mysql", connString)
	Check(dbConnErr)

	return db
}