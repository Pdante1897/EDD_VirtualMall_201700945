package Dato

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
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
	var cad strings.Builder
	arreglo := sha256.Sum256([]byte(nodo.Cadena))
	fmt.Fprintf(&cad, "%x", arreglo[:])
	nodo.Hash = cad.String()
	cad.Reset()

	return aux
}

func GetCadenas(this *NodoMerkle) string {
	if this == nil {
		return ""
	}
	var aux strings.Builder

	if this.Left == nil && this.Right == nil && this.Valor != -1 {
		fmt.Fprintf(&aux, "%s", this.Cadena)
	}
	if this.Left != nil {
		fmt.Fprintf(&aux, "%s", GetCadenas(this.Left))
	}
	if this.Right != nil {
		fmt.Fprintf(&aux, "%s", GetCadenas(this.Right))

	}

	return aux.String()
}

func (this *NodoMerkle) GenerarGraphvizMerk() string {
	var cadena = ""
	var mem strings.Builder
	var memH strings.Builder

	if this.Right == nil && this.Left == nil {
		fmt.Fprintf(&mem, "%v", &this.Hash)
		cad1 := mem.String()
		mem.Reset()
		cadena = "node" + cad1 + "[style=\"filled\"; fillcolor=\"green\" color=\"black\"; label=\"{" + strconv.Itoa(this.Valor) + " | " + this.Hash + " | " + this.Cadena + "}\"];\n"
	} else {
		fmt.Fprintf(&mem, "%v", &this.Hash)
		cad1 := mem.String()
		mem.Reset()
		cadena = "node" + cad1 + "[style=\"filled\"; fillcolor=\"green\" color=\"black\"; label=\"{" + strconv.Itoa(this.Valor) + " | " + this.Hash + " | " + this.Cadena + "}\"];\n"
	}
	if this.Left != nil {
		fmt.Fprintf(&memH, "%v", &this.Left.Hash)
		fmt.Fprintf(&mem, "%v", &this.Hash)
		cad1 := mem.String()
		cad2 := memH.String()
		mem.Reset()
		memH.Reset()
		cadena += this.Left.GenerarGraphvizMerk() + "node" + cad1 + "->node" + cad2 + ";\n"
	}
	if this.Right != nil {
		fmt.Fprintf(&memH, "%v", &this.Right.Hash)
		fmt.Fprintf(&mem, "%v", &this.Hash)
		cad1 := mem.String()
		cad2 := memH.String()
		mem.Reset()
		memH.Reset()
		cadena += this.Right.GenerarGraphvizMerk() + "node" + cad1 + "->node" + cad2 + ";\n"
	}
	mem.Reset()
	memH.Reset()

	return cadena
}

func (this *NodoMerkle) Graph() string {
	var cad strings.Builder
	fmt.Fprintf(&cad, "digraph G{ \n")
	fmt.Fprintf(&cad, "node[shape=\"record\"] \n")
	fmt.Fprintf(&cad, this.GenerarGraphvizMerk())
	fmt.Fprintf(&cad, "}")
	return cad.String()
}

func DirSvc(num string) {
	dir, err := filepath.Abs(filepath.Dir("./graphviz/graphviz.go"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "cd c:\\program files\\graphviz\\bin\n  ")
	fmt.Fprintf(&cadena, "dot -Tsvg \""+dir+"\\files\\"+num+".dot\" -o \""+dir+"\\files\\grafica"+num+".svg\"\n  ")
	fil, err := os.Create(dir + "\\files\\archivo.bat")
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := fil.WriteString(cadena.String())
	if err != nil {
		fmt.Println(err)
		fil.Close()
		return
	}
	fmt.Println(bytes, "bytes escritos satisfactoriamente! :D")
	err = fil.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd := exec.Command(dir + "\\files\\archivo.bat")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
