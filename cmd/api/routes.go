package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.home)

	router.HandlerFunc(http.MethodGet, "/notes", app.getAllNotes)
	router.HandlerFunc(http.MethodPost, "/notes", app.createNote)
	router.HandlerFunc(http.MethodGet, "/notes/:noteId", app.getNote)
	router.HandlerFunc(http.MethodPut, "/notes/:noteId", app.updateNote)
	router.HandlerFunc(http.MethodDelete, "/notes/:noteId", app.deleteNote)

	return middleware.Logger(router)
}
