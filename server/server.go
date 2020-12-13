package server

import (
	"net/http"

	"github.com/danielKim007/todolist/api"
	"github.com/danielKim007/todolist/db"
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

	return http.ListenAndServe(cfg.Address,
		loggingMiddleware(api.TodoListAPI()))
}
