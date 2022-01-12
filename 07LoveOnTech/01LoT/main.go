package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type userInfo struct {
	id    string
	email string
	pass  string
	fName string
	lName string
}

func main() {
	ui := &userInfo {
		id:    "johnson123",
		email: "xxx@example.com",
		pass:  "123",
		fName: "ニャロメ",
		lName: "みよじ",
	}

	err := saveUserInfo2DB(ui)
	fmt.Println("[+]Succeeded to save userInfo")

	id, ok, err := checkValidity(ui.email, ui.pass)
	if !ok || err != nil {
		log.Println(err)
	} else if id == "" {
		log.Printf("[-]Email and Password don't match [%s : %s]\n", ui.email, ui.pass)
	} else {
		fmt.Printf("[+]UserId id %s\n", id)
	}

	for i:=0; i<2; i++ {
		if ok := checkUID(id); ok {
			fmt.Printf("[+]UserId id %s exist.\n", id)
		} else {
			fmt.Printf("[+]UserId id %s not exist.\n", id)
		}
		id = "gomashio"
	}
}

func saveUserInfo2DB(user *userInfo) error {
	fmt.Println("[+]Save data to SQLite ...")
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	if err != nil {
		log.Println("[-]sql.Open")
		return err
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (UserID text PRIMARY KEY, Email text, Pass text, FirstName text, LastName text)")
	defer stmt.Close()
	if err != nil {
		log.Println("[-]sql.Prepare (create table: user)")
		return err
	}
	stmt.Exec()

	stmt, err = db.Prepare("INSERT INTO user (UserID, Email, Pass, FirstName, LastName) VALUES($1, $2, $3, $4, $5)")
	defer stmt.Close()
	if err != nil {
		log.Println("[-]sql.Prepare (INSERT INTO user)")
		return err
	}
	stmt.Exec(user.id, user.email, user.pass, user.fName, user.lName)

	return nil
}

func checkValidity(email string, pass string) (string, bool, error) {
	fmt.Println("[+]Checking if login info matches...", email, " : ", pass)
	var userID, e, p string

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	if err != nil {
		log.Println("[-]sql.Open", err)
		return "", false, err
	}

	//rows, err := db.Query("SELECT UserID, Email, Pass FROM user WHERE Email = $1 AND Pass = $2", email, pass) //andとかを小文字にすると正常に動かない?
	rows, err := db.Query("select userid, email, pass from USER where email = $1 and pass = $2", email, pass) //動いた
	defer rows.Close()
	if err != nil {
		log.Printf("[-]db.Query")
		return "", false, err
	}

	for rows.Next() {
		err := rows.Scan(&userID, &e, &p)
		if err != nil {
			log.Println("[-]rows.Scan")
			return "", false, err
		}
		fmt.Println("[+]Match!: ", e, " : ", p)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return "", false, err
	}
	return userID, true, nil
}

func checkUID(id string) bool {
	fmt.Println("[+]Checking if UID matches...", id)
	exist := false
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	if err != nil {
		log.Println("[-]sql.Open: ", err)
		return exist
	}
	row := db.QueryRow("select exists (select * from user where userid = $1)", id)
	if err := row.Scan(&exist); err != nil {
		log.Println("[-]No uid in DB: ", err)
		return exist
	}
	return exist
}

	     
