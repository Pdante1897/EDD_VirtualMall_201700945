package Graphviz

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"../Dato"
)

func Graficar(listad []Dato.ListaDoble, tamanio int) {
	var cadena strings.Builder
	var rank strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"box\" shape=\"record\"]\n")
	fmt.Fprintf(&cadena, "graph[splines=\"ortho\"]\n")
	fmt.Fprintf(&rank, "{rank=\"same\"")
	for i := tamanio * 5; i < (tamanio+1)*5; i++ {
		fmt.Fprintf(&cadena, "node"+strconv.Itoa(i)+"[label=\"%v|%v|%v|pos: %v\"];\n", listad[i].Indice, listad[i].Nombre, listad[i].Calificacion, i+1)
		archivo(listad[i].Inicio, &cadena, nil, strconv.Itoa(i))
		if i+1 < (tamanio+1)*5 {
			fmt.Fprintf(&cadena, "node"+strconv.Itoa(i)+"->node"+strconv.Itoa(i+1)+";\n")
		}
		fmt.Fprintf(&rank, "; node"+strconv.Itoa(i))

	}
	fmt.Fprintf(&rank, "}\n")
	fmt.Fprintf(&cadena, rank.String())
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), strconv.Itoa(tamanio+1), "lista")
}
func archivo(anterior *Dato.Nodo, s *strings.Builder, actual *Dato.Nodo, i string) {
	if anterior != nil {
		fmt.Fprintf(s, "node%p[label=\"%v|%v\"];\n", &(*anterior), anterior.Tienda.Nombre, anterior.Tienda.Contacto)
		if i != "" {
			fmt.Fprintf(s, "node"+i+"->node%p;\n", &(*anterior))

		}
		if actual != nil {
			fmt.Fprintf(s, "node%p->node%p;\n", &(*actual), &(*anterior))
			fmt.Fprintf(s, "node%p->node%p;\n", &(*anterior), &(*actual))
		}
		archivo(anterior.Siguiente, s, anterior, "")
	}
}
func guardarArchivo(cadena string, num string, nom string) {
	fil, err := os.Create("./Graphviz/files/" + nom + num + ".dot")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error")
		return
	}
	bytes, err := fil.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error")
		fil.Close()
		return
	}
	fmt.Println(bytes, "bytes escritos satisfactoriamente! :D")
	err = fil.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Println("error")
		return
	}
}

func GraficarArbol(s string, num string) {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"box\" shape=\"record\"]\n")
	fmt.Fprintf(&cadena, "graph[splines=\"ortho\"]\n")
	fmt.Fprintf(&cadena, s)
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), num, "Arbol")
}
