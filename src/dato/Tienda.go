package dato

import "fmt"

func (this *Tiendas) setNombre(nom string) {
	this.nombre = nom
}
func (this Tiendas) getNombre() string {
	return this.nombre
}

func hola() {
	fmt.Println("hola mundo")

}
