package Dato

type ArbolB struct {
	Num  int
	Raiz *NodoB
}

func NewArbol(nivel int) *ArbolB {
	arbol := ArbolB{nivel, nil}
	nodoRaiz := NewNodoB(nivel)
	arbol.Raiz = nodoRaiz
	return &arbol
}

func (this *ArbolB) Insert(nuevaKey *Key) {
	if this.Raiz.Keys[0] == nil {
		this.Raiz.Poner(0, nuevaKey)
	} else if this.Raiz.Keys[0].Left == nil {
		posInsertado := -1
		nodo := this.Raiz
		posInsertado = this.colocarNodo(nodo, nuevaKey)
		if posInsertado != -1 {
			if posInsertado == nodo.Max {
				medio := nodo.Max / 2
				keycentral := nodo.Keys[medio]
				right := NewNodoB(this.Num)
				left := NewNodoB(this.Num)
				indiceLeft := 0
				indiceRight := 0
				for i := 0; i < nodo.Max; i++ {
					if nodo.Keys[i].Value < keycentral.Value {
						left.Poner(indiceLeft, nodo.Keys[i])
						indiceLeft++
						nodo.Poner(i, nil)
					} else if nodo.Keys[i].Value > keycentral.Value {
						right.Poner(indiceRight, nodo.Keys[i])
						indiceRight++
						nodo.Poner(i, nil)
					}
				}
				nodo.Poner(medio, nil)
				this.Raiz = nodo
				this.Raiz.Poner(0, keycentral)
				left.NodoPadre = this.Raiz
				right.NodoPadre = this.Raiz
				keycentral.Left = left
				keycentral.Right = right
			}
		}
	} else if this.Raiz.Keys[0].Left != nil {
		nodo := this.Raiz
		for nodo.Keys[0].Left != nil {
			cont := 0
			for i := 0; i < nodo.Max; i, cont = i+1, cont+1 {
				if nodo.Keys[i] != nil {
					if nodo.Keys[i].Value > nuevaKey.Value {
						nodo = nodo.Keys[i].Left
						break
					}
				} else {
					nodo = nodo.Keys[i-1].Right
					break
				}
			}
			if cont == nodo.Max {
				nodo = nodo.Keys[cont-1].Right
			}
		}
		indiceCol := this.colocarNodo(nodo, nuevaKey)
		if indiceCol == nodo.Max-1 {
			for nodo.NodoPadre != nil {
				indiceMed := nodo.Max
				keycentral := nodo.Keys[indiceMed]
				left := NewNodoB(this.Num)
				right := NewNodoB(this.Num)
				indiceLeft, indiceRight := 0, 0
				for i := 0; i < nodo.Max; i++ {
					if nodo.Keys[i].Value < keycentral.Value {
						left.Poner(indiceLeft, nodo.Keys[i])
						indiceLeft++
						nodo.Poner(i, nil)
					} else if nodo.Keys[i].Value > keycentral.Value {
						right.Poner(indiceRight, nodo.Keys[i])
						indiceRight++
						nodo.Poner(i, nil)
					}
				}
				nodo.Poner(indiceMed, nil)
				keycentral.Left = left
				keycentral.Right = right
				nodo = nodo.NodoPadre
				left.NodoPadre = nodo
				right.NodoPadre = nodo
				for i := 0; i < left.Max; i++ {
					if left.Keys[i] != nil {
						if left.Keys[i].Left != nil {
							left.Keys[i].Left.NodoPadre = left
						}
						if left.Keys[i].Right != nil {
							left.Keys[i].Right.NodoPadre = left
						}
					}
				}
				for i := 0; i < right.Max; i++ {
					if right.Keys[i] != nil {
						if right.Keys[i].Right != nil {
							right.Keys[i].Right.NodoPadre = right
						}
						if right.Keys[i].Left != nil {
							right.Keys[i].Left.NodoPadre = right
						}
					}
				}
				posCol := this.colocarNodo(nodo, keycentral)
				if posCol == nodo.Max-1 {
					if nodo.NodoPadre == nil {
						indiceCentRaiz := nodo.Max / 2
						keycentralRaiz := nodo.Keys[indiceCentRaiz]
						leftRaiz := NewNodoB(this.Num)
						rightRaiz := NewNodoB(this.Num)
						indiceRightRaiz, indiceLeftRaiz := 0, 0
						for i := 0; i < nodo.Max; i++ {
							if nodo.Keys[i].Value < keycentralRaiz.Value {
								leftRaiz.Poner(indiceLeftRaiz, nodo.Keys[i])
								indiceLeftRaiz++
								nodo.Poner(i, nil)
							} else if nodo.Keys[i].Value > keycentralRaiz.Value {
								rightRaiz.Poner(indiceRightRaiz, nodo.Keys[i])
								indiceRightRaiz++
								nodo.Poner(i, nil)
							}
						}
						nodo.Poner(indiceCentRaiz, nil)
						nodo.Poner(0, keycentralRaiz)
						for i := 0; i < this.Num; i++ {
							if leftRaiz.Keys[i] != nil {
								leftRaiz.Keys[i].Left.NodoPadre = leftRaiz
								leftRaiz.Keys[i].Right.NodoPadre = leftRaiz
							}
						}
						for i := 0; i < this.Num; i++ {
							if rightRaiz.Keys[i] != nil {
								rightRaiz.Keys[i].Left.NodoPadre = rightRaiz
								rightRaiz.Keys[i].Right.NodoPadre = rightRaiz
							}
						}
						keycentralRaiz.Left = leftRaiz
						keycentralRaiz.Right = rightRaiz
						leftRaiz.NodoPadre = nodo
						rightRaiz.NodoPadre = nodo
						this.Raiz = nodo
					}
					continue
				} else {
					break
				}
			}
		}
	}
}

func (this *ArbolB) colocarNodo(nodo *NodoB, nuevaKey *Key) int {
	indice := -1
	for i := 0; i < nodo.Max; i++ {
		if nodo.Keys[i] == nil {
			inserted := false
			for j := i - 1; j >= 0; j-- {
				if nodo.Keys[j].Value > nuevaKey.Value {
					nodo.Poner(j+1, nodo.Keys[j])
				} else {
					nodo.Poner(j+1, nuevaKey)
					nodo.Keys[j].Right = nuevaKey.Left
					if (j+2) < this.Num && nodo.Keys[j+2] != nil {
						nodo.Keys[j+2].Left = nuevaKey.Right
					}
					inserted = true
					break
				}
			}
			if !inserted {
				nodo.Poner(0, nuevaKey)
				nodo.Keys[1].Left = nuevaKey.Right
			}
			indice = 1
			break
		}
	}
	return indice
}
