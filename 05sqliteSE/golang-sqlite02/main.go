package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	dbfile := "db/sqlite-database.db"
	//os.Remove(dbfile) //不要
	_ = os.DevNull
	
	log.Println("Creating sqlite-database.db...")
	// 既存でも初期化される
	// 作成出来ない場合はここで見つける。
	file, err := os.Create(dbfile)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	// sql.Openでは dbfileの生成不可とか、書き込み不可とかはエラーにならない。のでわからない
	sqliteDatabase, _ := sql.Open("sqlite3", dbfile)
	//sqliteDatabase, sm := sql.Open("sqlite3", dbfile)
	//fmt.Println(sqliteDatabase, sm)
	defer sqliteDatabase.Close()

	//書き込み権限がないのはここでわかる
	createTable(sqliteDatabase)

	// INSERT RECORDS
	insertStudent(sqliteDatabase, "0011", "Liana Kim", "Bachelor")
	insertStudent(sqliteDatabase, "0012", "Glen Rangel", "Bachelor")
	insertStudent(sqliteDatabase, "0013", "Martin Martins", "Master")
	insertStudent(sqliteDatabase, "0014", "Alayna Armitage", "PHD")
	insertStudent(sqliteDatabase, "0015", "Marni Benson", "Bachelor")
	insertStudent(sqliteDatabase, "0016", "Derrick Griffiths", "Master")
	insertStudent(sqliteDatabase, "0017", "Leigh Daly", "Bachelor")
	insertStudent(sqliteDatabase, "0018", "Marni Benson", "PHD")
	insertStudent(sqliteDatabase, "0019", "Klay Correa", "Bachelor")

	displayStudents(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"code" TEXT,
		"name" TEXT,
		"program" TEXT
		);`
	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("student table created")
}

func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var code string
		var name string
		var program string
		row.Scan(&id, &code , &name, &program)
		log.Printf("Student:%d   %s   %s   %s\n", id, code, name, program)
	}
}
