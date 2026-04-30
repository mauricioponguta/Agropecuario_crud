package controllers

import (
	"agropecuario_crud/config"
	"agropecuario_crud/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Helper respuesta JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// GET ALL
func GetCategoriasGanado(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_categoria_ganado, codigo, descripcion, activo
		FROM agropecuario.categoria_ganado
	`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.CategoriaGanado

	for rows.Next() {
		var c models.CategoriaGanado
		rows.Scan(&c.Id_categoria_ganado, &c.Codigo, &c.Descripcion, &c.Activo)
		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

// GET BY ID
func GetCategoriaGanadoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c models.CategoriaGanado

	err := config.DB.QueryRow(`
		SELECT id_categoria_ganado, codigo, descripcion, activo
		FROM agropecuario.categoria_ganado
		WHERE id_categoria_ganado = $1
	`, id).Scan(&c.Id_categoria_ganado, &c.Codigo, &c.Descripcion, &c.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Id no encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

// CREATE
func CreateCategoriaGanado(w http.ResponseWriter, r *http.Request) {
	var c models.CategoriaGanado
	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(`
		INSERT INTO agropecuario.categoria_ganado (codigo, descripcion, activo)
		VALUES ($1, $2, $3)
		RETURNING id_categoria_ganado
	`, c.Codigo, c.Descripcion, true).Scan(&c.Id_categoria_ganado)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

// UPDATE
func UpdateCategoriaGanado(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c models.CategoriaGanado
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(`
		UPDATE agropecuario.categoria_ganado
		SET codigo = $1, descripcion = $2
		WHERE id_categoria_ganado = $3
	`, c.Codigo, c.Descripcion, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato actualizado"})
}

// DELETE
func DeleteCategoriaGanado(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(`
		DELETE FROM agropecuario.categoria_ganado
		WHERE id_categoria_ganado = $1
	`, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}