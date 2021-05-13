package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

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
	prodEnc.Valor.Comentarios.Insertar(archivo.Id, archivo.Dpi, archivo.Cadena)
}

func PostSubComentarioProd(w http.ResponseWriter, r *http.Request) {
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
	archivo := new(Dato.ArrComent)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	buscar.Calificacion = calificacion
	var encontrado *Dato.Nodo
	var prodEnc *Dato.NodoAI
	var aux1 *Dato.NodoHash
	var tabla *Dato.TablaHash
	var aux Dato.Comentario
	encontrado = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
	prodEnc = Dato.BusquedaArbIn(encontrado.Tienda.Inventario.Raiz, producto)
	tabla = prodEnc.Valor.Comentarios
	for i := 0; i < len(archivo.Comentarios); i++ {

		for j := 0; j < len(tabla.Arreglo); j++ {
			if tabla.Arreglo[j] != nil {
				if tabla.Arreglo[j].Hash == archivo.Comentarios[i].Dpi && tabla.Arreglo[j].Valor == archivo.Comentarios[i].Cadena {
					if i == len(archivo.Comentarios)-2 {
						aux1 = tabla.Arreglo[j]
						break
					}
					tabla = tabla.Arreglo[j].SubComentarios

				}
			}

		}

	}
	aux = archivo.Comentarios[len(archivo.Comentarios)-1]
	if aux1.SubComentarios == nil {
		aux1.SubComentarios = Dato.NewTablaHash(7, 50, 30)

	}
	aux1.SubComentarios.Insertar(aux.Id, aux.Dpi, aux.Cadena)
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
	encontrado.Tienda.Comentarios.Insertar(archivo.Id, archivo.Dpi, archivo.Cadena)

	encontrado.Tienda.Comentarios.Imprimir()
}

func PostSubComentario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var buscar Dato.Busqueda
	buscar.Nombre = strings.ReplaceAll(vars["nombre"], "-", " ")
	buscar.Departamento = strings.ReplaceAll(vars["departamento"], "-", " ")
	calificacion, err := strconv.Atoi(vars["calificacion"])

	if err != nil {
		fmt.Fprintf(w, "Id invalida")
		return
	}
	archivo := new(Dato.ArrComent)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	buscar.Calificacion = calificacion
	var encontrado *Dato.Nodo
	var aux1 *Dato.NodoHash
	var tabla *Dato.TablaHash
	var aux Dato.Comentario
	encontrado = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
	tabla = encontrado.Tienda.Comentarios
	for i := 0; i < len(archivo.Comentarios); i++ {

		for j := 0; j < len(tabla.Arreglo); j++ {
			if tabla.Arreglo[j] != nil {
				if tabla.Arreglo[j].Hash == archivo.Comentarios[i].Dpi && tabla.Arreglo[j].Valor == archivo.Comentarios[i].Cadena {
					if i == len(archivo.Comentarios)-2 {
						aux1 = tabla.Arreglo[j]
						break
					}
					tabla = tabla.Arreglo[j].SubComentarios

				}
			}

		}

	}
	aux = archivo.Comentarios[len(archivo.Comentarios)-1]
	if aux1.SubComentarios == nil {
		aux1.SubComentarios = Dato.NewTablaHash(7, 50, 30)

	}
	aux1.SubComentarios.Insertar(aux.Id, aux.Dpi, aux.Cadena)
}

func GenerarBloque() {
	if MerkleTiendas == *Dato.NewMerkel() && MerklePedidos == *Dato.NewMerkel() && MerkleProductos == *Dato.NewMerkel() && MerkleUsuarios == *Dato.NewMerkel() {
		fmt.Println("No se han realizado cambios en el sistema")
	} else {

		Blockchain.AgregarBlock(Dato.GetCadenas(MerkleTiendas.Raiz) + "?" + Dato.GetCadenas(MerkleProductos.Raiz) + "?" + Dato.GetCadenas(MerkleUsuarios.Raiz) + "?" + Dato.GetCadenas(MerklePedidos.Raiz))
		for _, block := range Blockchain.Blocks {
			fmt.Printf("Indice: %v \n", block.Indice)
			fmt.Printf("Fecha: %s \n", block.Fecha)
			fmt.Printf("Data: %s \n", block.Data)
			fmt.Printf("PrevHash: %x \n", block.PrevHash)
			fmt.Printf("Hash: %x \n", block.Hash)
			fmt.Printf("Nonce: %v \n", block.Nonce)

		}
		if MerkleTiendas != *Dato.NewMerkel() {
			Dato.GenerarHashArbol(MerkleTiendas.Raiz)
			Dato.GuardarArchivo(MerkleTiendas.Raiz.Graph(), strconv.Itoa(len(Blockchain.Blocks)), "ArbolMerkleTiendas")
			Dato.DirSvc("ArbolMerkleTiendas" + strconv.Itoa(len(Blockchain.Blocks)))
		}
		if MerklePedidos != *Dato.NewMerkel() {
			Dato.GenerarHashArbol(MerklePedidos.Raiz)
			Dato.GuardarArchivo(MerklePedidos.Raiz.Graph(), strconv.Itoa(len(Blockchain.Blocks)), "ArbolMerklePedidos")
			Dato.DirSvc("ArbolMerklePedidos" + strconv.Itoa(len(Blockchain.Blocks)))
		}
		if MerkleProductos != *Dato.NewMerkel() {
			Dato.GenerarHashArbol(MerkleProductos.Raiz)
			Dato.GuardarArchivo(MerkleProductos.Raiz.Graph(), strconv.Itoa(len(Blockchain.Blocks)), "ArbolMerkleProductos")
			Dato.DirSvc("ArbolMerkleProductos" + strconv.Itoa(len(Blockchain.Blocks)))
		}
		if MerkleUsuarios != *Dato.NewMerkel() {
			Dato.GenerarHashArbol(MerkleUsuarios.Raiz)
			Dato.GuardarArchivo(MerkleUsuarios.Raiz.Graph(), strconv.Itoa(len(Blockchain.Blocks)), "ArbolUsuarios")
			Dato.DirSvc("ArbolUsuarios" + strconv.Itoa(len(Blockchain.Blocks)))
		}

	}

	MerkleTiendas = *Dato.NewMerkel()
	MerklePedidos = *Dato.NewMerkel()
	MerkleProductos = *Dato.NewMerkel()
	MerkleUsuarios = *Dato.NewMerkel()
	fmt.Println("todsa")
}

