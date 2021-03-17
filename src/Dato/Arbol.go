package Dato

type Arbol struct {
	Raiz *NodoA
}
type NodoA struct {
	Valor  ListaDPe
	Factor int
	Left   *NodoA
	Right  *NodoA
}

func NewArbol() *Arbol {
	return &Arbol{nil}
}

func NewNodo(valor ListaDPe) *NodoA {
	return &NodoA{valor, 0, nil, nil}
}

func rotLL(nodo *NodoA, nodo1 *NodoA) *NodoA {
	nodo.Left = nodo1.Right
	nodo1.Right = nodo
	if nodo1.Factor == -1 {
		nodo.Factor = 0
		nodo1.Factor = 0
	} else {
		nodo.Factor = -1
		nodo1.Factor = 1
	}
	return nodo1
}

func rotRR(nodo *NodoA, nodo1 *NodoA) *NodoA {
	nodo.Right = nodo1.Left
	nodo1.Left = nodo
	if nodo1.Factor == 1 {
		nodo.Factor = 0
		nodo1.Factor = 0
	} else {
		nodo.Factor = 1
		nodo1.Factor = -1
	}
	return nodo1
}

func rotLR(nodo *NodoA, nodo1 *NodoA) *NodoA {
	nodo2 := nodo1.Right
	nodo.Left = nodo2.Right
	nodo2.Right = nodo
	nodo1.Right = nodo2.Left
	nodo2.Left = nodo1
	if nodo2.Factor == 1 {
		nodo1.Factor = -1
	} else {
		nodo1.Factor = 0
	}
	if nodo.Factor == -1 {
		nodo.Factor = -1
	} else {
		nodo.Factor = 0
	}
	nodo2.Factor = 0
	return nodo2
}

func rotRL(nodo *NodoA, nodo1 *NodoA) *NodoA {
	nodo2 := nodo1.Left
	nodo.Right = nodo2.Left
	nodo2.Left = nodo
	nodo1.Left = nodo2.Right
	nodo2.Right = nodo1
	if nodo2.Factor == 1 {
		nodo.Factor = -1
	} else {
		nodo.Factor = 0
	}
	if nodo2.Factor == -1 {
		nodo1.Factor = 1
	} else {
		nodo1.Factor = 0
	}
	nodo2.Factor = 0
	return nodo2
}

func insert(raiz *NodoA, valor ListaDPe, bol *bool) *NodoA {
	var nodo1 *NodoA
	if raiz == nil {
		raiz = NewNodo(valor)
		*bol = true
	} else if valor.Indice < raiz.Valor.Indice {
		izq := insert(raiz.Left, valor, bol)
		raiz.Left = izq
		if *bol {
			switch raiz.Factor {
			case 1:
				raiz.Factor = 0
				*bol = false
				break
			case 0:
				raiz.Factor = -1
				break
			case -1:
				nodo1 = raiz.Left
				if nodo1.Factor == -1 {
					raiz = rotLL(raiz, nodo1)
				} else {
					raiz = rotLR(raiz, nodo1)
				}
				*bol = false
			}

		}
	} else if valor.Indice > raiz.Valor.Indice {
		derc := insert(raiz.Right, valor, bol)
		raiz.Right = derc
		if *bol {
			switch raiz.Factor {
			case 1:
				nodo1 = raiz.Right
				if nodo1.Factor == 1 {
					raiz = rotRR(raiz, nodo1)
				} else {
					raiz = rotRL(raiz, nodo1)
				}
				*bol = false
				break
			case 0:
				raiz.Factor = 1
				break
			case -1:
				raiz.Factor = 0
				*bol = false
			}
		}
	}
	return raiz
}

func (this *Arbol) Insert(valor ListaDPe) {
	bol1 := false
	bol2 := &bol1
	this.Raiz = insert(this.Raiz, valor, bol2)
}

func BusquedaArb(nodo *NodoA, anio int) *NodoA {
	var aux *NodoA
	if anio > nodo.Valor.Indice {
		aux = BusquedaArb(nodo.Right, anio)
	} else if anio < nodo.Valor.Indice {
		aux = BusquedaArb(nodo.Left, anio)
	} else if anio == nodo.Valor.Indice {
		aux = nodo
	} else {
		aux = nil
	}

	return aux
}
