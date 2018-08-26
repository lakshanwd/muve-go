package repo

import (
	"container/list"
	"database/sql"
	"log"
)

//Repo - abstract repository interface
type Repo interface {
	Select(doc interface{}) (*list.List, error)
	Insert(doc interface{}) (int64, error)
	Update(doc interface{}) (int64, error)
	Remove(doc interface{}) (int64, error)
}

type delegate func(rows *sql.Rows, collection *list.List) error

func executeReader(db *sql.DB, query string, fn delegate, params ...interface{}) (*list.List, error) {
	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := list.New()
	start := make(chan bool)
	end := make(chan bool)
	go func() {
		<-start
		for rows.Next() {
			if err = fn(rows, data); err != nil {
				log.Fatal(err)
			}
		}
		end <- true
	}()
	start <- true
	<-end
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return data, err
}

func executeInsert(db *sql.DB, query string, params ...interface{}) (int64, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func executeUpdateDelete(db *sql.DB, query string, params ...interface{}) (int64, error) {
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
