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
		// permiten cualquier origen de la peticion
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
	// Conectar a la base de datos
	config.ConnectDB()

	// Crear el router
	r := mux.NewRouter()

	// Registrar las rutas
	routes.RegisterCategoriaGanadoRoutes(r)
	routes.RegisterSubastaRoutes(r)
	routes.RegisterPrecioSubastaGanadoRoutes(r)

	log.Println("Servidor corriendo en el puerto :8082")

	http.ListenAndServe(":8082", enableCORS(r))

}
