package repo

import (
	"container/list"

	"github.com/lakshanwd/muve-go/go-crud/db"
)

//Select - Read from database
func Select(query string, fn delegate, params ...interface{}) (*list.List, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return executeReader(db, query, fn, params...)
}

//Insert - Insert into database
func Insert(query string, params ...interface{}) (int64, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	return executeInsert(db, query, params...)
}

//UpdateDelete - Updates or Delete from database
func UpdateDelete(query string, params ...interface{}) (int64, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	return executeUpdateDelete(db, query, params...)
}
