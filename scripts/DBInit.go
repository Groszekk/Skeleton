package scripts

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	var err error
	sqlitedb, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	// defer sqlitedb.Close()
	CreateTable(sqlitedb)
	return sqlitedb
}

func CreateTable(db *sql.DB) {
	user_table := `CREATE TABLE users (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"nick" TEXT,
		"email" TEXT,
		"password" TEXT,
		"verify" TEXT,
		"active" BIT
		);`

	statement, err := db.Prepare(user_table)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	statement.Exec()
	fmt.Println("users table created")
}
