export class Producto {
    Nombre: string
    Codigo: number
    Descripcion: string
    Precio: number
    Cantidad: number
    Imagen: string

    constructor(nombre: string, codigo: number, descripcion: string, precio: number, calificacion: number, foto: string){
        this.Nombre=nombre
        this.Codigo=codigo
        this.Precio=precio
        this.Descripcion=descripcion
        this.Cantidad=calificacion
        this.Imagen=foto
    }
}

export class ArrProducto{
    Productos: Producto[]=[]
    constructor(){

    }
}