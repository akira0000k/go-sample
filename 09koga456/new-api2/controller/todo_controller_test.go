package controller

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"go-sample/09koga456/new-api2/controller/dto"
	"go-sample/09koga456/new-api2/test"
)

func TestGetTodos_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepository{})
	target.GetTodos(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var todosResponse dto.TodosResponse
	json.Unmarshal(body, &todosResponse)
	if len(todosResponse.Todos) != 0 {
		t.Errorf("Response is %v", todosResponse.Todos)
	}
}

func TestGetTodos_ExistTodo(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepositoryGetTodosExist{})
	target.GetTodos(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var todosResponse dto.TodosResponse
	json.Unmarshal(body, &todosResponse.Todos)
	if len(todosResponse.Todos) != 2 {
		t.Errorf("Response is %v", todosResponse.Todos)
	}
}

func TestGetTodos_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.GetTodos(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	if w.Body.Len() != 0 {
		t.Errorf("body is %v", w.Body.Len())
	}
}

func TestPostTodo_OK(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/todos/", json)

	target := NewTodoController(&test.MockTodoRepository{})
	target.PostTodo(w, r)

	if w.Code != 201 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Location") != r.Host+r.URL.Path+"2" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPostTodo_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/todos/", json)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.PostTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Location") != "" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPutTodo_OK(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todos/2", json)

	target := NewTodoController(&test.MockTodoRepository{})
	target.PutTodo(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTodo_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepository{})
	target.PutTodo(w, r)

	if w.Code != 400 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTodo_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todos/2", json)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.PutTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo_OK(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/2", nil)

	target := NewTodoController(&test.MockTodoRepository{})
	target.DeleteTodo(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.DeleteTodo(w, r)

	if w.Code != 400 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/2", nil)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.DeleteTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
}
// Akira@MBP controller % ls
// total 56
// drwxr-xr-x  10 Akira  staff   320  1 10 17:08 ./
// drwxr-xr-x  13 Akira  staff   416  1  9 23:17 ../
// drwxr-xr-x   3 Akira  staff    96 11 24 22:09 dto/
// -rw-r--r--   1 Akira  staff   482  1  9 22:06 router.go
// -rw-r--r--   1 Akira  staff   130  1  8 22:48 router_interface.go
// -rw-r--r--   1 Akira  staff  1551  1  8 22:56 router_test.goo
// drwxr-xr-x   2 Akira  staff    64 11 29 20:26 tes/
// -rw-r--r--   1 Akira  staff  2224  1  9 22:07 todo_controller.go
// -rw-r--r--   1 Akira  staff   279  1  8 22:48 todo_controller_interface.go
// -rw-r--r--   1 Akira  staff  4590  1  8 22:56 todo_controller_test.go
// |
// Akira@MBP controller % go test -v
// === RUN   TestGetTodos_NotFound
// --- PASS: TestGetTodos_NotFound (0.00s)
// === RUN   TestGetTodos_ExistTodo
// --- PASS: TestGetTodos_ExistTodo (0.00s)
// === RUN   TestGetTodos_Error
// --- PASS: TestGetTodos_Error (0.00s)
// === RUN   TestPostTodo_OK
// --- PASS: TestPostTodo_OK (0.00s)
// === RUN   TestPostTodo_Error
// --- PASS: TestPostTodo_Error (0.00s)
// === RUN   TestPutTodo_OK
// --- PASS: TestPutTodo_OK (0.00s)
// === RUN   TestPutTodo_InvalidPath
// --- PASS: TestPutTodo_InvalidPath (0.00s)
// === RUN   TestPutTodo_Error
// --- PASS: TestPutTodo_Error (0.00s)
// === RUN   TestDeleteTodo_OK
// --- PASS: TestDeleteTodo_OK (0.00s)
// === RUN   TestDeleteTodo_InvalidPath
// --- PASS: TestDeleteTodo_InvalidPath (0.00s)
// === RUN   TestDeleteTodo_Error
// --- PASS: TestDeleteTodo_Error (0.00s)
// PASS
// ok  	go-sample/09koga456/new-api2/controller	0.313s
// |
// Akira@MBP controller %
