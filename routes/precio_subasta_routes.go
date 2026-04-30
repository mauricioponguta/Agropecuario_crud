package routes

import (
	"agropecuario_crud/controllers"

	"github.com/gorilla/mux"
)

func RegisterPrecioSubastaRoutes(r *mux.Router) {

	r.HandleFunc("/precio_subasta", controllers.CreatePrecioSubasta).Methods("POST")
	r.HandleFunc("/precio_subasta", controllers.GetPreciosSubasta).Methods("GET")
	r.HandleFunc("/precio_subasta/{id}", controllers.DeletePrecioSubasta).Methods("DELETE")
}