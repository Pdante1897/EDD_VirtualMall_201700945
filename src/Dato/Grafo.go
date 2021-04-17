package Dato

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ArchivoJsonGrafo struct {
	Nodos []struct {
		Nombre  string `json:"Nombre"`
		Enlaces []struct {
			Nombre    string  `json:"Nombre"`
			Distancia float64 `json:"Distancia"`
		} `json:"Enlaces"`
	} `json:"Nodos"`
	PosicionInicialrobot string `json:"PosicionInicialRobot"`
	Entrega              string `json:"Entrega"`
}

type ListaAdy struct {
	Lista   *ListaDobleVert
	Inicial string
	Entrega string
}

type Vertice struct {
	Nombre     string
	Adyacentes *ListaDobleVert
}

type NodoVert struct {
	Vertice   *Vertice
	Distancia float64
	Anterior  *NodoVert
	Siguiente *NodoVert
}

func NewVertice(nombre string) *Vertice {
	lista := ListaDobleVert{nil, nil}
	return &Vertice{nombre, &lista}
}
func NewListaDobleVert() *ListaDobleVert {
	lista := ListaDobleVert{nil, nil}
	return &lista
}

func NewLista() *ListaAdy {
	lista := ListaDobleVert{nil, nil}
	return &ListaAdy{&lista, "", ""}
}

func (this *ListaAdy) GetVertice(nombre string) *Vertice {
	nodo := this.Lista.Buscar(nombre)
	if nodo != nil {
		return nodo.Vertice

	} else {
		return nil
	}
}

func (this *ListaAdy) Insert(nombre string) {
	if this.GetVertice(nombre) == nil {
		nuevo := NewVertice(nombre)
		this.Lista.Insertar(nuevo, 0)
	} else {
		fmt.Println("Vertice ya agregado")
	}
}

func (this *ListaAdy) Enlazar(a string, b string, distancia float64) {
	var origen *Vertice
	var dest *Vertice
	origen = this.GetVertice(a)
	dest = this.GetVertice(b)

	if origen == nil || dest == nil {
		fmt.Println("no se encontro el vertice")
		return
	}

	origen.Adyacentes.Insertar(dest, distancia)
}

func contiene(busc *ListaDobleVert, element *Vertice) bool {
	for e := busc.Inicio; e != nil; e = e.Siguiente {
		if e.Vertice == element {
			return true
		}
	}
	return false
}

func (this *ListaAdy) Draw() {
	aux := NewListaDobleVert()
	var s strings.Builder
	fmt.Fprintf(&s, "digraph G{\n")
	for e := this.Lista.Inicio; e != nil; e = e.Siguiente {
		temp := e.Vertice
		dist := e.Distancia
		if !contiene(aux, temp) {
			aux.Insertar(temp, dist)
			fmt.Fprintf(&s, "node%p[label=\"%v\"]\n", &(*temp), temp.Nombre)

			if this.Inicial == temp.Nombre {
				fmt.Fprintf(&s, "node%p[style=filled; color=\"blue\"]\n", &(*temp))

			}
		}

		for temporal := temp.Adyacentes.Inicio; temporal != nil; temporal = temporal.Siguiente {
			verTemp := temporal.Vertice
			distemp := temporal.Distancia
			fmt.Fprintf(&s, "node%p->node%p[label=\"%v\"]\n", &(*temp), &(*verTemp), distemp)
			if contiene(aux, verTemp) == false {
				aux.Insertar(verTemp, dist)
				fmt.Fprintf(&s, "node%p[label=\"%v\"]\n", &(*verTemp), verTemp.Nombre)

				if this.Entrega == verTemp.Nombre {
					fmt.Fprintf(&s, "node%p[style=filled; color=\"green\"]\n", &(*temp))

				}
			}
		}
	}

	fmt.Fprintf(&s, "}")
	guardarArchivoGrafo(s.String(), "", "Grafo")
	DirGrafo("Grafo")

}

func guardarArchivoGrafo(cadena string, num string, nom string) {
	fil, err := os.Create("./Graphviz/files/" + nom + num + ".circo")
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

func DirGrafo(num string) {
	dir, err := filepath.Abs(filepath.Dir("./graphviz/graphviz.go"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "cd c:\\program files\\graphviz\\bin\n  ")
	fmt.Fprintf(&cadena, "circo -Tpdf \""+dir+"\\files\\"+num+".circo\" -o \""+dir+"\\files\\grafica"+num+".pdf\"\n  ")
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

type ListaDobleVert struct {
	Inicio *NodoVert
	Fin    *NodoVert
}

func (this *ListaDobleVert) Vacio() bool {
	if this.Inicio == nil {
		return true
	} else {
		return false
	}

}
func (this *ListaDobleVert) Insertar(dato *Vertice, dist float64) {
	aux := &NodoVert{
		Vertice:   dato,
		Distancia: dist,
		Siguiente: nil,
		Anterior:  nil,
	}

	if this.Vacio() {
		this.Inicio = aux
		this.Fin = this.Inicio
	} else if this.Inicio.Siguiente == nil {
		aux.Anterior = this.Inicio
		this.Fin = aux
		this.Inicio.Siguiente = this.Fin
	} else {
		aux.Anterior = this.Fin
		this.Fin.Siguiente = aux
		this.Fin = this.Fin.Siguiente

	}
}

func (this *ListaDobleVert) Buscar(cadena string) *NodoVert {
	var auxiliar *NodoVert
	auxiliar = this.Inicio
	for auxiliar != nil {
		if auxiliar.Vertice.Nombre == cadena {
			return auxiliar
		}
		auxiliar = auxiliar.Siguiente

	}
	return nil
}
