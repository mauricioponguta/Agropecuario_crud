package models

type PrecioSubasta struct {
	Id_precio_subasta   int  `json:"id"`
	Id_subasta          int  `json:"id_subasta"`
	Id_categoria_ganado int  `json:"id_categoria_ganado"`
	Activo              bool `json:"activo"`
}