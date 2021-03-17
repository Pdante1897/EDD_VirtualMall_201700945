package Dato

type JsonInventario struct {
	Invetarios []struct {
		Tienda       string     `json:"Tienda"`
		Departamento string     `json:"Departamento"`
		Calificacion int        `json:"Calificacion"`
		Productos    []Producto `json:"Productos"`
	} `json:"Invetarios"`
}

type Producto struct {
	Nombre      string  `json:"Nombre"`
	Codigo      int     `json:"Codigo"`
	Descripcion string  `json:"Descripcion"`
	Precio      float64 `json:"Precio"`
	Cantidad    int     `json:"Cantidad"`
	Imagen      string  `json:"Imagen"`
}

type ArbolIn struct {
	Raiz *NodoAI
}
type NodoAI struct {
	Valor  Producto
	Factor int
	Left   *NodoAI
	Right  *NodoAI
}

func NewArbolIn() *ArbolIn {
	return &ArbolIn{nil}
}

func NewNodoIn(valor Producto) *NodoAI {
	return &NodoAI{valor, 0, nil, nil}
}

func rotLLIn(nodo *NodoAI, nodo1 *NodoAI) *NodoAI {
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

func rotRRIn(nodo *NodoAI, nodo1 *NodoAI) *NodoAI {
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

func rotLRIn(nodo *NodoAI, nodo1 *NodoAI) *NodoAI {
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

func rotRLIn(nodo *NodoAI, nodo1 *NodoAI) *NodoAI {
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

func insertProd(raiz *NodoAI, valor Producto, bol *bool) *NodoAI {
	var nodo1 *NodoAI
	if raiz == nil {
		raiz = NewNodoIn(valor)
		*bol = true
	} else if valor.Codigo < raiz.Valor.Codigo {
		izq := insertProd(raiz.Left, valor, bol)
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
					raiz = rotLLIn(raiz, nodo1)
				} else {
					raiz = rotLRIn(raiz, nodo1)
				}
				*bol = false
			}

		}
	} else if valor.Codigo > raiz.Valor.Codigo {
		derc := insertProd(raiz.Right, valor, bol)
		raiz.Right = derc
		if *bol {
			switch raiz.Factor {
			case 1:
				nodo1 = raiz.Right
				if nodo1.Factor == 1 {
					raiz = rotRRIn(raiz, nodo1)
				} else {
					raiz = rotRLIn(raiz, nodo1)
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

func (this *ArbolIn) Insert(valor Producto) {
	bol1 := false
	bol2 := &bol1
	this.Raiz = insertProd(this.Raiz, valor, bol2)
}

func BusquedaArbIn(nodo *NodoAI, Codigo int) *NodoAI {
	var aux *NodoAI
	if Codigo > nodo.Valor.Codigo {
		aux = BusquedaArbIn(nodo.Right, Codigo)
	} else if Codigo < nodo.Valor.Codigo {
		aux = BusquedaArbIn(nodo.Left, Codigo)
	} else if Codigo == nodo.Valor.Codigo {
		aux = nodo
	} else {
		aux = nil
	}

	return aux
}
