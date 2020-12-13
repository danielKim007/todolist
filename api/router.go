package api

import (
	"net/http"
)

func TodoListAPI() http.Handler {
	router := router.NewRouter()
	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)

	return router
}
