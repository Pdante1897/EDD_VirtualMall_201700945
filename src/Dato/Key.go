package Dato

type Key struct {
	Value int
	Left  *NodoB
	Right *NodoB
}

func NewKey(valor int) *Key {
	key := Key{valor, nil, nil}
	return &key
}
