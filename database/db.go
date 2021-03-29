package database

import (
	"database"
	"database/sql"
	"fmt"
)

func connect() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:admin12345@tcp(127.0.0.1:3306)/events")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
