package routes

import (
	"housybe/handlers"
	"housybe/pkg/mysql"
	"housybe/repositories"

	"github.com/gorilla/mux"
)

func HouseRoutes(r *mux.Router) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerHouse(houseRepository)

	r.HandleFunc("/houses", h.AddHouse).Methods("POST")
	r.HandleFunc("/house", h.FindHouses).Methods("GET")
	// r.HandleFunc("/house/{id}", h.GetHouse).Methods("GET")
}
