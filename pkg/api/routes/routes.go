package routes

import (
	"github.com/andrleite/relayr-app/pkg/api/controllers"
	"github.com/gorilla/mux"
)

// NewRouter initiate a new mux router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.PostUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.PutUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/feedbacks", controllers.GetFeedbacks).Methods("GET")
	r.HandleFunc("/feedbacks", controllers.PostFeedback).Methods("POST")
	r.HandleFunc("/healthz", controllers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/metrics", controllers.MetricsHandler).Methods("GET")
	return r
}
