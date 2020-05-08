package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zerefwayne/that-meme/handlers"
)

func configureCORS(r *mux.Router) http.Handler {

	handler := cors.AllowAll().Handler(r)

	return handler

}

func newRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.DefaultHandler)
	r.HandleFunc("/api/upload", handlers.UploadImageHandler)
	r.HandleFunc("/api/search", handlers.SearchHandler)

	return r

}

// NewHandler ...
func NewHandler() http.Handler {

	router := newRouter()

	handler := configureCORS(router)

	return handler

}
