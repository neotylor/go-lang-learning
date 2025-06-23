package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	RegisterAuthRoutes(r) // Register /login and /register routes
	RegisterTaskRoutes(r) // Register protected task routes
	// Add future route groups here
}
