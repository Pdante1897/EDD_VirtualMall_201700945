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

	"github.com/gorilla/handlers"
	"github.com/manucorporat/try"

	"./Dato"
	"./Graphviz"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}

var Listad []Dato.ListaDoble
var ListaS Dato.ListaE
var Arbol = Dato.NewArbol(5)
var Grafo Dato.ListaAdy
var MerkleTiendas Dato.Merkle
var MerkleProductos Dato.Merkle
var MerkleUsuarios Dato.Merkle
var MerklePedidos Dato.Merkle

var Ind int = 0
var Dep int = 0

func agregarPedidos(w http.ResponseWriter, r *http.Request) {
	archivo := new(Dato.JsonPedidos)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	for i := 0; i < len(archivo.Pedidos); i++ {
		var buscar Dato.Busqueda
		var listaDP *Dato.ListaDPe
		var matriz *Dato.Matriz
		fecha := strings.Split(archivo.Pedidos[i].Fecha, "-")
		var dia, err1 = strconv.Atoi(fecha[0])
		var mes, err2 = strconv.Atoi(fecha[1])
		var anio, err3 = strconv.Atoi(fecha[2])
		var mesL = meses(mes)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("error")
		}
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
		buscar.Departamento = archivo.Pedidos[i].Departamento
		buscar.Nombre = archivo.Pedidos[i].Tienda
		buscar.Calificacion = archivo.Pedidos[i].Calificacion
		nodo := Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
		if nodo != nil && Arbol.Raiz.Buscar(archivo.Pedidos[i].Cliente) != nil {
			for j := 0; j < len(archivo.Pedidos[i].Productos); j++ {
				var nodoAI = Dato.BusquedaArbIn(nodo.Tienda.Inventario.Raiz, archivo.Pedidos[i].Productos[j].Codigo)
				if nodoAI != nil {
					var cola *Dato.Cola
					cola = matriz.Buscar(buscar.Departamento + strconv.Itoa(dia))
					if cola != nil {

						cola.Push(nodoAI.Valor)

					} else {
						var nuevo *Dato.NodoPedido
						nuevo = new(Dato.NodoPedido)
						var colaN *Dato.Cola
						colaN = new(Dato.Cola)
						colaN.Nombre = buscar.Departamento + strconv.Itoa(dia)
						colaN.Push(nodoAI.Valor)
						nuevo.Cola = colaN
						nuevo.Departamento = buscar.Departamento
						nuevo.Dia = dia
						nuevo.Cliente = archivo.Pedidos[i].Cliente
						matriz.Add(nuevo)

					}

				}
			}
			matriz.Imprimir()
			matriz.Imprimir2()

		} else {
			fmt.Println("No se encontro el cliente")

			continue
		}

	}

}
func meses(num int) string {
	switch num {
	case 1:
		return "Enero"
	case 2:
		return "Febrero"
	case 3:
		return "Marzo"
	case 4:
		return "Abril"
	case 5:
		return "Mayo"
	case 6:
		return "Junio"
	case 7:
		return "Julio"
	case 8:
		return "Agosto"
	case 9:
		return "Septiembre"
	case 10:
		return "Octubre"
	case 11:
		return "Noviembre"
	case 12:
		return "Diciembre"
	}
	return ""
}

func agregarInv(w http.ResponseWriter, r *http.Request) {
	archivo := new(Dato.JsonInventario)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	for i := 0; i < len(archivo.Inventarios); i++ {
		var buscar Dato.Busqueda
		buscar.Departamento = archivo.Inventarios[i].Departamento
		buscar.Nombre = archivo.Inventarios[i].Tienda
		buscar.Calificacion = archivo.Inventarios[i].Calificacion
		for j := 0; j < len(archivo.Inventarios[i].Productos); j++ {
			try.This(func() {
				var producto = archivo.Inventarios[i].Productos[j]
				var nodo = new(Dato.Nodo)
				nodo = Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre)
				var proBus = Dato.BusquedaArbIn(nodo.Tienda.Inventario.Raiz, producto.Codigo)
				if proBus != nil {
					proBus.Valor.Cantidad += producto.Cantidad
				} else {
					Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre).Tienda.Inventario.Insert(producto)
				}
				fmt.Println(Listad[Dato.RowMajor(Ind, Dep, buscar.NumDep(Listad), buscar.Calificacion-1)].Buscar(buscar.Nombre).Tienda.Inventario.Raiz.GenerarGraphviz())
			}).Catch(func(e try.E) {
				fmt.Println(e)
				fmt.Println("Tienda no encontrada")
			})
		}
	}

}

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

