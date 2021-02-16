package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./Dato"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}

func agregar(w http.ResponseWriter, r *http.Request) {
	listad := new(Dato.ListaDoble)
	archivo := new(Dato.ArchivoJson)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	for i := 0; i < len(archivo.Datos); i++ {
		fmt.Fprintln(w, archivo.Datos[i].Indice)
		for j := 0; j < len(archivo.Datos[i].Departamentos); j++ {
			fmt.Fprintln(w, archivo.Datos[i].Departamentos[j].Nombre)
			for k := 0; k < len(archivo.Datos[i].Departamentos[j].Tiendas); k++ {
				fmt.Fprintln(w, archivo.Datos[i].Departamentos[j].Tiendas[k].ToString())
				listad.Insertar(archivo.Datos[i].Departamentos[j].Tiendas[k])
			}
		}
	}
	listad.To_String()
}

func main() {
	fmt.Println("un server papu")
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
