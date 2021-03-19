package Dato

import "fmt"

type NodoL struct {
	Dato      ListaDPe
	Siguiente *NodoL
}

type ListaE struct {
	Inicio *NodoL
	Fin    *NodoL
}

func (lista *ListaE) Insertar(dato ListaDPe) {
	aux := NodoL{
		Dato:      dato,
		Siguiente: nil,
	}
	if lista.Inicio == nil {
		lista.Inicio = &aux
		lista.Fin = lista.Inicio
	} else {
		lista.Fin.Siguiente = &aux
		lista.Fin = lista.Fin.Siguiente
	}
}
func (l *ListaE) Mostrar() {
	var auxiliar *NodoL
	auxiliar = l.Inicio

	for auxiliar != nil {
		fmt.Println(auxiliar.Dato.Indice)
		auxiliar = auxiliar.Siguiente

	}
}
func (l *ListaE) Buscar(dato int) *NodoL {
	if l.Inicio == nil {
		return nil
	}
	var auxiliar *NodoL
	auxiliar = l.Inicio

	for auxiliar != nil {
		if auxiliar.Dato.Indice == dato {
			fmt.Println(auxiliar.Dato.Indice)
			return auxiliar
		}
		auxiliar = auxiliar.Siguiente

	}
	var null *NodoL
	return null
}
func (l *ListaE) Eliminar(dato int) {
	var auxiliar1, actual, siguiente *NodoL
	auxiliar1 = l.Buscar(dato)
	actual = l.Inicio

	for {
		if auxiliar1 == l.Inicio && l.Inicio.Siguiente == nil {
			l.Inicio = nil
			break
		} else if auxiliar1 == l.Inicio {
			auxiliar1 = l.Inicio.Siguiente
			l.Inicio = nil
			l.Inicio = auxiliar1
			break
		} else if auxiliar1 == actual.Siguiente {
			siguiente = auxiliar1.Siguiente
			l.Fin = actual
			l.Fin.Siguiente = siguiente
			fmt.Println("simon")
			break
		}

		if actual == l.Fin {
			break
		}
		actual = actual.Siguiente

	}
}
