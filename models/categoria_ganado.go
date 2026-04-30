package models

type CategoriaGanado struct {
	Id_categoria_ganado int    `json:"id"`
	Codigo              string `json:"codigo"`
	Descripcion         string `json:"descripcion"`
	Activo              bool   `json:"activo"`
}