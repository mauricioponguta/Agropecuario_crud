package routes

import (
	"github.com/gorilla/mux"
	"agropecuario_crud/controllers"
)

func SubastaRoutes(r *mux.Router) {

	r.HandleFunc("/subastas", controllers.GetSubastas).Methods("GET")
	r.HandleFunc("/subastas/{id}", controllers.GetSubastaByID).Methods("GET")
	r.HandleFunc("/subastas", controllers.CreateSubasta).Methods("POST")
	r.HandleFunc("/subastas/{id}", controllers.UpdateSubasta).Methods("PUT")
	r.HandleFunc("/subastas/{id}", controllers.DeleteSubasta).Methods("DELETE")
}