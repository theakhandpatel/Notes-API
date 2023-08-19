package main

import (
	"encoding/json"
	"fmt"
	"movieapi/internal/data"
	"net/http"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello , Kicking and Alive")
}

func (app *application) getAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := app.models.Notes.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func (app *application) createNote(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		Body  string
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	note := data.Note{
		Title:     input.Title,
		Body:      input.Body,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	err = app.models.Notes.Insert(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (app *application) getNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := app.readIDParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := app.models.Notes.Get(noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (app *application) updateNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := app.readIDParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := app.models.Notes.Get(noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var input struct {
		Title *string `json:"title"`
		Body  *string `json:"body"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if input.Title != nil {
		note.Title = *input.Title
	}
	if input.Body != nil {
		note.Body = *input.Body
	}

	note.UpdatedAt = time.Now()

	err = app.models.Notes.Update(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (app *application) deleteNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := app.readIDParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.models.Notes.Delete(noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
