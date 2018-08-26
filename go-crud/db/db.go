package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

//GetDatabase - returns a Database object
func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "developer:123@/XYZ_lakshan_dissanayake")
	if err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		err = db.Ping()
	}
	return db, err
}
