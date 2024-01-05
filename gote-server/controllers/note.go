package controllers

import (
	"encoding/json"
	"net/http"

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
