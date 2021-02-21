package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"./Dato"
	"./Graphviz"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}

var Listad []Dato.ListaDoble
var Ind int = 0
var Dep int = 0

func agregar(w http.ResponseWriter, r *http.Request) {
	archivo := new(Dato.ArchivoJson)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	var tamanio int = 5 * len(archivo.Datos) * len(archivo.Datos[0].Departamentos)
	Listad = make([]Dato.ListaDoble, tamanio)
	fmt.Println(strconv.Itoa(tamanio))
	//var contador int = 0
	var calificacion int = 1
	var lista int = 0
	var indice int = 0
	Ind = len(archivo.Datos)
	Dep = len(archivo.Datos[0].Departamentos)
	for q := 0; q < tamanio; q++ {
		Listad[q].Calificacion = calificacion
		calificacion++
		if calificacion > 5 {
			calificacion = 1
		}
		if q == 0 {
			Listad[q].Nombre = archivo.Datos[0].Departamentos[lista].Nombre
		} else if q%5 == 0 {
			lista++
			if lista == len(archivo.Datos[0].Departamentos) {
				lista = 0
				indice++
			}
			Listad[q].Nombre = archivo.Datos[0].Departamentos[lista].Nombre
		} else {
			Listad[q].Nombre = archivo.Datos[0].Departamentos[lista].Nombre
		}
		Listad[q].Indice = archivo.Datos[indice].Indice
		for i := 0; i < len(archivo.Datos); i++ {
			//fmt.Fprintln(w, archivo.Datos[i].Indice)
			for j := 0; j < len(archivo.Datos[i].Departamentos); j++ {
				//fmt.Fprintln(w, archivo.Datos[i].Departamentos[j].Nombre)
				for k := 0; k < len(archivo.Datos[i].Departamentos[j].Tiendas); k++ {
					if Listad[q].Indice == archivo.Datos[i].Indice && Listad[q].Nombre == archivo.Datos[i].Departamentos[j].Nombre && Listad[q].Calificacion == archivo.Datos[i].Departamentos[j].Tiendas[k].Calificacion {
						//fmt.Fprintln(w, archivo.Datos[i].Departamentos[j].Tiendas[k].ToString())
						Listad[q].Insertar(archivo.Datos[i].Departamentos[j].Tiendas[k])
					}
				}
			}
		}
	}
	/*for i := 0; i < tamanio; i++ {
		Listad[i].To_String()

	}*/

}
func BuscarEsp(w http.ResponseWriter, r *http.Request) {
	buscar := new(Dato.Busqueda)
	encontrado := new(Dato.Nodo)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al enviar")
		return
	}
	json.Unmarshal(reqBody, &buscar)
	encontrado = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
	archivo, err := json.Marshal(&encontrado.Tienda)
	if err != nil {
		fmt.Fprintf(w, "Error ")
		return
	}
	fmt.Fprintf(w, string(archivo))
}

func dir(num string) {
	dir, err := filepath.Abs(filepath.Dir("./graphviz/graphviz.go"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "cd c:\\program files\\graphviz\\bin\n  ")
	fmt.Fprintf(&cadena, "dot -Tpdf \""+dir+"\\files\\lista"+num+".dot\" -o \""+dir+"\\files\\grafica"+num+".pdf\"\n  ")
	fil, err := os.Create(dir + "\\files\\archivo.cmd")
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
	cmd := exec.Command(dir + "\\files\\archivo.cmd")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
func generarImg(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(Listad)/5; i++ {
		Graphviz.Graficar(Listad, i)
		dir(strconv.Itoa(i + 1))
	}
}
func main() {
	//Listad[1].Ordenar("Hola mundo")
	fmt.Println("un server papu")
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/cargartienda", agregar).Methods("POST")
	router.HandleFunc("/getArreglo", generarImg).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", BuscarEsp).Methods("POST")
	router.HandleFunc("/Eliminar", Eliminar).Methods("DELETE")
	router.HandleFunc("/id/{numero}", BuscId).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
func Eliminar(w http.ResponseWriter, r *http.Request) {
	eliminar := new(Dato.Eliminar)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al enviar")
		return
	}
	json.Unmarshal(reqBody, &eliminar)
	Listad[Dato.RowMajor(Ind, Dep, eliminar.NumDep(Listad), eliminar.Calificacion-1)].Eliminar(eliminar.Nombre)
	fmt.Fprintf(w, "eliminado")
}
func BuscId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ArchivoTiendas []byte
	lista, err := strconv.Atoi(vars["numero"])
	if err != nil {
		fmt.Fprintf(w, "Id invalida")
		return
	}
	ArchivoTiendas, err = json.Marshal(Listad[lista].BusquedaId())
	if err != nil {
		fmt.Fprintf(w, "Error ")
		return
	}

	fmt.Fprintf(w, string(ArchivoTiendas))

}
func Guardar() {

}
