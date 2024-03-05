package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jonatasemanuel/gote-server/helpers"
	"github.com/jonatasemanuel/gote-server/services"
)

var tag services.Tag

// POST/tags
func CreateTag(w http.ResponseWriter, r *http.Request) {
	var tagData services.Tag
	err := json.NewDecoder(r.Body).Decode(&tagData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	tagCreated, err := tag.CreateTag(tagData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, tagCreated)
}

// GET/tags
func GetAllTags(w http.ResponseWriter, r *http.Request) {
	all, err := tag.GetAllTags()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"tags": all})
}
