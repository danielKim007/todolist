package db

import (
	"github.com/danielKim007/todolist/todo"
)

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

// CreateTodoList creates a new todo list
func CreateTodoList(name string) (list todo.List, err error) {
	list.Name = name
	err = db.QueryRow(`INSERT INTO todo_list (name) VALUES ($1)
		 RETURNING id`, name).Scan(&list.ID)
	return
}

// RenameTodoList renames a todo list
func RenameTodoList(id int, newName string) error {
	res, err := db.Exec(`UPDATE todo_list
		SET name = $1
		WHERE id = $2`, newName, id)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
