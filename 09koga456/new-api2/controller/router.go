package controller

import (
	"net/http"
)

type router struct {
	tc TodoController
}

//func NewRouter(tc TodoController) Router {
func NewRouter(tc TodoController) *router {
	return &router{tc}
}

func (ro *router) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetTodos(w, r)
	case "POST":
		ro.tc.PostTodo(w, r)
	case "PUT":
		ro.tc.PutTodo(w, r)
	case "DELETE":
		ro.tc.DeleteTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}
