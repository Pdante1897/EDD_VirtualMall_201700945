package Dato

import (
	"fmt"
	"reflect"
)

type NodoPedido struct {
	Norte        interface{}
	Sur          interface{}
	Este         interface{}
	Oeste        interface{}
	Cola         Producto
	Departamento string
	Dia          int
}

type NodoCabVert struct {
	Norte        interface{}
	Sur          interface{}
	Este         interface{}
	Oeste        interface{}
	Departamento string
}
type NodoCabHor struct {
	Norte interface{}
	Sur   interface{}
	Este  interface{}
	Oeste interface{}
	Dia   int
}

type Matriz struct {
	CabHor *NodoCabHor
	CabVer *NodoCabVert
}

func (this *Matriz) getVertical(dato string) interface{} {
	if this.CabVer == nil {
		return nil
	}
	var aux interface{} = this.CabVer
	for aux != nil {
		if aux.(*NodoCabVert).Departamento == dato {
			return aux
		}
		aux = aux.(*NodoCabVert).Sur
	}
	return nil
}
func (this *Matriz) getHorizontal(dato int) interface{} {
	if this.CabHor == nil {
		return nil
	}
	var aux interface{} = this.CabHor
	for aux != nil {
		if aux.(*NodoCabHor).Dia == dato {
			return aux
		}
		aux = aux.(*NodoCabVert).Este
	}
	return nil
}

func (this *Matriz) crearVert(dato string) *NodoCabVert {
	if this.CabVer == nil {
		nuevo := &NodoCabVert{
			Este:         nil,
			Oeste:        nil,
			Sur:          nil,
			Norte:        nil,
			Departamento: dato,
		}
		this.CabVer = nuevo
		return nuevo
	}
	var aux interface{} = this.CabVer
	if dato < aux.(*NodoCabVert).Departamento {
		nuevo := &NodoCabVert{
			Este:         nil,
			Oeste:        nil,
			Sur:          nil,
			Norte:        nil,
			Departamento: dato,
		}
		nuevo.Sur = this.CabVer
		this.CabVer.Norte = nuevo
		this.CabVer = nuevo
		return nuevo
	}
	for aux.(*NodoCabVert).Sur != nil {
		if dato > aux.(*NodoCabVert).Departamento && dato <= aux.(*NodoCabVert).Sur.(*NodoCabVert).Departamento {
			nuevo := &NodoCabVert{
				Este:         nil,
				Oeste:        nil,
				Sur:          nil,
				Norte:        nil,
				Departamento: dato,
			}
			temp := aux.(*NodoCabVert).Sur
			temp.(*NodoCabVert).Norte = nuevo
			nuevo.Sur = temp
			aux.(*NodoCabVert).Sur = nuevo
			nuevo.Norte = aux
		}
		aux = aux.(*NodoCabVert).Sur
	}
	nuevo := &NodoCabVert{
		Este:         nil,
		Oeste:        nil,
		Sur:          nil,
		Norte:        nil,
		Departamento: dato,
	}
	aux.(*NodoCabVert).Sur = nuevo
	nuevo.Norte = aux
	return nuevo
}

func (this *Matriz) crearHor(dato int) *NodoCabHor {
	if this.CabHor == nil {
		nuevo := &NodoCabHor{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dia:   dato,
		}
		this.CabHor = nuevo
		return nuevo
	}
	var aux interface{} = this.CabHor
	if dato < aux.(*NodoCabHor).Dia {
		nuevo := &NodoCabHor{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dia:   dato,
		}
		nuevo.Este = this.CabHor
		this.CabHor.Oeste = nuevo
		this.CabHor = nuevo
		return nuevo
	}
	for aux.(*NodoCabHor).Este != nil {
		if dato > aux.(*NodoCabHor).Dia && dato <= aux.(*NodoCabHor).Este.(*NodoCabHor).Dia {
			nuevo := &NodoCabHor{
				Este:  nil,
				Oeste: nil,
				Sur:   nil,
				Norte: nil,
				Dia:   dato,
			}
			temp := aux.(*NodoCabHor).Este
			temp.(*NodoCabHor).Oeste = nuevo
			nuevo.Este = temp
			aux.(*NodoCabHor).Este = nuevo
			nuevo.Oeste = aux
		}
		aux = aux.(*NodoCabHor).Este
	}
	nuevo := &NodoCabHor{
		Este:  nil,
		Oeste: nil,
		Sur:   nil,
		Norte: nil,
		Dia:   dato,
	}
	aux.(*NodoCabHor).Este = nuevo
	nuevo.Oeste = aux
	return nuevo
}

