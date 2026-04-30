package routes

import (
	"agropecuario_crud/controllers"

	"github.com/gorilla/mux"
)

func RegisterSubastaRoutes(r *mux.Router) {

	r.HandleFunc("/subasta", controllers.CreateSubasta).Methods("POST")
	r.HandleFunc("/subasta", controllers.GetSubastas).Methods("GET")
	r.HandleFunc("/subasta/{id}", controllers.GetSubastaByID).Methods("GET")
	r.HandleFunc("/subasta/{id}", controllers.UpdateSubasta).Methods("PUT")
	r.HandleFunc("/subasta/{id}", controllers.DeleteSubasta).Methods("DELETE")
}