package controllers

import (
	"agropecuario_crud/config"
	"agropecuario_crud/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL
func GetSubastas(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_subasta, fecha_subasta, ubicacion, precio, id_usuario_registra, activo
		FROM agropecuario.subasta
	`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Subasta

	for rows.Next() {
		var s models.Subasta
		rows.Scan(&s.Id_subasta, &s.Fecha_subasta, &s.Ubicacion, &s.Precio, &s.Id_usuario_registra, &s.Activo)
		list = append(list, s)
	}

	respondJSON(w, 200, list)
}

// GET BY ID
func GetSubastaByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var s models.Subasta

	err := config.DB.QueryRow(`
		SELECT id_subasta, fecha_subasta, ubicacion, precio, id_usuario_registra, activo
		FROM agropecuario.subasta
		WHERE id_subasta = $1
	`, id).Scan(&s.Id_subasta, &s.Fecha_subasta, &s.Ubicacion, &s.Precio, &s.Id_usuario_registra, &s.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Id no encontrado"})
		return
	}

	respondJSON(w, 200, s)
}

// CREATE
func CreateSubasta(w http.ResponseWriter, r *http.Request) {
	var s models.Subasta
	json.NewDecoder(r.Body).Decode(&s)

	err := config.DB.QueryRow(`
		INSERT INTO agropecuario.subasta
		(fecha_subasta, ubicacion, precio, id_usuario_registra, activo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id_subasta
	`, s.Fecha_subasta, s.Ubicacion, s.Precio, s.Id_usuario_registra, true).Scan(&s.Id_subasta)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, s)
}

// UPDATE
func UpdateSubasta(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var s models.Subasta
	json.NewDecoder(r.Body).Decode(&s)

	_, err := config.DB.Exec(`
		UPDATE agropecuario.subasta
		SET fecha_subasta = $1, ubicacion = $2, precio = $3, id_usuario_registra = $4
		WHERE id_subasta = $5
	`, s.Fecha_subasta, s.Ubicacion, s.Precio, s.Id_usuario_registra, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato actualizado"})
}

// DELETE
func DeleteSubasta(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(`
		DELETE FROM agropecuario.subasta
		WHERE id_subasta = $1
	`, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}