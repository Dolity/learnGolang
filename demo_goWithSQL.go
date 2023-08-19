package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func creatingTable(db *sql.DB) {
	query := `CREATE TABLE user (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func Inset(db *sql.DB) {
	var username string
	var password string
	fmt.Scan(&username)
	fmt.Scan(&password)
	createdAt := time.Now()

	result, err := db.Exec("INSERT INTO user (username, password, created_at) VALUES (?, ?, ?)", username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
}

func Delete(db *sql.DB) {
	var deleteid int
	fmt.Scan(&deleteid)

	_, err := db.Exec("DELETE FROM user WHERE id = ?", deleteid)
	if err != nil {
		log.Fatal(err)
	}
}

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
	// creatingTable(db)
	// Inset(db)
	Delete(db)
}
