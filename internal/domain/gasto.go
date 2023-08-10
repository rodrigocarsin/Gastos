package domain

type Gasto struct {
	ID          int64   `json:"id"`
	Descripcion string  `json:"descripcion"`
	Monto       float64 `json:"monto"`
	Fecha       string  `json:"fecha"`
	Categoria   string  `json:"categoria"`
	TipoPago    string  `json:"tipoPago"`
	Comercio    string  `json:"comercio"`
}
