package repo

import (
	"container/list"
	"database/sql"

	util "github.com/lakshanwd/db-helper/mysql"
	"github.com/lakshanwd/muve-go/go-crud/db"
)

var connection *sql.DB

//SetupRepo - setup database connections
func SetupRepo() (err error) {
	connection, err = db.GetDatabase()
	return
}

//CloseRepo - close database connections
func CloseRepo() {
	if connection != nil {
		connection.Close()
	}
}

//Select - Read from database
func Select(query string, fn util.Delegate, params ...interface{}) (*list.List, error) {
	return util.ExecuteReader(connection, query, fn, params...)
}

//Insert - Insert into database
func Insert(query string, params ...interface{}) (int64, error) {
	return util.ExecuteInsert(connection, query, params...)
}

//UpdateDelete - Updates or Delete from database
func UpdateDelete(query string, params ...interface{}) (int64, error) {
	return util.ExecuteUpdateDelete(connection, query, params...)
}