func Dir(num string) {
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
		Dir(strconv.Itoa(i + 1))
	}
}
func generarImgM(w http.ResponseWriter, r *http.Request) {
	ListaS.Recorrer()
}
func generarImgInv(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(Listad); i++ {
		Listad[i].Recorrer()
	}
}
func main() {
	var arbol = Dato.NewMerkel()
	for i := 1; i < 10; i++ {

		arbol.Insertar(i, strconv.Itoa(i))
	}
	Dato.GenerarHashArbol(arbol.Raiz)
	Dato.GuardarArchivo(arbol.Raiz.GenerarGraphvizMerk(), "prueba", "")

	tabla := Dato.NewTablaHash(7, 50, 25)
	tabla.Insertar(100, "Hola mundo")
	tabla.Insertar(100, "Que wen servicio")
	tabla.Insertar(101, "Que wen servicio")

	tabla.Insertar(101, "a cuanto el arroz?")
	tabla.Insertar(1, "me gustan las papas fritas")
	tabla.Insertar(35, "por la orda!")
	tabla.Imprimir()

	fmt.Println("un server papu")

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/cargartienda", agregar).Methods("POST")
	router.HandleFunc("/getArreglo", generarImg).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", BuscarEsp).Methods("GET")
	router.HandleFunc("/Eliminar", Eliminar).Methods("DELETE")
	router.HandleFunc("/id/{numero}", BuscId).Methods("GET")
	router.HandleFunc("/cargarinventario", agregarInv).Methods("POST")
	router.HandleFunc("/cargarpedido", agregarPedidos).Methods("POST")
	router.HandleFunc("/getMatriz", generarImgM).Methods("GET")
	router.HandleFunc("/getImgInv", generarImgInv).Methods("GET")
	router.HandleFunc("/getTiendas", getTiendas).Methods("GET")
	router.HandleFunc("/getUsuarios", getUsuarios).Methods("GET")
	router.HandleFunc("/getUsuariosSha256", getUsuariosSha256).Methods("GET")
	router.HandleFunc("/getUsuariosBcrypt", getUsuariosBcrypt).Methods("GET")
	router.HandleFunc("/cargarUsuarios", cargarUsuarios).Methods("POST")
	router.HandleFunc("/cargarGrafos", cargarGrafos).Methods("POST")
	router.HandleFunc("/getProductos/{nombre}+{departamento}+{calificacion}", getInventarios).Methods("GET")
	router.HandleFunc("/getPDFs", getPDFs).Methods("GET")
	router.HandleFunc("/buscarUser", BuscarUser).Methods("POST")
	router.HandleFunc("/getComentarios/{nombre}+{departamento}+{calificacion}", GetComentarios).Methods("GET")
	router.HandleFunc("/postComentario/{nombre}+{departamento}+{calificacion}", PostComentario).Methods("POST")
	router.HandleFunc("/getComentariosProd/{nombre}+{departamento}+{calificacion}+{producto}", GetComentariosProd).Methods("GET")
	router.HandleFunc("/postComentarioProd/{nombre}+{departamento}+{calificacion}+{producto}", PostComentarioProd).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func cargarGrafos(w http.ResponseWriter, r *http.Request) {
	Grafo = *Dato.NewLista()
	archivo := new(Dato.ArchivoJsonGrafo)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	for i := 0; i < len(archivo.Nodos); i++ {
		Grafo.Insert(archivo.Nodos[i].Nombre)
	}
	for i := 0; i < len(archivo.Nodos); i++ {
		for j := 0; j < len(archivo.Nodos[i].Enlaces); j++ {
			Grafo.Enlazar(archivo.Nodos[i].Nombre, archivo.Nodos[i].Enlaces[j].Nombre, archivo.Nodos[i].Enlaces[j].Distancia)
		}
	}
	Grafo.Inicial = archivo.PosicionInicialrobot
	Grafo.Entrega = archivo.Entrega
	Grafo.Draw()
	//user := Arbol.Raiz.Buscar(7771566947243)

	//fmt.Println(user.Nombre)

}

func cargarUsuarios(w http.ResponseWriter, r *http.Request) {
	archivo := new(Dato.JsonUsuarios)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)
	for i := 0; i < len(archivo.Usuarios); i++ {
		nuevakey := Dato.NewKey(archivo.Usuarios[i].Dpi, archivo.Usuarios[i])

		Arbol.Insert(nuevakey)
		fmt.Println(archivo.Usuarios[i].Nombre)
	}
	fmt.Println(Arbol.Raiz.Keys[0].Value)
}

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	Dato.GraficarArbolB(Arbol.Raiz)
}
func getUsuariosSha256(w http.ResponseWriter, r *http.Request) {
	Dato.GraficarArbolBSha256(Arbol.Raiz)
}
func getUsuariosBcrypt(w http.ResponseWriter, r *http.Request) {
	Dato.GraficarArbolBBcript(Arbol.Raiz)
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
func getTiendas(w http.ResponseWriter, r *http.Request) {
	var tiendas []Dato.TiendaF
	for i := 0; i < len(Listad); i++ {
		var aux = Listad[i].GetTiendas()
		for j := 0; j < len(aux); j++ {
			tiendas = append(tiendas, aux[j])
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Dato.ArrTienda{Tiendas: tiendas})

}
func getPDFs(w http.ResponseWriter, r *http.Request) {
	dir, err := filepath.Abs(filepath.Dir("./graphviz/files/"))
	if err != nil {
		log.Fatal(err)
	}
	archivos, err := ioutil.ReadDir("./Graphviz/files")
	if err != nil {
		log.Fatal(err)
	}
	listapdfs := Graphviz.NewListaPDF()
	var pdfs []Graphviz.Archivopdf
	for _, archivo := range archivos {
		if strings.Contains(archivo.Name(), ".pdf") {
			fmt.Println("Nombre:", dir+"\\"+archivo.Name())
			var pdf Graphviz.Archivopdf
			pdf.Nombre = archivo.Name()
			pdf.Ruta = dir + "\\" + archivo.Name()
			pdfs = append(pdfs, pdf)
			copiar(pdf)
		}

	}
	listapdfs.Pdfs = pdfs
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Graphviz.ListaPdf{Pdfs: listapdfs.Pdfs})

}

func copiar(pdf Graphviz.Archivopdf) {
	dir, err := filepath.Abs(filepath.Dir("./graphviz/graphviz.go"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "copy  "+pdf.Ruta+" C:\\Users\\gerar\\Desktop\\ProyectoEdd\\EDD_VirtualMall_201700945\\virtualMallFront\\src\\assets\\download\\ \n  ")
	fil, err := os.Create(dir + "\\files\\copiar.bat")
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
	cmd := exec.Command(dir + "\\files\\copiar.bat")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func getInventarios(w http.ResponseWriter, r *http.Request) {
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
	productos := encontrado.GetProductos()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Dato.ArrProducto{Productos: productos.Productos})

}
