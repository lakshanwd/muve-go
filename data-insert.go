package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

var channel chan bool
var quit chan bool

func main() {
	channel = make(chan bool)
	quit = make(chan bool)
	if conn, err := getDatabase(); err == nil {
		defer fmt.Println("data inserted")
		defer conn.Close()
		query := "insert into tbl_booking(vehicle_id, pick_up_address, drop_address, total_distance, created_on) values (?,?,?,?,?)"
		stmt, err := conn.Prepare(query)
		if err != nil {
			return
		}
		defer stmt.Close()
		for i := 0; i < 3; i++ {
			go insertData(stmt, i*100, (i+1)*100)
			channel <- true
		}
		for i := 0; i < 10; i++ {
			<-quit
		}
	}
}

func getDatabase() (*sql.DB, error) {
	return sql.Open("mysql", "developer:123@/XYZ_lakshan_dissanayake")
}

func insertData(statement *sql.Stmt, start int, end int) {
	<-channel
	fmt.Println("started")
	for i := start; i < end; i++ {
		temp := (i + 1) % 10000
		statement.Exec(temp, fmt.Sprintf("pick up address %d", (temp+1)), fmt.Sprintf("drop address %d", (temp+1)), 10, "2018-08-24 20:39:00")
	}
	fmt.Println("done")
	quit <- true
}
