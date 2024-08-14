package server

import (
	"github.com/gorilla/mux"
	"github.com/owaisnadeemdev/roadguard/internal/api/http/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignupHandle)
	return router
}
