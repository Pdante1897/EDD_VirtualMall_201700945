package Dato

type NodoCabVert struct {
	Norte interface{}
	Sur   interface{}
	Este  interface{}
	Oeste interface{}
	Dato  int
}
type NodoCabHor struct {
	Norte interface{}
	Sur   interface{}
	Este  interface{}
	Oeste interface{}
	letra string
}

type Matriz struct {
	CabHor *NodoCabHor
	CabVer *NodoCabVert
}

func (this *Matriz) getVertical(dato int) interface{} {
	if this.CabVer == nil {
		return nil
	}
	var aux interface{} = this.CabVer
	for aux != nil {
		if aux.(*NodoCabVert).Dato == dato {
			return aux
		}
		aux = aux.(*NodoCabVert).Sur
	}
	return nil
}
func (this *Matriz) getHorizontal(nom string) interface{} {
	if this.CabHor == nil {
		return nil
	}
	var aux interface{} = this.CabHor
	for aux != nil {
		if aux.(*NodoCabHor).letra == string(nom[0]) {
			return aux
		}
		aux = aux.(*NodoCabVert).Este
	}
	return nil
}

func (this *Matriz) crearVert(dato int) *NodoCabVert {
	if this.CabVer == nil {
		nuevo := &NodoCabVert{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dato:  dato,
		}
		this.CabVer = nuevo
		return nuevo
	}
	var aux interface{} = this.CabVer
	if dato < aux.(*NodoCabVert).Dato {
		nuevo := &NodoCabVert{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dato:  dato,
		}
		nuevo.Sur = this.CabVer
		this.CabVer.Norte = nuevo
		this.CabVer = nuevo
		return nuevo
	}
	for aux.(*NodoCabVert).Sur != nil {
		if dato > aux.(*NodoCabVert).Dato && dato < aux.(*NodoCabVert).Sur.(*NodoCabVert).Dato {
			nuevo := &NodoCabVert{
				Este:  nil,
				Oeste: nil,
				Sur:   nil,
				Norte: nil,
				Dato:  dato,
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
		Este:  nil,
		Oeste: nil,
		Sur:   nil,
		Norte: nil,
		Dato:  dato,
	}
	aux.(*NodoCabVert).Sur = nuevo
	nuevo.Norte = aux
	return nuevo
}

func (this *Matriz) crearHor(nom string) *NodoCabHor {
	if this.CabHor == nil {
		nuevo := &NodoCabHor{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			letra: string(nom[0]),
		}
		this.CabHor = nuevo
		return nuevo
	}
	var aux interface{} = this.CabHor
	if string(nom[0]) < aux.(*NodoCabHor).letra {
		nuevo := &NodoCabHor{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			letra: string(nom[0]),
		}
		nuevo.Sur = this.CabHor
		this.CabHor.Norte = nuevo
		this.CabHor = nuevo
		return nuevo
	}
	for aux.(*NodoCabHor).Sur != nil {
		if string(nom[0]) > aux.(*NodoCabHor).letra && string(nom[0]) < aux.(*NodoCabHor).Sur.(*NodoCabHor).letra {
			nuevo := &NodoCabHor{
				Este:  nil,
				Oeste: nil,
				Sur:   nil,
				Norte: nil,
				letra: string(nom[0]),
			}
			temp := aux.(*NodoCabHor).Sur
			temp.(*NodoCabHor).Norte = nuevo
			nuevo.Sur = temp
			aux.(*NodoCabHor).Sur = nuevo
			nuevo.Norte = aux
		}
		aux = aux.(*NodoCabHor).Sur
	}
	nuevo := &NodoCabHor{
		Este:  nil,
		Oeste: nil,
		Sur:   nil,
		Norte: nil,
		letra: string(nom[0]),
	}
	aux.(*NodoCabHor).Sur = nuevo
	nuevo.Norte = aux
	return nuevo
}
