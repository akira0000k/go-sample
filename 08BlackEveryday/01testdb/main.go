package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("main start")
	dbfile := "./test.db"
	//os.Remove(dbfile) //不要
	_ = os.DevNull
	
	log.Println("Creating", dbfile, "...")
	// 既存でも初期化される
	// 作成出来ない場合はここで見つける。
	file, err := os.Create(dbfile)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println(dbfile, " created")
	//-rw-r--r--   1 Akira  staff        0  1 14 00:16 test.db
	//return
	
	//テープルの作成
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		//`CREATE TABLE IF NOT EXISTS "BOOKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255))`,
		`create table if not exists "BOOKS" ("ID" integer primary key, "TITLE" varchar(255))`,
	)
	if err != nil {
		panic(err)
	}
	
	//レコードの挿入
	res, err := db.Exec(
		//`INSERT INTO BOOKS (ID, TITLE) VALUES (?, ?)`,
		`insert into BOOKS (ID, TITLE) values (?,?)`,
		123,
		"伝染るんです",
	)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("res.LastInsertId:", id)

	//レコードの挿入2
	res, err = db.Exec(
		//`INSERT INTO BOOKS (ID, TITLE) VALUES (?, ?)`,
		`insert into BOOKS (ID, TITLE) values (?,?)`,
		456,
		"熊のプーさん",
	)
	if err != nil {
		panic(err)
	}

	id, err = res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("res.LastInsertId:", id)

	//複数レコードの取得
	rows, err := db.Query(
		`select * from books`,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}
		fmt.Printf("id: %d, title: %s\n", id, title)
	}
	rows.Close()
	
	//レコードの更新
	res, err = db.Exec(
		`update books set title=? where id=?`,
		"熊のプー太郎",
		id)
	if err != nil{
		panic(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("affected by update: %d\n", affect)

	//1件のみ取得
	row := db.QueryRow(
		`select * from books where id=?`, id)
	var title string
	err = row.Scan(&id, &title)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Not found")
	case err != nil:
		panic(err)
	default:
		fmt.Printf("id: %d, title: %s\n", id, title)
	}

	//レコードの削除
	res, err = db.Exec(
		`delete from books where id=?`, id)
	if err != nil{
		panic(err)
	}
	affect, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("affected by delete: %d\n", affect)

	
	//1件のみ取得
	row = db.QueryRow(
		`select * from books where id=?`, id)
	err = row.Scan(&id, &title)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Not found")
	case err != nil:
		panic(err)
	default:
		fmt.Printf("id: %d, title: %s\n", id, title)
	}

	//複数レコードの取得
	rows, err = db.Query(
		`select * from books`,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}
		fmt.Printf("id: %d, title: %s\n", id, title)
	}
	
	fmt.Println()
}
