package main

import (
	"database/sql"
	"fmt"
	"time"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var arg1 string
	if len(os.Args) > 1 {
		arg1 = os.Args[1]
	}
	fmt.Println("** start transact **", arg1)

	type Book struct {
		ID    int64
		Title string
	}

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`create table if not exists "BOOKS" ("ID" integer primary key, "TITLE" varchar(255))`)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	//10レコード挿入する
	stmt, err := tx.Prepare(`insert into books (id, title) values (?, ?)`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i:=0; i<16; i++ {
		id := fmt.Sprintf("%d%d", time.Now().Unix(), i)
		if _, err := stmt.Exec(id, fmt.Sprintf("title-%d", i)); err != nil {
			panic(err)
		}
	}

	if arg1 == "roll" {
		tx.Rollback()
	} else {
		tx.Commit()
	}


	rows, err := db.Query(`select * from books`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	nrec := 0
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title); err != nil {
			panic(err)
		}
		fmt.Printf("id: %d, title: %s\n", b.ID, b.Title)
		nrec++
	}
	fmt.Println(nrec, "records found")
	fmt.Println("** END **")
}
