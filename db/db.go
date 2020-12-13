package db

import (
	"database/sql"

	"github.com/danielKim007/todolist/todo"
	_ "github.com/lib/pq"
)

var db *sql.DB

//Connect to database
func Connect(url string) error {
	c, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	db = c
	return nil
}

func GetTodoLists() ([]todo.List, error) {
	rows, err := db.Query(`SELECT id, name FROM todo_list`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := []todo.List{}
	for rows.Next() {
		var list todo.List
		if err := rows.Scan(&list.ID, &list.Name); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}
