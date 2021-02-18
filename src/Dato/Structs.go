package Dato

import (
	"fmt"
	"strconv"
)

type ArchivoJson struct {
	Datos []struct {
		Indice        string `json:"Indice"`
		Departamentos []struct {
			Nombre  string   `json:"Nombre"`
			Tiendas []Tienda `json:"Tiendas"`
		} `json:"Departamentos"`
	} `json:"Datos"`
}
type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
}

func (this Tienda) ToString() string {
	return ("Nombre: " + this.Nombre + " Descripcion: " + this.Descripcion + " Contacto: " + this.Contacto + " Calificacion: " + strconv.Itoa(this.Calificacion))
}

func Hola() {
	fmt.Println("hola mundo")

}

type Nodo struct {
	Tienda    Tienda
	Anterior  *Nodo
	Siguiente *Nodo
}
type ListaDoble struct {
	Indice       string
	Nombre       string
	Calificacion int
	Inicio       *Nodo
	Fin          *Nodo
}

func (this *ListaDoble) Vacio() bool {
	if this.Inicio == nil {
		return true
	} else {
		return false
	}

}
func (this *ListaDoble) Insertar(dato Tienda) {
	aux := Nodo{
		Tienda:    dato,
		Siguiente: nil,
		Anterior:  nil,
	}
	if this.Vacio() {
		this.Inicio = &aux
		this.Fin = this.Inicio
	} else {
		aux.Anterior = this.Fin
		this.Fin.Siguiente = &aux
		this.Fin = this.Fin.Siguiente

	}
}

func (this ListaDoble) To_String() {
	var aux *Nodo
	aux = this.Inicio
	fmt.Println("__________________________________________________________________________________________________")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println(this.Indice + "\n" + this.Nombre + "\n" + strconv.Itoa(this.Calificacion))
	fmt.Println("-----------------------------------------------------------------------")

	for {
		if this.Vacio() {
			break
		}
		if aux == this.Inicio {
			fmt.Println(aux.Tienda.ToString())
			aux = aux.Siguiente
		} else if aux != nil {
			fmt.Println(aux.Tienda.ToString())
			aux = aux.Siguiente
		} else if aux == this.Fin {
			fmt.Println(aux.Tienda.ToString())
			break
		} else if aux == this.Inicio {
			break
		} else {
			break
		}

	}
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("__________________________________________________________________________________________________")

}
