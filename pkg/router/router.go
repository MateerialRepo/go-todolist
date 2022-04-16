package router

import (
	"github.com/MateerialRepo/golang-todo/pkg/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", HomeHandler)
	r.Mount("todo", handler.TodoHandlers())

	return r
}
