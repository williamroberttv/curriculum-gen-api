package database

import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq"
)

func Connect(user, password, dbname string) {
	//get the connection string url
	fmt.Println("Connecting to database...")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            user, password, dbname)
	// Connect to the database.
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
  defer db.Close()
	fmt.Println("Successfully connected to database!")
}

func checkErr(err error) {
	if err != nil {
			panic(err)
	}
}