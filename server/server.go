package server

import (
	"net/http"
	"https://github.com/danielKim007/todolist/api"
	"https://github.com/danielKim007/todolist/db"
)

//Config server config
type Config struct {
	Address     string
	DatabaseURL string
}

// ListenAndServer starts up the server
func ListenAndServer(cfg Config) error {
	if err := db.Connect(cfg.DatabaseURL); err != nil {
		return err
	}

	return http.ListenAndServer(cfg.Address,
		loggingMiddleware(api.TodoListAPI()))
}
