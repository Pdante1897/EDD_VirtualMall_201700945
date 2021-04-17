package Dato

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manucorporat/try"
)

type NodoB struct {
	Max       int
	NodoPadre *NodoB
	Keys      []*Key
}

func NewNodoB(max int) *NodoB {
	keys := make([]*Key, max)
	nodo := NodoB{max, nil, keys}
	return &nodo
}

func (this *NodoB) Poner(pos int, key *Key) {
	this.Keys[pos] = key
}

func (this *NodoB) GenerarGraphviz(s *strings.Builder) string {
	var cadena = ""
	var llaves strings.Builder

	for i := 0; i < 5; i++ {
		try.This(func() {

			if this.Keys[i] != nil {
				fmt.Fprintf(&llaves, "<f"+strconv.FormatInt(this.Keys[i].Value, 10)+">|"+"{Dpi: "+strconv.FormatInt(this.Keys[i].Value, 10)+"| Nombre: "+this.Keys[i].Usuario.Nombre+"}|")

			}
		}).Catch(func(e try.E) {

		})

		try.This(func() {

			if &this.Keys[i].Left != nil {
				fmt.Println("left")
				cadena += this.Keys[i].Left.GenerarGraphviz(s) + ";\n"

			}

		}).Catch(func(e try.E) {

		})

		try.This(func() {

			if &this.Keys[i].Right != nil {
				fmt.Println("right")
				cadena += this.Keys[i].Right.GenerarGraphviz(s) + ";\n"

			}

		}).Catch(func(e try.E) {

		})
	}

	fmt.Fprintf(s, "node"+strconv.FormatInt(this.Keys[0].Value, 10)+"[label=\""+llaves.String()+"\"];\n")
	if this.NodoPadre != nil {
		fmt.Println("padre")

		for i := 0; i < this.Max; i++ {
			if this.NodoPadre.Keys[i].Left == this || this.NodoPadre.Keys[i].Right == this {
				fmt.Fprintf(s, "node"+strconv.FormatInt(this.NodoPadre.Keys[0].Value, 10)+": f"+strconv.FormatInt(this.NodoPadre.Keys[i].Value, 10)+"->node"+strconv.FormatInt(this.Keys[0].Value, 10)+";\n")
			}
		}
	}
	cadena = s.String()
	return cadena
}

func GraficarArbolB(Nodo *NodoB) {
	var cadena strings.Builder
	var cadena2 strings.Builder

	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"box\" shape=\"record\"]\n")
	fmt.Fprintf(&cadena, Nodo.GenerarGraphviz(&cadena2))
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "", "ArbolUser")
	Dir("ArbolUser")
}

func (this *NodoB) Buscar(key int64) *Usuario {
	for i := 0; i < len(this.Keys); i++ {
		//try.This(func() {
		if this.Keys[i] != nil {
			if this.Keys[i].Value == key {
				user := this.Keys[i].Usuario
				return &user
			}
			if key < this.Keys[i].Value {
				return this.Keys[i].Left.Buscar(key)

			} else if i == 4 && this.Keys[i] != nil {
				return this.Keys[i].Right.Buscar(key)
			} else if key > this.Keys[i].Value {
				a := false
				try.This(func() {
					if this.Keys[i+1] == nil {
						a = true
					} else {
						a = false
					}
				}).Catch(func(e try.E) {
					a = true
				})
				if a {
					return this.Keys[i].Right.Buscar(key)
				} else {
					continue
				}

			}
		}

		//}).Catch(func(e try.E) {

		//})
	}

	return nil
}
