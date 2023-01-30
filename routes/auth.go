package routes

import (
	"housybe/handlers"
	"housybe/pkg/mysql"
	"housybe/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/signup", h.Register).Methods("POST")
	r.HandleFunc("/signin", h.Login).Methods("POST")
}
