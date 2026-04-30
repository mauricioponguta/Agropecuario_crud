package main

import (
	"agropecuario_crud/config"
	"agropecuario_crud/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// middleware CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	// DB
	config.ConnectDB()

	// Router
	r := mux.NewRouter()

	// Registrar rutas (igual que tu ejemplo)
	routes.RegisterCategoriaGanadoRoutes(r)
	routes.RegisterSubastaRoutes(r)
	routes.RegisterPrecioSubastaRoutes(r)

	log.Println("Servidor corriendo en el puerto :8082")

	http.ListenAndServe(":8082", enableCORS(r))
}