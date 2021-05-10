package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"./Dato"
	"github.com/gorilla/mux"
)

func BuscarUser(w http.ResponseWriter, r *http.Request) {
	buscar := new(Dato.BusquedaUsuario)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al enviar")
		return
	}
	json.Unmarshal(reqBody, &buscar)
	encontrado := Arbol.Raiz.Buscar(buscar.Dpi)
	archivo, err := json.Marshal(&encontrado)
	if err != nil {
		fmt.Fprintf(w, "Error ")
		return
	}
	fmt.Fprintf(w, string(archivo))
}

func GetComentarios(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var buscar Dato.Busqueda
	buscar.Nombre = strings.ReplaceAll(vars["nombre"], "-", " ")
	buscar.Departamento = strings.ReplaceAll(vars["departamento"], "-", " ")
	calificacion, err := strconv.Atoi(vars["calificacion"])
	if err != nil {
		fmt.Fprintf(w, "Id invalida")
		return
	}
	buscar.Calificacion = calificacion
	var encontrado *Dato.Nodo
	encontrado = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
	comentarios := encontrado.GetComentarios()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Dato.ArrComent{Comentarios: comentarios.Comentarios})

}

func GetComentariosProd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var buscar Dato.Busqueda
	buscar.Nombre = strings.ReplaceAll(vars["nombre"], "-", " ")
	buscar.Departamento = strings.ReplaceAll(vars["departamento"], "-", " ")
	calificacion, err := strconv.Atoi(vars["calificacion"])
	producto, err := strconv.Atoi(vars["producto"])

	if err != nil {
		fmt.Fprintf(w, "Id invalida")
		return
	}
	buscar.Calificacion = calificacion
	var encontrado *Dato.Nodo
	var prodEnc *Dato.NodoAI
	encontrado = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
	prodEnc = Dato.BusquedaArbIn(encontrado.Tienda.Inventario.Raiz, producto)
	comentarios := prodEnc.GetComentarios()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Dato.ArrComent{Comentarios: comentarios.Comentarios})

}
func PostComentarioProd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var buscar Dato.Busqueda
	buscar.Nombre = strings.ReplaceAll(vars["nombre"], "-", " ")
	buscar.Departamento = strings.ReplaceAll(vars["departamento"], "-", " ")
	calificacion, err := strconv.Atoi(vars["calificacion"])
	producto, err := strconv.Atoi(vars["producto"])

	if err != nil {
		fmt.Fprintf(w, "Id invalida")
		return
	}
	archivo := new(Dato.Comentario)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	buscar.Calificacion = calificacion
	var encontrado *Dato.Nodo
	var prodEnc *Dato.NodoAI
	encontrado = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
	prodEnc = Dato.BusquedaArbIn(encontrado.Tienda.Inventario.Raiz, producto)
	prodEnc.Valor.Comentarios.Insertar(archivo.Id, archivo.Cadena)
}

func PostComentario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var buscar Dato.Busqueda
	buscar.Nombre = strings.ReplaceAll(vars["nombre"], "-", " ")
	buscar.Departamento = strings.ReplaceAll(vars["departamento"], "-", " ")
	calificacion, err := strconv.Atoi(vars["calificacion"])
	if err != nil {
		fmt.Fprintf(w, "Id invalida")
		return
	}
	archivo := new(Dato.Comentario)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	buscar.Calificacion = calificacion
	var encontrado *Dato.Nodo
	encontrado = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
	encontrado.Tienda.Comentarios.Insertar(archivo.Id, archivo.Cadena)
}
