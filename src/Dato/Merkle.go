package Dato

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Merkle struct {
	Raiz *NodoMerkle
}
type NodoMerkle struct {
	Valor  int
	Hash   string
	Cadena string
	Left   *NodoMerkle
	Right  *NodoMerkle
}

func (this *NodoMerkle) Suma() int {
	if this.Right != nil && this.Left != nil {
		return this.Left.Valor + this.Right.Valor
	}
	return -1
}

func BusquedaMk(nodo *NodoMerkle, Codigo string) *NodoMerkle {
	if nodo == nil {
		return nil
	}
	var aux *NodoMerkle
	if Codigo == nodo.Cadena {
		aux = nodo
	} else if nodo.Left != nil {
		aux = BusquedaMk(nodo.Left, Codigo)
	} else if nodo.Right != nil {
		aux = BusquedaMk(nodo.Right, Codigo)
	}

	return aux
}
func NewMerkel() *Merkle {
	return &Merkle{}
}
func NewNodoMk(valor int, cadena string, right *NodoMerkle, left *NodoMerkle) *NodoMerkle {
	nodo := NodoMerkle{valor, "", cadena, left, right}
	return &nodo
}

func (this *Merkle) Insertar(valor int, cadena string) {
	n := NewNodoMk(valor, cadena, nil, nil)
	if this.Raiz == nil {
		lista := list.New()
		lista.PushBack(n)
		lista.PushBack(NewNodoMk(-1, "-1", nil, nil))
		this.construirArbol(lista)
	} else {
		lista := this.GetLista()
		lista.PushBack(n)
		this.construirArbol(lista)
	}
}

func (this *Merkle) GetLista() *list.List {
	lista := list.New()
	getLista(lista, this.Raiz.Left)
	getLista(lista, this.Raiz.Right)
	return lista
}

func getLista(lista *list.List, actual *NodoMerkle) {
	if actual != nil {
		getLista(lista, actual.Left)
		if actual.Right == nil && actual.Valor != -1 {
			lista.PushBack(actual)
		}
		getLista(lista, actual.Right)
	}
}

func (this *Merkle) construirArbol(lista *list.List) {
	size := float64(lista.Len())
	cant := 1
	for (size)/2 > 1 {
		cant++
		size = size / 2
	}
	nodosT := math.Pow(2, float64(cant))
	for lista.Len() < int(nodosT) {
		lista.PushBack(NewNodoMk(-1, "-1", nil, nil))

	}
	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoMerkle)
		nodo2 := segundo.Value.(*NodoMerkle)
		nuevo := NewNodoMk(nodo1.Valor+nodo2.Valor, nodo1.Cadena+nodo2.Cadena, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*NodoMerkle)
}

func GenerarHashArbol(nodo *NodoMerkle) *NodoMerkle {
	if nodo == nil {
		return nil
	}
	var aux *NodoMerkle
	if nodo.Left == nil && nodo.Right == nil {
		aux = nodo
	}
	if nodo.Left != nil {
		aux = GenerarHashArbol(nodo.Left)
		nodo.Cadena = nodo.Left.Hash
	}
	if nodo.Right != nil {
		aux = GenerarHashArbol(nodo.Right)

		nodo.Cadena = nodo.Cadena + nodo.Right.Hash
	}
	//value := sha256.Sum256([]byte(nodo.Cadena))
	var cad strings.Builder
	arreglo := sha256.Sum256([]byte(nodo.Cadena))
	fmt.Fprintf(&cad, "%x", arreglo[:])
	fmt.Println(cad.String())
	nodo.Hash = cad.String()
	cad.Reset()

	return aux
}

func (this *NodoMerkle) GenerarGraphvizMerk() string {
	var cadena = ""
	var dir strings.Builder
	fmt.Fprintf(&dir, "%x", &this)
	var dirl strings.Builder
	fmt.Fprintf(&dirl, "%x", &this.Left)
	var dirr strings.Builder
	fmt.Fprintf(&dirr, "%x", &this.Right)
	if this.Right == nil && this.Left == nil {
		cadena = "node" + dir.String() + "[label=\"{" + strconv.Itoa(this.Valor) + " | " + this.Cadena + " | " + this.Hash + "}\"];\n"
	} else {
		cadena = "node" + dir.String() + "[label=\"{" + strconv.Itoa(this.Valor) + " | " + this.Cadena + " | " + this.Hash + "}\"];\n"
	}
	if this.Left != nil {
		cadena += this.Left.GenerarGraphvizMerk() + "node" + dir.String() + "->node" + dirl.String() + ";\n"
	}
	if this.Right != nil {
		cadena += this.Right.GenerarGraphvizMerk() + "node" + dir.String() + "->node" + dirr.String() + ";\n"
	}
	dir.Reset()
	dirl.Reset()
	dirr.Reset()

	return cadena
}
