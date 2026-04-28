package controllers

import (
	"encoding/json"
	"net/http"

	"agropecuario_crud/config"
	"agropecuario_crud/models"
)

func GetRelaciones(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT id_precio_subasta, id_subasta, id_categoria_ganado, activo
	FROM agropecuario.Tr_PrecioSubastaGanado
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var relaciones []models.PrecioSubastaGanado

	for rows.Next() {
		var rel models.PrecioSubastaGanado
		rows.Scan(&rel.ID, &rel.IdSubasta, &rel.IdCategoriaGanado, &rel.Activo)
		relaciones = append(relaciones, rel)
	}

	json.NewEncoder(w).Encode(relaciones)
}

func CreateRelacion(w http.ResponseWriter, r *http.Request) {
	var rel models.PrecioSubastaGanado
	json.NewDecoder(r.Body).Decode(&rel)

	query := `
	INSERT INTO agropecuario.Tr_PrecioSubastaGanado
	(id_subasta, id_categoria_ganado, activo)
	VALUES ($1, $2, $3)
	`

	_, err := config.DB.Exec(query, rel.IdSubasta, rel.IdCategoriaGanado, rel.Activo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "relación creada"})
}