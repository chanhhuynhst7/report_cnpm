package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "chanh"
)

var (
	id    int
	name  string
	major string
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	/*
		//insert data
		sqlStatement := `INSERT INTO Student (id, name, major)
		    VALUES ($1, $2, $3)`
		_, err = db.Exec(sqlStatement, 01, "Chanh", "IT")
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
	*/
	//select data để print ra
	rows, err := db.Query(`SELECT * FROM Student`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&id, &name, &major)
		fmt.Println(id, name, major)
	}
}
