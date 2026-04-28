package models

type Subasta struct {
	ID        int     `json:"id_subasta"`
	Fecha     string  `json:"fecha_subasta"`
	Ubicacion string  `json:"ubicacion"`
	Precio    float64 `json:"precio"`
	Activo    bool    `json:"activo"`
}