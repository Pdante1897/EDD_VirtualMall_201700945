package Dato

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	MatrizD   *Matriz
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
func (this ListaDPe) Recorrer(anio string) {
	aux := this.Inicio
	for aux != nil {
		GraficarMatriz(aux.MatrizD.Graphviz(), anio+aux.Mes)
		Dir("Matriz" + anio + aux.Mes)
		aux = aux.Siguiente
	}
}

func Dir(num string) {
	dir, err := filepath.Abs(filepath.Dir("./graphviz/graphviz.go"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "cd c:\\program files\\graphviz\\bin\n  ")
	fmt.Fprintf(&cadena, "dot -Tpdf \""+dir+"\\files\\"+num+".dot\" -o \""+dir+"\\files\\grafica"+num+".pdf\"\n  ")
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
func GraficarMatriz(s *strings.Builder, num string) {
	var cadena strings.Builder
	var rank strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"box\" shape=\"record\"]\n")
	fmt.Fprintf(&cadena, "graph[splines=\"ortho\"]\n")
	fmt.Fprintf(&cadena, rank.String())
	fmt.Fprintf(&cadena, s.String())
	fmt.Fprintf(&cadena, "}\n")
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), num, "Matriz")
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
