package Dato

type NodoC struct {
	Producto  Producto
	siguiente *NodoC
}

type Cola struct {
	inicio *NodoC
	fin    *NodoC
}

func (this Cola) vacia() bool {
	if this.inicio == nil {
		return true
	} else {
		return false
	}
}

func (this *Cola) push(dato Producto) {
	auxiliar := new(NodoC)
	auxiliar2 := new(NodoC)
	auxiliar.Producto = dato
	if this.vacia() {
		this.inicio = auxiliar
		this.fin = auxiliar
	} else if this.inicio.siguiente == nil {
		auxiliar2 = this.inicio
		this.inicio = auxiliar
		this.fin = auxiliar2
		this.inicio.siguiente = this.fin
	} else {
		auxiliar2 = this.inicio
		this.inicio = auxiliar
		this.inicio.siguiente = auxiliar2
	}
}

func (this *Cola) pop() Producto {
	var dato Producto
	if !this.vacia() {
		auxiliar := new(NodoC)
		dato = this.inicio.Producto
		if this.inicio.siguiente != nil {
			auxiliar = this.inicio.siguiente
			this.inicio = auxiliar
		} else {

		}
		return dato
	} else {
		return dato
	}

}
