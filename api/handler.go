package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danielKim007/todolist/db"
	"github.com/danielKim007/todolist/todo"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	if err != nil {
		log.Println("internal error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{ Error string }{"Internal Error"})
		return
	}
	json.NewEncoder(w).Encode(lists)
}

func cretaeTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	todoList, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, todoList)
}
