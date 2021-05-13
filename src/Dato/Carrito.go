package Dato

type JsonPedido struct {
	Cliente   string        `json:"Cliente"`
	Anio      int           `json:"Anio"`
	Mes       int           `json:"Mes"`
	Dia       int           `json:"Dia"`
	Productos []ProductoPed `json:"Productos"`
	Total     float64       `json:"Total"`
}

type ProductoPed struct {
	Tienda       string  `json:"Tienda"`
	Departamento string  `json:"Departamento"`
	Calificacion string  `json:"Calificacion"`
	Nombre       string  `json:"Nombre"`
	Codigo       int     `json:"Codigo"`
	Descripcion  string  `json:"Descripcion"`
	Precio       float64 `json:"Precio"`
	Cantidad     int     `json:"Cantidad"`
	Imagen       string  `json:"Imagen"`
}
