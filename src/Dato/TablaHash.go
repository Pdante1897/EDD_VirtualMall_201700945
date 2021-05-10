package Dato

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type NodoHash struct {
	Hash           int
	Valor          string
	SubComentarios *TablaHash
}

type TablaHash struct {
	Size            int
	Carga           int
	Porcentaje      int
	Porcentaje_crec int
	arreglo         []*NodoHash
}

func (this *TablaHash) Posicion(clave int, valor string) int {

	i, p := 0, 0
	d := 0.2520*float64(clave) - float64(int(0.2520*float64(clave)))
	p = int(float64(len(this.arreglo)) * d)
	fmt.Println(this.Size)
	fmt.Println(this.Size)
	fmt.Println(p)
	//p = int(clave % this.Size)
	for this.arreglo[p] != nil && this.arreglo[p].Valor != valor {
		i++
		j := i * i
		p = p + j
		if p >= len(this.arreglo) {
			p = p - len(this.arreglo)
		}
	}
	return p
}

func (this *TablaHash) Insertar(nuevo int, valor string) {
	nuevoNodo := NodoHash{nuevo, valor, nil}
	pos := this.Posicion(nuevo, valor)
	this.arreglo[pos] = &nuevoNodo
	this.Carga++
	if ((this.Carga * 100) / this.Size) > this.Porcentaje {
		sizenuevo := this.Size
		for {
			sizenuevo++
			if ((this.Carga * 100) / sizenuevo) <= this.Porcentaje_crec {
				break
			}
		}
		nuevoArray := make([]*NodoHash, sizenuevo)
		viejo := this.arreglo
		this.arreglo = nuevoArray
		this.Size = sizenuevo
		aux := 0
		for i := 0; i < len(viejo); i++ {
			if viejo[i] != nil {
				aux = this.Posicion(viejo[i].Hash, valor)
				nuevoArray[aux] = viejo[i]
			}
		}
	}
}

func (this *TablaHash) Imprimir() {
	data := make([][]string, this.Size)
	for i := 0; i < len(this.arreglo); i++ {
		temp := make([]string, 2)
		aux := this.arreglo[i]
		if aux != nil {
			temp[0] = strconv.Itoa(aux.Hash)
			temp[1] = aux.Valor
		} else {
			temp[0] = "-"
			temp[1] = "-"
		}
		data[i] = temp
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hash", "Valor"})
	table.SetFooter([]string{"size", strconv.Itoa(this.Size), "Carga", strconv.Itoa(this.Carga)})
	table.AppendBulk(data)
	table.Render()

}

func NewTablaHash(size int, porcentaje int, porcentaje_crec int) *TablaHash {
	arreglo := make([]*NodoHash, size)
	return &TablaHash{size, 0, porcentaje, porcentaje_crec, arreglo}
}

func (this *TablaHash) ListaComent() []Comentario {
	var comentarios []Comentario
	for i := 0; i < len(this.arreglo); i++ {
		if this.arreglo[i] != nil {
			var aux Comentario
			aux.Id = this.arreglo[i].Hash
			aux.Cadena = this.arreglo[i].Valor
			comentarios = append(comentarios, aux)
		}

	}
	return comentarios
}
