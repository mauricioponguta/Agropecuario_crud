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