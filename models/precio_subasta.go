package models

import "time"

type PrecioSubastaGanado struct {
	IDPrecioSubasta   int       `json:"id_precio_subasta"`
	IDSubasta         int       `json:"id_subasta"`
	IDCategoriaGanado int       `json:"id_categoria_ganado"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
