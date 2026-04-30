package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"agropecuario_crud/config"
	"agropecuario_crud/models"

	"github.com/gorilla/mux"
)


// GET ALL
func GetSubastas(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT id_subasta, fecha_subasta, ubicacion, precio,
	       id_usuario_registra, activo, fecha_creacion, fecha_modificacion
	FROM agropecuario.subasta
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
		rows.Scan(
			&s.ID,
			&s.Fecha,
			&s.Ubicacion,
			&s.Precio,
			&s.IdUsuarioRegistra,
			&s.Activo,
			&s.FechaCreacion,
			&s.FechaModificacion,
		)
		subastas = append(subastas, s)
	}

	json.NewEncoder(w).Encode(subastas)
}


// GET BY ID
func GetSubastaByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	query := `
	SELECT id_subasta, fecha_subasta, ubicacion, precio,
	       id_usuario_registra, activo, fecha_creacion, fecha_modificacion
	FROM agropecuario.subasta
	WHERE id_subasta = $1
	`

	var s models.Subasta

	err := config.DB.QueryRow(query, id).Scan(
		&s.ID,
		&s.Fecha,
		&s.Ubicacion,
		&s.Precio,
		&s.IdUsuarioRegistra,
		&s.Activo,
		&s.FechaCreacion,
		&s.FechaModificacion,
	)

	if err != nil {
		http.Error(w, "No encontrado", 404)
		return
	}

	json.NewEncoder(w).Encode(s)
}


// CREATE
func CreateSubasta(w http.ResponseWriter, r *http.Request) {
	var s models.Subasta

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	query := `
	INSERT INTO agropecuario.subasta 
	(fecha_subasta, ubicacion, precio, id_usuario_registra, activo)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id_subasta
	`

	err = config.DB.QueryRow(
		query,
		s.Fecha,
		s.Ubicacion,
		s.Precio,
		s.IdUsuarioRegistra,
		s.Activo,
	).Scan(&s.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(s)
}


// UPDATE
func UpdateSubasta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var s models.Subasta
	json.NewDecoder(r.Body).Decode(&s)

	query := `
	UPDATE agropecuario.subasta
	SET fecha_subasta=$1,
	    ubicacion=$2,
	    precio=$3,
	    id_usuario_registra=$4,
	    activo=$5,
	    fecha_modificacion=now()
	WHERE id_subasta=$6
	`

	_, err := config.DB.Exec(
		query,
		s.Fecha,
		s.Ubicacion,
		s.Precio,
		s.IdUsuarioRegistra,
		s.Activo,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "actualizado"})
}


// DELETE (físico)
func DeleteSubasta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	query := `DELETE FROM agropecuario.subasta WHERE id_subasta=$1`

	_, err := config.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "eliminado"})
}