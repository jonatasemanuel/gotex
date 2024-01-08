package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonatasemanuel/gote-server/helpers"
	"github.com/jonatasemanuel/gote-server/services"
)

var note services.Note

// POST/notes
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var noteData services.Note
	err := json.NewDecoder(r.Body).Decode(&noteData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	noteCreated, err := note.CreateNote(noteData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, noteCreated)
}

// GET/notes
func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	all, err := note.GetAllNotes()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"notes": all})
}

// GET/notes/{id}
func GetNoteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	note, err := note.GetNoteById(id)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, note)
}

// PUT/notes/{id}
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var noteData services.Note

	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&noteData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	noteUpdated, err := note.UpdateNote(id, noteData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	helpers.WriteJSON(w, http.StatusOK, noteUpdated)
}

// DELETE/notes/{id}
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := note.DeleteNote(id); err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"message": "Successfull deletion"})
}
