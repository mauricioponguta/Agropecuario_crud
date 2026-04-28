package config

import (
	"database/sql"// Para conexiones con sql
	"fmt"
	"log" // imprimir y manejar logs

	_"github.com/lib/pq" // driver de postgres
)

var DB *sql.DB // Instancia global de la base de datos

// ConnectDB establece la conexión a la base de datos

func ConnectDB() {
	// Variables para la conexion 
	host:= "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "AgroCampo"
	schema:= "Agropecuario"

psqlInfo := fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
	host, port, user, password, dbname,schema,
)
// Cadena de conexión

// Abrir la conexión a la base de datos
db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	log.Fatalf("Error al conectar a la base de datos: %v", err)
}

err = db.Ping()
if err != nil {
	log.Fatalf("No se pudo conectar a la base de datos: %v", err)
}

fmt.Println("Conexión exitosa a la base de datos")
fmt.Println("Conectado a la db:", dbname, "Y esquema:", schema)
DB = db // Asignar la conexión a la variable global


}
