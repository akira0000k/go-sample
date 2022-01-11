package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"go-sample/09koga456/new-api2/test"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	target := NewRouter(&test.MockTodoController{})
	mux = http.NewServeMux()
	mux.HandleFunc("/todos/", target.HandleTodosRequest)

}

func TestGetTodos(t *testing.T) {
	r, _ := http.NewRequest("GET", "/todos/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPostTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	r, _ := http.NewRequest("POST", "/todos/", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 201 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	r, _ := http.NewRequest("PUT", "/todos/2", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/todos/2", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/todos/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 405 {
		t.Errorf("Response cod is %v", w.Code)
	}
}
// Akira@MBP controller % ls
// total 56
// drwxr-xr-x  10 Akira  staff   320  1 10 17:07 ./
// drwxr-xr-x  13 Akira  staff   416  1  9 23:17 ../
// drwxr-xr-x   3 Akira  staff    96 11 24 22:09 dto/
// -rw-r--r--   1 Akira  staff   482  1  9 22:06 router.go
// -rw-r--r--   1 Akira  staff   130  1  8 22:48 router_interface.go
// -rw-r--r--   1 Akira  staff  1551  1  8 22:56 router_test.go
// drwxr-xr-x   2 Akira  staff    64 11 29 20:26 tes/
// -rw-r--r--   1 Akira  staff  2224  1  9 22:07 todo_controller.go
// -rw-r--r--   1 Akira  staff   279  1  8 22:48 todo_controller_interface.go
// -rw-r--r--   1 Akira  staff  4590  1  8 22:56 todo_controller_test.goo
// |
// Akira@MBP controller % go test -v
// === RUN   TestGetTodos
// --- PASS: TestGetTodos (0.00s)
// === RUN   TestPostTodo
// --- PASS: TestPostTodo (0.00s)
// === RUN   TestPutTodo
// --- PASS: TestPutTodo (0.00s)
// === RUN   TestDeleteTodo
// --- PASS: TestDeleteTodo (0.00s)
// === RUN   TestInvalidMethod
// --- PASS: TestInvalidMethod (0.00s)
// PASS
// ok  	go-sample/09koga456/new-api2/controller	0.430s
// |
// Akira@MBP controller %
