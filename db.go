package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"
)

func bootstrap() *sql.DB {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	folder := usr.HomeDir
	file := ".rowing.db"
	database := folder + "/" + file
	needTable := false

	if _, err := os.Stat(database); os.IsNotExist(err) {
		f, err := os.Create(database) // Create SQLite file
		check(err)

		err = f.Close()
		check(err)

		fmt.Printf("file: %s created\n", database)
		needTable = true
	}

	db, _ := sql.Open("sqlite3", database)
	if needTable == true {
		createTable(db)
	}
	return db
}

func createTable(db *sql.DB) {
	createTableSql := `CREATE TABLE rowing (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"date" TEXT,
		"distance" INTEGER,
		"duration" INTEGER,
		"actual" INTEGER,
		"pace" INTEGER,
		"power" INTEGER
	);`

	log.Println("Create rowing table...")
	statement, err := db.Prepare(createTableSql) // Prepare SQL Statement
	check(err)

	_, err = statement.Exec()
	check(err)

	log.Println("rowing table created")
}
