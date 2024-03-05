package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jonatasemanuel/gote-server/controllers"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/api/v1/notes", controllers.GetAllNotes)
	router.Get("/api/v1/notes/{id}", controllers.GetNoteById)
	router.Post("/api/v1/notes", controllers.CreateNote)
	router.Put("/api/v1/notes/{id}", controllers.UpdateNote)
	router.Delete("/api/v1/notes/{id}", controllers.DeleteNote)

	router.Get("/api/v1/tags", controllers.GetAllTags)
	router.Post("/api/v1/tags", controllers.CreateTag)
	return router
}
