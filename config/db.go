package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB establece la conexión a la base de datos
func ConnectDB() {

	// Variables para la conexión
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres" // Cambia esto por tu contraseña real si es diferente
	dbname := "AgroCampo"
	schema := "agropecuario"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	// Abrir la conexión
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	fmt.Println("Conexión exitosa a la base de datos")
	fmt.Println("Conectado a la db:", dbname, "y esquema:", schema)

	DB = db
}