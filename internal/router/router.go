package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"go-crm/internal/handler"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", handler.Login).Methods("POST")
	r.HandleFunc("/api/me", handler.Me).Methods("GET")
	return r
}