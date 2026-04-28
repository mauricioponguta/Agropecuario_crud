package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"agropecuario_crud/config"
	"agropecuario_crud/models"

	"github.com/gorilla/mux"
)

func GetSubastas(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT id_subasta, fecha_subasta, ubicacion, precio, activo 
	FROM agropecuario.Subasta
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var subastas []models.Subasta

	for rows.Next() {
		var s models.Subasta
		rows.Scan(&s.ID, &s.Fecha, &s.Ubicacion, &s.Precio, &s.Activo)
		subastas = append(subastas, s)
	}

	json.NewEncoder(w).Encode(subastas)
}

func CreateSubasta(w http.ResponseWriter, r *http.Request) {
	var s models.Subasta
	json.NewDecoder(r.Body).Decode(&s)

	query := `
	INSERT INTO agropecuario.Subasta 
	(fecha_subasta, ubicacion, precio, activo)
	VALUES ($1, $2, $3, $4) RETURNING id_subasta
	`

	err := config.DB.QueryRow(query, s.Fecha, s.Ubicacion, s.Precio, s.Activo).Scan(&s.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(s)
}

func UpdateSubasta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var s models.Subasta
	json.NewDecoder(r.Body).Decode(&s)

	query := `
	UPDATE agropecuario.Subasta
	SET fecha_subasta=$1, ubicacion=$2, precio=$3, activo=$4
	WHERE id_subasta=$5
	`

	_, err := config.DB.Exec(query, s.Fecha, s.Ubicacion, s.Precio, s.Activo, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "actualizado"})
}

func DeleteSubasta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	query := `DELETE FROM agropecuario.Subasta WHERE id_subasta=$1`

	_, err := config.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "eliminado"})
}