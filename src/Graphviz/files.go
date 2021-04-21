package Graphviz

type ListaPdf struct {
	Pdfs []Archivopdf `json:"Archivos"`
}

type Archivopdf struct {
	Nombre string `json:"Nombre"`
	Ruta   string `json:"Ruta"`
}

func NewListaPDF() ListaPdf {
	lista := ListaPdf{nil}
	return lista
}