func (this *Matriz) getUltV(cab *NodoCabHor, dato string) interface{} {
	if cab.Sur == nil {
		return cab
	}
	aux := cab.Sur
	if dato <= aux.(*NodoPedido).Departamento {
		return cab
	}
	for aux.(*NodoPedido).Sur != nil {
		if dato > aux.(*NodoPedido).Departamento && dato <= aux.(*NodoPedido).Sur.(*NodoPedido).Departamento {
			return aux
		}
		aux = aux.(*NodoPedido).Sur
	}
	if dato <= aux.(*NodoPedido).Departamento {
		return aux.(*NodoPedido).Norte
	}
	return aux
}

func (this *Matriz) getUltH(cab *NodoCabVert, dato int) interface{} {
	if cab.Este == nil {
		return cab
	}
	aux := cab.Este
	if dato <= aux.(*NodoPedido).Dia {
		return cab
	}
	for aux.(*NodoPedido).Este != nil {
		if dato > aux.(*NodoPedido).Dia && dato <= aux.(*NodoPedido).Este.(*NodoPedido).Dia {
			return aux
		}
		aux = aux.(*NodoPedido).Este
	}
	if dato <= aux.(*NodoPedido).Dia {
		return aux.(*NodoPedido).Oeste
	}
	return aux
}

func (this *Matriz) Add(nuevo *NodoPedido) {
	vert := this.getVertical(nuevo.Departamento)
	hor := this.getHorizontal(nuevo.Dia)
	if vert == nil {
		vert = this.crearVert(nuevo.Departamento)
	}
	if hor == nil {
		hor = this.crearHor(nuevo.Dia)
	}
	izq := this.getUltH(vert.(*NodoCabVert), nuevo.Dia)
	sup := this.getUltV(hor.(*NodoCabHor), nuevo.Departamento)
	fmt.Println(reflect.TypeOf(izq).String())
	fmt.Println(reflect.TypeOf(sup).String())

	if reflect.TypeOf(izq).String() == "*Dato.NodoPedido" {
		if izq.(*NodoPedido).Este == nil {
			izq.(*NodoPedido).Este = nuevo
			nuevo.Oeste = izq
		} else {
			temp := izq.(*NodoPedido).Este
			izq.(*NodoPedido).Este = nuevo
			nuevo.Oeste = izq
			temp.(*NodoPedido).Oeste = nuevo
			nuevo.Este = temp
		}
	} else {
		if izq.(*NodoCabVert).Este == nil {
			izq.(*NodoCabVert).Este = nuevo
			nuevo.Oeste = izq
		} else {
			temp := izq.(*NodoCabVert).Este
			izq.(*NodoCabVert).Este = nuevo
			nuevo.Oeste = izq
			temp.(*NodoPedido).Oeste = nuevo
			nuevo.Este = temp
		}
	}

	if reflect.TypeOf(sup).String() == "*Dato.NodoPedido" {
		if sup.(*NodoPedido).Sur == nil {
			sup.(*NodoPedido).Sur = nuevo
			nuevo.Norte = sup
		} else {
			temp := sup.(*NodoPedido).Sur
			sup.(*NodoPedido).Sur = nuevo
			nuevo.Norte = sup
			temp.(*NodoPedido).Norte = nuevo
			nuevo.Sur = temp
		}
	} else {
		if sup.(*NodoCabHor).Sur == nil {
			sup.(*NodoCabHor).Sur = nuevo
			nuevo.Norte = sup
		} else {
			temp := sup.(*NodoCabHor).Sur
			sup.(*NodoCabHor).Sur = nuevo
			nuevo.Norte = sup
			temp.(*NodoPedido).Norte = nuevo
			nuevo.Sur = temp
		}
	}
}
func (this *Matriz) Imprimir() {
	var aux interface{} = this.CabVer
	for aux != nil {
		fmt.Print(aux.(*NodoCabVert).Departamento, "***************")
		tmp := aux.(*NodoCabVert).Este
		for tmp != nil {
			fmt.Printf("%v,%v------", tmp.(*NodoPedido).Cola.Nombre, tmp.(*NodoPedido).Departamento)
			tmp = tmp.(*NodoPedido).Este
		}
		fmt.Print("\n")
		aux = aux.(*NodoCabVert).Sur
	}
}

func (this *Matriz) Imprimir2() {
	var aux interface{} = this.CabHor
	for aux != nil {
		fmt.Print(aux.(*NodoCabHor).Dia, "*****************")
		tmp := aux.(*NodoCabHor).Sur
		for tmp != nil {
			fmt.Printf("%v,%v-------", tmp.(*NodoPedido).Cola.Nombre, tmp.(*NodoPedido).Departamento)
			tmp = tmp.(*NodoPedido).Sur
		}
		fmt.Println("")
		aux = aux.(*NodoCabHor).Este
	}
}
