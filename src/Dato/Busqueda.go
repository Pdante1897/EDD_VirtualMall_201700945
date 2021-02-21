package Dato

import (
	"fmt"
	"strconv"
)

type Busqueda struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}
type Eliminar struct {
	Nombre       string `json:"Nombre"`
	Categoria    string `json:"Categoria"`
	Calificacion int    `json:"Calificacion"`
}

func RowMajor(indices int, departamentos int, posiciones [2]int, calificacion int) int {
	var primero int
	var segundo int
	var tercero int
	primero = posiciones[0]
	segundo = (primero * departamentos) + posiciones[1]
	tercero = (segundo * 5) + calificacion
	fmt.Println(strconv.Itoa(indices) + " " + strconv.Itoa(departamentos) + " " + strconv.Itoa(posiciones[0]) + " " + strconv.Itoa(posiciones[1]) + " " + strconv.Itoa(tercero))
	return tercero
}
func (this Eliminar) NumDep(lista []ListaDoble) [2]int {
	var contador [2]int
	var temp string
	var temp1 string
	contador[0] = 0
	contador[1] = 0
	temp = string(this.Nombre[0])
	for i := 0; i < len(lista); i++ {
		temp1 = lista[i+1].Indice
		if temp1 != lista[i].Indice {
			contador[0]++
		}
		if temp == lista[i].Indice {
			break
		}

	}
	for i := 0; i < len(lista); i++ {
		if this.Categoria != lista[i].Nombre {
			contador[1]++
			i = i + 4
		} else if this.Categoria == lista[i].Nombre {
			break
		}
	}
	return contador
}
func (this Busqueda) NumDep(lista []ListaDoble) [2]int {
	var contador [2]int
	var temp string
	var temp1 string
	contador[0] = 0
	contador[1] = 0
	temp = string(this.Nombre[0])
	for i := 0; i < len(lista); i++ {
		temp1 = lista[i+1].Indice
		if temp1 != lista[i].Indice {
			contador[0]++
		}
		if temp == lista[i].Indice {
			break
		}

	}
	for i := 0; i < len(lista); i++ {
		if this.Departamento != lista[i].Nombre {
			contador[1]++
			i = i + 4
		} else if this.Departamento == lista[i].Nombre {
			break
		}
	}
	return contador
}
