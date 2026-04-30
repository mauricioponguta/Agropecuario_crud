package models

type Subasta struct {
	Id_subasta          int     `json:"id"`
	Fecha_subasta       string  `json:"fecha_subasta"`
	Ubicacion           string  `json:"ubicacion"`
	Precio              float64 `json:"precio"`
	Id_usuario_registra int     `json:"id_usuario_registra"`
	Activo              bool    `json:"activo"`
}