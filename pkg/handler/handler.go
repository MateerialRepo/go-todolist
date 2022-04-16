package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

func TodoHandlers() http.Handler {
	//defining a group route for "/todo/" after a request
	rg := chi.NewRouter()

	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodoList)
		r.Post("/", createTodoList)
		r.Put("/{id}", editTodoList)
		r.Delete("/{id}", deleteTodoList)
	})

	return rg
}
