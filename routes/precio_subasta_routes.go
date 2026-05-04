package routes

import (
	"agropecuario_crud/controllers"
	"github.com/gorilla/mux"
)

// Registro de rutas para la tabla tr_precioSubastaGanado

func RegisterPrecioSubastaGanadoRoutes(r *mux.Router) {
	r.HandleFunc("/precio_subasta_ganado", controllers.GetPreciosSubastaGanado).Methods("GET")
	r.HandleFunc("/precio_subasta_ganado/{id}", controllers.GetPrecioSubastaGanadoByID).Methods("GET")
	r.HandleFunc("/precio_subasta_ganado", controllers.CreatePrecioSubastaGanado).Methods("POST")
	r.HandleFunc("/precio_subasta_ganado/{id}", controllers.UpdatePrecioSubastaGanado).Methods("PUT")
	r.HandleFunc("/precio_subasta_ganado/{id}", controllers.DeletePrecioSubastaGanado).Methods("DELETE")
}
