package routes

import (
	"net/http"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/controllers"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/middleware"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router) {
	authenticated := r.PathPrefix("/tasks").Subrouter()
	authenticated.Use(middleware.JWTMiddleware)

	authenticated.HandleFunc("", controllers.GetTasks).Methods("GET")
	authenticated.HandleFunc("/{id}", controllers.GetTask).Methods("GET")
	authenticated.HandleFunc("", controllers.CreateTask).Methods("POST")
	authenticated.HandleFunc("/{id}", controllers.UpdateTask).Methods("PUT")
	authenticated.HandleFunc("/{id}", controllers.DeleteTask).Methods("DELETE")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go Task Manager API"))
	})
}