func Hilo() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	stop := make(chan bool)
	go func() {
		defer func() { stop <- true }()

		count := 0
		for {
			select {
			case <-Ticker.C:
				count++
				fmt.Println("Tick ", count, time.Now())
				GenerarBloque()

			case <-stop:
				fmt.Println("\nGuardando Blocke")
				return
			}
		}
	}()
	<-c
	Ticker.Stop()
	stop <- true
	<-stop

	fmt.Println("Aplicacion detenida")

}

func PostTiempo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	minutos, err := strconv.Atoi(vars["minutos"])
	fmt.Println(minutos)
	if err != nil {
		fmt.Fprintf(w, "Id invalida")
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, nil)
	T = time.Duration(minutos) * 60 * time.Second
	Ticker = time.NewTicker(T)

}

func addPedido(w http.ResponseWriter, r *http.Request) {
	archivo := new(Dato.JsonPedido)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	var buscar Dato.Busqueda
	var listaDP *Dato.ListaDPe
	var matriz *Dato.Matriz
	var dia = archivo.Dia
	var mes = archivo.Mes
	var anio = archivo.Anio
	var mesL = meses(mes)
	var nodo *Dato.Nodo
	var nodoAnio = ListaS.Buscar(anio)
	if nodoAnio == nil {
		var temp = new(Dato.ListaDPe)
		temp.Indice = anio
		ListaS.Insertar(temp)
		listaDP = ListaS.Buscar(anio).Dato
	} else {
		listaDP = ListaS.Buscar(anio).Dato
	}
	var nodoMes = listaDP.Buscar(mesL)
	if nodoMes == nil {
		listaDP.Insertar(mesL)
		listaDP.Buscar(mesL).MatrizD = new(Dato.Matriz)
		nodoMes = listaDP.Buscar(mesL)
		matriz = nodoMes.MatrizD
	} else {
		nodoMes = listaDP.Buscar(mesL)
		matriz = nodoMes.MatrizD
	}
	var nuevo *Dato.NodoPedido
	nuevo = new(Dato.NodoPedido)
	var colaN *Dato.Cola
	colaN = new(Dato.Cola)
	colaN.Nombre = buscar.Departamento + strconv.Itoa(dia)
	for i := 0; i < len(archivo.Productos); i++ {
		buscar.Departamento = archivo.Productos[i].Departamento
		buscar.Nombre = archivo.Productos[i].Tienda
		buscar.Calificacion, err = strconv.Atoi(archivo.Productos[i].Calificacion)
		nodo = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
		aux := *Dato.BusquedaArbIn(nodo.Tienda.Inventario.Raiz, archivo.Productos[i].Codigo)
		aux.Valor.Cantidad = aux.Valor.Cantidad - 1
		colaN.Push(aux.Valor)
	}
	nuevo.Cola = colaN
	nuevo.Departamento = buscar.Departamento
	nuevo.Dia = dia
	cliente, err := strconv.Atoi(archivo.Cliente)
	nuevo.Cliente = int64(cliente)
	matriz.Add(nuevo)
	MercklePedidos()
}

func MercklePedidos() {
	cont := 1
	vector := ListaS.RecorrerL()
	for i := 0; i < len(vector); i++ {
		anio := strconv.Itoa(vector[i].Dato.Indice)
		vector2 := vector[i].Dato.RecorrerL()
		for j := 0; j < len(vector2); j++ {
			mes := vector2[j].Mes
			dias := vector2[j].MatrizD.RecorrerM()
			pedido := "|" + anio + "?" + mes + "?" + dias
			MerklePedidos.Insertar(cont, pedido)
			cont++
		}
	}
}
