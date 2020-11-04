package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/go-sql-driver/mysql.v1"
)

var db *sql.DB

// GetMongoDB function to return DB connection
func GetDBconn() *sql.DB {
	dbName := "go_account"
	fmt.Println("conn info:", dbName)
	db, err := sql.Open("mysql","root:priya@1298@tcp(127.0.0.1:3306)/go_account")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	return db
}
