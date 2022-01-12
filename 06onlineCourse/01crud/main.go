package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`

	_, err := DbConnection.Exec(cmd)
	if err != nil {
		fmt.Println("エラー")
		log.Fatalln(err)
	}

	// C/CRUD
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"

	_, err = DbConnection.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = DbConnection.Exec(cmd, "Mike", 0)
	if err != nil {
		log.Fatalln(err)
	}

	// U/CRUD
	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DbConnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}

	// R/CRUD
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// R/CRUD
	cmd = "SELECT * FROM person where age = ?"

	row := DbConnection.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// D/CRUD
	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	// R/CRUD
	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows1, _ := DbConnection.Query(cmd)
	defer rows1.Close()
	var pp1 []Person
	for rows1.Next() {
		var p Person
		err := rows1.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp1 = append(pp1, p)
	}
	err = rows1.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp1 {
		fmt.Println(p.Name, p.Age)
	}
}
