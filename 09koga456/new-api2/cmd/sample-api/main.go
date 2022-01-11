package main

import (
	"net/http"
	"fmt"
	"os"

	"go-sample/09koga456/new-api2/controller"
	"go-sample/09koga456/new-api2/model/repositorymy"
	"go-sample/09koga456/new-api2/model/repositorylite"
)

func main() {
	var arg1 string
	if len(os.Args) > 1 {
		arg1 = os.Args[1]
	}
	fmt.Println("main start...", arg1)

	//var tr model.TodoRepository
	//var tc controller.TodoController
	var ro controller.Router
	if arg1 == "mysql" {
		ro = controller.NewRouter(
			controller.NewTodoController(
				repositorymy.NewTodoRepository()))
	} else {
		ro = controller.NewRouter(
			controller.NewTodoController(
				repositorylite.NewTodoRepository()))
	}

	server := http.Server{
		Addr: ":8080",
	}

	fmt.Println("main http.HandleFunc(/todos/")

	http.HandleFunc("/todos/", ro.HandleTodosRequest)

	fmt.Println("server.ListenAndServe()")

	server.ListenAndServe()
}
