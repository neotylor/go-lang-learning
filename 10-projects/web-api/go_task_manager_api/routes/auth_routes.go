package routes

import (
	"net/http"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/controllers"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.Handle("/me", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
}
