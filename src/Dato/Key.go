package Dato

type Key struct {
	Value   int64
	Usuario Usuario
	Left    *NodoB
	Right   *NodoB
}

func NewKey(valor int64, us Usuario) *Key {
	key := Key{valor, us, nil, nil}
	return &key
}
