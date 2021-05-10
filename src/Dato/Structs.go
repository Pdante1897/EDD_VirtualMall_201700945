package Dato

import (
	"fmt"
	"strconv"
	"strings"
)

type ArchivoJson struct {
	Datos []struct {
		Indice        string `json:"Indice"`
		Departamentos []struct {
			Nombre  string   `json:"Nombre"`
			Tiendas []Tienda `json:"Tiendas"`
		} `json:"Departamentos"`
	} `json:"Datos"`
}
type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Logo         string `json:"Logo"`
	Inventario   ArbolIn
	Comentarios  *TablaHash
}

type TiendaF struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Logo         string `json:"Logo"`
}
type ArrTienda struct {
	Tiendas []TiendaF `json:"Tiendas"`
}

type ArrComent struct {
	Comentarios []Comentario `json: "Comentarios"`
}

type Comentario struct {
	Id     int    `json: "Id"`
	Cadena string `json: "Cadena"`
}

func (this Tienda) ToString() string {
	return ("Nombre: " + this.Nombre + " Descripcion: " + this.Descripcion + " Contacto: " + this.Contacto + " Calificacion: " + strconv.Itoa(this.Calificacion))
}

func Hola() {
	fmt.Println("hola mundo")

}

type Nodo struct {
	Tienda    Tienda
	Anterior  *Nodo
	Siguiente *Nodo
}
type ListaDoble struct {
	Indice       string
	Nombre       string
	Calificacion int
	Inicio       *Nodo
	Fin          *Nodo
}

func (this *ListaDoble) Vacio() bool {
	if this.Inicio == nil {
		return true
	} else {
		return false
	}

}
func (this *ListaDoble) Insertar(dato Tienda) {
	aux := Nodo{
		Tienda:    dato,
		Siguiente: nil,
		Anterior:  nil,
	}
	aux.Tienda.Comentarios = NewTablaHash(7, 50, 25)
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
	//this.ordenar()
}
func ascii(nodo *Nodo) int {
	var numero int
	for i := 0; i < len(nodo.Tienda.Nombre); i++ {
		aux := int(nodo.Tienda.Nombre[i])
		numero += aux
	}
	return numero
}

func (this *ListaDoble) ordenar() {
	var auxiliar = this.Inicio.Siguiente
	var actual = this.Inicio
	if this.Inicio.Siguiente == nil {
		return
	} else {
		if ascii(this.Inicio) > ascii(this.Inicio.Siguiente) {
			auxiliar = this.Inicio
			this.Inicio = this.Inicio.Siguiente
			this.Inicio.Siguiente = auxiliar

		} else {
			for actual.Siguiente != nil {
				if ascii(actual) > ascii(actual.Siguiente) {
					auxiliar = actual.Siguiente
					auxiliar.Siguiente = actual
					this.Fin = actual.Anterior
					this.Fin.Siguiente = auxiliar
					this.Fin = this.Fin.Siguiente
				} else {

				}
				actual = actual.Siguiente
			}

		}
		actual = this.Fin
		auxiliar = actual.Anterior

		if ascii(this.Fin) < ascii(this.Fin.Anterior) {
			auxiliar = this.Fin
			this.Fin = this.Fin.Anterior
			this.Fin.Anterior = auxiliar
		} else {
			for actual.Anterior != nil {
				if ascii(actual) < ascii(actual.Anterior) {
					auxiliar = actual.Anterior
					auxiliar.Anterior = actual
					this.Inicio = actual.Siguiente
					this.Inicio.Anterior = auxiliar
					this.Inicio = this.Fin.Anterior
				} else {
				}
				fmt.Println("aqui")
				actual = actual.Anterior

			}

		}
	}
}

func (this *ListaDoble) Buscar(cadena string) *Nodo {
	var auxiliar *Nodo
	auxiliar = this.Inicio
	for auxiliar != nil {
		if auxiliar.Tienda.Nombre == cadena {
			return auxiliar
		}
		auxiliar = auxiliar.Siguiente

	}
	fmt.Println("no se pudo encontrar")
	return nil
}
func (this *ListaDoble) Eliminar(dato string) {
	var auxiliar1, actual, siguiente *Nodo
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

func (this ListaDoble) To_String() {
	var aux *Nodo
	aux = this.Inicio
	fmt.Println("__________________________________________________________________________________________________")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println(this.Indice + "\n" + this.Nombre + "\n" + strconv.Itoa(this.Calificacion))
	fmt.Println("-----------------------------------------------------------------------")

	for {
		if this.Vacio() {
			break
		}
		if aux == this.Inicio {
			fmt.Println(aux.Tienda.ToString())
			aux = aux.Siguiente
		} else if aux != nil {
			fmt.Println(aux.Tienda.ToString())
			aux = aux.Siguiente
		} else if aux == this.Fin {
			fmt.Println(aux.Tienda.ToString())
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

func (this *ListaDoble) BusquedaId() []Tienda {
	var aux *Nodo
	aux = this.Inicio
	var tiendas []Tienda
	var i int = 0
	for {
		if this.Vacio() {
			tiendas = nil
			break
		}
		if aux == this.Inicio {
			fmt.Println(aux.Tienda.ToString())
			tiendas = append(tiendas, aux.Tienda)
			aux = aux.Siguiente
		} else if aux != nil {
			fmt.Println(aux.Tienda.ToString())
			tiendas = append(tiendas, aux.Tienda)
			aux = aux.Siguiente
		} else if aux == this.Fin {
			tiendas = append(tiendas, aux.Tienda)
			fmt.Println(aux.Tienda.ToString())
			break
		} else {
			break
		}
		i++
	}

	return tiendas
}

func (this *ListaDoble) Recorrer() {
	aux := this.Inicio

	for aux != nil {
		GraficarArbol(aux.Tienda.Inventario.Raiz.GenerarGraphviz(), aux.Tienda.Nombre+strconv.Itoa(aux.Tienda.Calificacion))
		nombre := strings.Replace("Arbol"+aux.Tienda.Nombre+strconv.Itoa(aux.Tienda.Calificacion), " ", "", -1)
		Dir(strings.Replace(nombre, ",", "", -1))
		aux = aux.Siguiente
	}
}

func GraficarArbol(s string, num string) {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"box\" shape=\"record\"]\n")
	fmt.Fprintf(&cadena, s)
	fmt.Fprintf(&cadena, "}\n")
	nombre := strings.Replace(num, " ", "", -1)
	GuardarArchivo(cadena.String(), strings.Replace(nombre, ",", "", -1), "Arbol")
}

func (this *ListaDoble) GetTiendas() []TiendaF {
	var tiendas []TiendaF
	aux := this.Inicio

	for aux != nil {
		aux2 := TiendaF{this.Nombre,
			aux.Tienda.Nombre,
			aux.Tienda.Descripcion,
			aux.Tienda.Contacto,
			aux.Tienda.Calificacion,
			aux.Tienda.Logo,
		}
		tiendas = append(tiendas, aux2)
		aux = aux.Siguiente
	}
	return tiendas
}

func (this *Nodo) GetProductos() ArrProducto {
	var arreglo ArrProducto
	arreglo.Productos = this.Tienda.Inventario.Raiz.ObtenerProductos(nil)
	return arreglo
}

func (this *Nodo) GetComentarios() ArrComent {
	var arreglo ArrComent
	arreglo.Comentarios = this.Tienda.Comentarios.ListaComent()
	return arreglo
}
