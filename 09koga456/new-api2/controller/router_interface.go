package controller

import (
	"net/http"
)

type Router interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}
