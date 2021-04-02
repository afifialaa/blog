package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:admin12345@tcp(127.0.0.1:3306)/events")

	if err != nil {
		fmt.Println("Failed to connect to db")
		panic(err.Error())
	}

	defer db.Close()
}

// Create database
func createDB(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS events")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database was created")
	db.Close()
}

func Seed() {
	fmt.Println("Seeding")

	db, err := sql.Open("mysql", "root:admin12345@tcp(127.0.0.1:3306)/events")
	createDB(db)

	if err != nil {
		fmt.Println("Failed to connect to db")
		panic(err.Error())
	}

	defer db.Close()

	_, err = db.Query("INSERT INTO users (email, password) VALUES ( 'afifi@gmail.com', 'afifi123' )")
	_, err = db.Query("INSERT INTO users (email, password) VALUES ( 'omneya@gmail.com', 'omneya123' )")
	_, err = db.Query("INSERT INTO users (email, password) VALUES ( 'luka@gmail.com', 'luka123' )")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Data inserted into database")
	}
}
