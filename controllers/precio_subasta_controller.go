package controllers

import (
	"agropecuario_crud/config"
	"agropecuario_crud/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // Libreria para crear rutas
)

// GetPreciosSubastaGanado obtiene todos los precios de subasta ganado con filtros opcionales
func GetPreciosSubastaGanado(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_precio_subasta, id_subasta, id_categoria_ganado, activo, fecha_creacion, fecha_modificacion FROM \"tr_precioSubastaGanado\" WHERE 1=1"

	idSubasta := r.URL.Query().Get("id_subasta")
	idCategoriaGanado := r.URL.Query().Get("id_categoria_ganado")

	if idSubasta != "" {
		query += " AND id_subasta = " + idSubasta
	}

	if idCategoriaGanado != "" {
		query += " AND id_categoria_ganado = " + idCategoriaGanado
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.PrecioSubastaGanado

	for rows.Next() {
		var p models.PrecioSubastaGanado
		rows.Scan(&p.IDPrecioSubasta, &p.IDSubasta, &p.IDCategoriaGanado, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)
		list = append(list, p)
	}

	respondJSON(w, 200, list)
}

// GetPrecioSubastaGanadoByID obtiene un precio de subasta ganado por su ID
func GetPrecioSubastaGanadoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Printf("ID recibido: %s\n", id) // Imprime el ID recibido para depuración

	var p models.PrecioSubastaGanado

	err := config.DB.QueryRow(
		"SELECT id_precio_subasta, id_subasta, id_categoria_ganado, activo, fecha_creacion, fecha_modificacion FROM \"tr_precioSubastaGanado\" WHERE id_precio_subasta = $1",
		id,
	).Scan(&p.IDPrecioSubasta, &p.IDSubasta, &p.IDCategoriaGanado, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Id no encontrado"})
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, 200, p)
}

// CreatePrecioSubastaGanado crea un nuevo precio de subasta ganado
func CreatePrecioSubastaGanado(w http.ResponseWriter, r *http.Request) {
	var p models.PrecioSubastaGanado

	json.NewDecoder(r.Body).Decode(&p)

	error := config.DB.QueryRow(
		"INSERT INTO \"tr_precioSubastaGanado\" (id_subasta, id_categoria_ganado, activo) VALUES ($1, $2, $3) RETURNING id_precio_subasta, fecha_creacion, fecha_modificacion",
		p.IDSubasta, p.IDCategoriaGanado, p.Activo,
	).Scan(&p.IDPrecioSubasta, &p.FechaCreacion, &p.FechaModificacion)

	if error != nil {
		respondJSON(w, 500, map[string]string{"error": error.Error()})
		return
	}
	respondJSON(w, 201, p)
}

// UpdatePrecioSubastaGanado actualiza un precio de subasta ganado existente
func UpdatePrecioSubastaGanado(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.PrecioSubastaGanado
	json.NewDecoder(r.Body).Decode(&p)

	_, err := config.DB.Exec(
		"UPDATE \"tr_precioSubastaGanado\" SET id_subasta = $1, id_categoria_ganado = $2, activo = $3, fecha_modificacion = now() WHERE id_precio_subasta = $4",
		p.IDSubasta, p.IDCategoriaGanado, p.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato actualizado"})
}

// DeletePrecioSubastaGanado elimina un precio de subasta ganado
func DeletePrecioSubastaGanado(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM \"tr_precioSubastaGanado\" WHERE id_precio_subasta = $1", id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}
