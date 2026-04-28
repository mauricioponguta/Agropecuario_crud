package main

import (
	"fmt"
	"log"
	"net/http"

	"agropecuario_crud/config"
	"agropecuario_crud/routes"

	"github.com/gorilla/mux"
)

func main() {

	// Inicializar conexión a la base de datos
	config.ConnectDB()

	// Crear router
	r := mux.NewRouter()

	// Registrar rutas
	routes.SubastaRoutes(r)

	// Puerto del servidor
	port := "3000"

	fmt.Println("Servidor corriendo en http://localhost:" + port)

	// Levantar servidor
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}