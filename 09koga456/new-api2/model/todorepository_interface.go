package model

import (
	"go-sample/09koga456/new-api2/model/entity"
)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
}
