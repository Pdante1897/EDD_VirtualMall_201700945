package Dato

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type NodoHash struct {
	Id             int
	Hash           int
	Valor          string
	SubComentarios *TablaHash
}

type TablaHash struct {
	Size            int
	Carga           int
	Porcentaje      int
	Porcentaje_crec int
	Arreglo         []*NodoHash
}

func (this *TablaHash) Posicion(clave int, valor string) int {

	i, p := 0, 0
	d := 0.2520*float64(clave) - float64(int(0.2520*float64(clave)))
	p = int(float64(len(this.Arreglo)) * d)
	fmt.Println(this.Size)
	fmt.Println(this.Size)
	fmt.Println(p)
	//p = int(clave % this.Size)
	for this.Arreglo[p] != nil && this.Arreglo[p].Valor != valor {
		i++
		j := i * i
		p = p + j
		for p >= len(this.Arreglo) {
			p = p - len(this.Arreglo)
		}

	}
	return p
}
func isPrime(candidate int64) bool {
	var i, limit int64

	if candidate == 2 {
		return true
	}

	if math.Mod(float64(candidate), 2) == 0 {
		return false
	}

	limit = int64(math.Ceil(math.Sqrt(float64(candidate))))
	for i = 3; i <= limit; i += 2 { //Only odd numbers
		if math.Mod(float64(candidate), float64(i)) == 0 {
			return false
		}
	}
	return true
}
func (this *TablaHash) Insertar(id int, nuevo int, valor string) {
	nuevoNodo := NodoHash{0, nuevo, valor, nil}
	pos := this.Posicion(nuevo, valor)
	this.Arreglo[pos] = &nuevoNodo
	this.Carga++
	if ((this.Carga * 100) / this.Size) > this.Porcentaje {
		sizenuevo := this.Size + 1
		for !isPrime(int64(sizenuevo)) {
			sizenuevo++

			fmt.Println(sizenuevo)
		}
		nuevoArray := make([]*NodoHash, sizenuevo)
		viejo := this.Arreglo
		this.Arreglo = nuevoArray
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
	for i := 0; i < len(this.Arreglo); i++ {
		temp := make([]string, 2)
		aux := this.Arreglo[i]
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
	j := 0
	for i := 0; i < len(this.Arreglo); i++ {
		if this.Arreglo[i] != nil {
			var aux Comentario
			this.Arreglo[i].Id = j
			aux.Id = this.Arreglo[i].Id
			aux.Dpi = this.Arreglo[i].Hash
			aux.Cadena = this.Arreglo[i].Valor
			if this.Arreglo[i].SubComentarios != nil {
				aux.SubComentarios = this.Arreglo[i].SubComentarios.ListaComent()

			}
			comentarios = append(comentarios, aux)

			j++

		}

	}
	return comentarios
}
