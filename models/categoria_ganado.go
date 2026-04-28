package models

type CategoriaGanado struct {
	ID          int    `json:"id_categoria_ganado"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Activo      bool   `json:"activo"`
}