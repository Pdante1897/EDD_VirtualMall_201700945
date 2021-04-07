package Dato

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
