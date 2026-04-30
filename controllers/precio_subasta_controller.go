package controllers

import (
	"agropecuario_crud/config"
	"agropecuario_crud/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL (con JOIN)
func GetPreciosSubasta(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT 
			ps.id_precio_subasta,
			ps.id_subasta,
			s.ubicacion,
			ps.id_categoria_ganado,
			c.codigo,
			ps.activo
		FROM agropecuario."tr_precioSubastaGanado" ps
		JOIN agropecuario.subasta s ON ps.id_subasta = s.id_subasta
		JOIN agropecuario.categoria_ganado c ON ps.id_categoria_ganado = c.id_categoria_ganado
	`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		var id, idSubasta, idCategoria int
		var ubicacion, codigo string
		var activo bool

		rows.Scan(&id, &idSubasta, &ubicacion, &idCategoria, &codigo, &activo)

		item := map[string]interface{}{
			"id": id,
			"id_subasta": idSubasta,
			"ubicacion": ubicacion,
			"id_categoria_ganado": idCategoria,
			"categoria_codigo": codigo,
			"activo": activo,
		}

		list = append(list, item)
	}

	respondJSON(w, 200, list)
}

// CREATE (valida relaciones)
func CreatePrecioSubasta(w http.ResponseWriter, r *http.Request) {
	var p models.PrecioSubasta
	json.NewDecoder(r.Body).Decode(&p)

	err := config.DB.QueryRow(`
		INSERT INTO agropecuario."tr_precioSubastaGanado"
		(id_subasta, id_categoria_ganado, activo)
		VALUES ($1, $2, $3)
		RETURNING id_precio_subasta
	`, p.Id_subasta, p.Id_categoria_ganado, true).Scan(&p.Id_precio_subasta)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, p)
}

// DELETE
func DeletePrecioSubasta(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(`
		DELETE FROM agropecuario."tr_precioSubastaGanado"
		WHERE id_precio_subasta = $1
	`, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}