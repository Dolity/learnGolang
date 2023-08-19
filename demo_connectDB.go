package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func query(db *sql.DB) {
	var (
		id         int
		coursename string
		price      float64
		instructor string
	)

	query := "SELECT id, coursename, price, instructor FROM onlinecourse WHERE id = ?"
	if err := db.QueryRow(query, 2).Scan(&id, &coursename, &price, &instructor); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, coursename, price, instructor)
}

func main() {
	db, err := sql.Open("mysql", "root:abc123@tcp(127.0.0.1:3305)/todojojo")
	if err != nil {
		fmt.Println("Connect Error: ", err)
	} else {
		fmt.Println("Connect Success")
	}
	query(db)
}
