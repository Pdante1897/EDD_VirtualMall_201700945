package Dato

type NodoHash struct {
	Hash  int
	Valor int
}

type TablaHash struct {
	Size            int
	Carga           int
	Porcentaje      int
	Porcentaje_crec int
	arreglo         []*NodoHash
}

func (this *TablaHash) Posicion(clave int, valor int) int {
	i, p := 0, 0
	p = int(clave % this.Size)
	for this.arreglo[p] != nil && this.arreglo[p].Valor != valor {
		i++
		p = p + 1
		if p >= this.Size {
			p = p - this.Size
		}
	}
	return p
}

func (this *TablaHash) Insertar(nuevo int, valor int) {
	nuevoNodo := NodoHash{nuevo, valor}
	pos := this.Posicion(nuevo, valor)
	this.arreglo[pos] = &nuevoNodo
	this.Carga++
	if ((this.Carga * 100) / this.Size) > this.Porcentaje {
		sizenuevo := this.Size
		for {
			sizenuevo++
			if ((this.Carga * 100) / sizenuevo) <= this.Porcentaje_crec {
				break
			}
		}
		nuevoArray := make([]*NodoHash, sizenuevo)
		viejo := this.arreglo
		this.arreglo = nuevoArray
		this.Size = sizenuevo
		aux := 0
		for i := 0; i < len(viejo); i++ {
			if viejo[i] != nil {
				aux = this.Posicion(viejo[i].Hash, valor)
				nuevoArray[aux] = viejo[i]
			}
		}
	}
}
