package models

type PrecioSubastaGanado struct {
	ID                  int `json:"id_precio_subasta"`
	IdSubasta           int `json:"id_subasta"`
	IdCategoriaGanado   int `json:"id_categoria_ganado"`
	Activo              bool `json:"activo"`
}