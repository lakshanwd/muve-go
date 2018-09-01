package repo

import (
	"container/list"
	"database/sql"

	"github.com/lakshanwd/muve-go/go-crud/db"
)

var connection *sql.DB

//SetupRepo - setup database connections
func SetupRepo() {
	connection, _ = db.GetDatabase()
}

//CloseRepo - close database connections
func CloseRepo() {
	connection.Close()
}

//Select - Read from database
func Select(query string, fn delegate, params ...interface{}) (*list.List, error) {
	return executeReader(connection, query, fn, params...)
}

//Insert - Insert into database
func Insert(query string, params ...interface{}) (int64, error) {
	return executeInsert(connection, query, params...)
}

//UpdateDelete - Updates or Delete from database
func UpdateDelete(query string, params ...interface{}) (int64, error) {
	return executeUpdateDelete(connection, query, params...)
}
