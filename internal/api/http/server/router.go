package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/owaisnadeemdev/roadguard/internal/api/http/handlers"
	"github.com/rs/cors"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignupHandle)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8100"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(router)

	return corsHandler
}
