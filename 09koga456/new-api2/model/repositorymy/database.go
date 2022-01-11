package repositorymy

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func initMysql() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"todo-app", "todo-password",
		"localhost:3306", //"sample-api-db:3306",
		"todo",
	)
	fmt.Println("--------------------")
	fmt.Println(dataSourceName)
	fmt.Println("--------------------")
	
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
