package Dato

import (
	"fmt"
)

type JsonPedidos struct {
	Pedidos []Pedido `json:"Pedidos"`
}
type Pedido struct {
	Fecha        string `json:"Fecha"`
	Tienda       string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int    `json:"Calificacion"`
	Productos    []struct {
		Codigo int `json:"Codigo"`
	} `json:"Productos"`
}

type NodoP struct {
	Mes       string
	MatrizD   Matriz
	Anterior  *NodoP
	Siguiente *NodoP
}
type ListaDPe struct {
	Indice int
	Inicio *NodoP
	Fin    *NodoP
}

func (this *ListaDPe) Vacio() bool {
	if this.Inicio == nil {
		return true
	} else {
		return false
	}

}
func (this *ListaDPe) Insertar(dato string) {
	aux := NodoP{
		Mes:       dato,
		Siguiente: nil,
		Anterior:  nil,
	}
	if this.Vacio() {
		this.Inicio = &aux
		this.Fin = this.Inicio
	} else if this.Inicio.Siguiente == nil {
		aux.Anterior = this.Inicio
		this.Fin = &aux
		this.Inicio.Siguiente = this.Fin
	} else {
		aux.Anterior = this.Fin
		this.Fin.Siguiente = &aux
		this.Fin = this.Fin.Siguiente

	}
}

func (this ListaDPe) Buscar(cadena string) *NodoP {
	var auxiliar *NodoP
	auxiliar = this.Inicio
	fmt.Println("_____________________")
	for auxiliar != nil {
		if auxiliar.Mes == cadena {
			fmt.Println(auxiliar.Mes)
			return auxiliar
		}
		auxiliar = auxiliar.Siguiente

	}
	fmt.Println("no se pudo encontrar")
	return nil
}
func (this *ListaDPe) Eliminar(dato string) {
	var auxiliar1, actual, siguiente *NodoP
	auxiliar1 = this.Buscar(dato)
	actual = this.Inicio

	for {
		if auxiliar1 == this.Inicio && this.Inicio.Siguiente == nil {
			this.Inicio = nil
			break
		} else if auxiliar1 == this.Inicio {
			auxiliar1 = this.Inicio.Siguiente
			this.Inicio = nil
			this.Inicio = auxiliar1
			this.Inicio.Anterior = nil
			break
		} else if auxiliar1 == actual.Siguiente {
			siguiente = auxiliar1.Siguiente
			this.Fin = actual
			this.Fin.Siguiente = siguiente
			fmt.Println("simon")
			break
		}

		if actual == this.Fin {
			break
		}
		actual = actual.Siguiente

	}
}

func (this ListaDPe) To_String() {
	var aux *NodoP
	aux = this.Inicio
	fmt.Println("__________________________________________________________________________________________________")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println(this.Indice)
	fmt.Println("-----------------------------------------------------------------------")

	for {
		if this.Vacio() {
			break
		}
		if aux == this.Inicio {
			fmt.Println(aux.Mes)
			aux = aux.Siguiente
		} else if aux != nil {
			fmt.Println(aux.Mes)
			aux = aux.Siguiente
		} else if aux == this.Fin {
			fmt.Println(aux.Mes)
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
