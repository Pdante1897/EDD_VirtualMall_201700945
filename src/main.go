package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"./Dato"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}

var Listad []Dato.ListaDoble

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
	var contador int = 0
	var calificacion int = 1
	var lista int = 0
	var indice int = 0
	for i := 0; i < tamanio; i++ {
		Listad[i].Calificacion = calificacion
		calificacion++
		if calificacion > 5 {
			calificacion = 1
		}
		if i == 0 {
			Listad[i].Nombre = archivo.Datos[0].Departamentos[lista].Nombre

		} else if i%5 == 0 {
			indice++
			if indice == len(archivo.Datos) {
				indice = 0
				lista++

			}
			Listad[i].Nombre = archivo.Datos[0].Departamentos[lista].Nombre
		} else {
			Listad[i].Nombre = archivo.Datos[0].Departamentos[lista].Nombre

		}
		Listad[i].Indice = archivo.Datos[indice].Indice

	}
	for i := 0; i < len(archivo.Datos); i++ {
		fmt.Fprintln(w, archivo.Datos[i].Indice)
		for j := 0; j < len(archivo.Datos[i].Departamentos); j++ {
			fmt.Fprintln(w, archivo.Datos[i].Departamentos[j].Nombre)
			for k := 0; k < len(archivo.Datos[i].Departamentos[j].Tiendas); k++ {

				//fmt.Fprintln(w, archivo.Datos[i].Departamentos[j].Tiendas[k].ToString())
				//Listad[contador].Insertar(archivo.Datos[i].Departamentos[j].Tiendas[k])

			}
			contador++

		}

	}
	for i := 0; i < tamanio; i++ {
		Listad[i].To_String()

	}

}

func main() {
	fmt.Println("un server papu")
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/cargartienda", agregar).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
