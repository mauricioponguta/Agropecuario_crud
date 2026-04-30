package routes

import (
	"agropecuario_crud/controllers"

	"github.com/gorilla/mux"
)

func RegisterCategoriaGanadoRoutes(r *mux.Router) {

	r.HandleFunc("/categoria_ganado", controllers.CreateCategoriaGanado).Methods("POST")
	r.HandleFunc("/categoria_ganado", controllers.GetCategoriasGanado).Methods("GET")
	r.HandleFunc("/categoria_ganado/{id}", controllers.GetCategoriaGanadoByID).Methods("GET")
	r.HandleFunc("/categoria_ganado/{id}", controllers.UpdateCategoriaGanado).Methods("PUT")
	r.HandleFunc("/categoria_ganado/{id}", controllers.DeleteCategoriaGanado).Methods("DELETE")
}