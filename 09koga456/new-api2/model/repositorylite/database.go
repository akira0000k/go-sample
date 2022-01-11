package repositorylite

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func initSqlite3() {
	fmt.Println("repository.init() called...")
	if Db != nil { //only once
		return
	}
	var err error
	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	// 	"todo-app", "todo-password", "sample-api-db:3306", "todo",
	//)
	//Db, err = sql.Open("mysql", dataSourceName)

	fmt.Println(`Db = sql.Open("sqlite3")`)
	Db, err = sql.Open("sqlite3", "kogatodo.sql3")
	if err != nil {
		panic(err)
	}
	createTable(Db)
}
//    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
// AUTOINCREMENT is only allowed on an INTEGER PRIMARY KEY

func createTable(db *sql.DB) {
	createTodoTableSQL := `
CREATE TABLE IF NOT EXISTS todo (
    id             INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title          VARCHAR(40) NOT NULL,
    content       VARCHAR(100) NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp
);`
	//    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp

	log.Println("Create todo table...")
	statement, err := db.Prepare(createTodoTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("todo table created")
}
