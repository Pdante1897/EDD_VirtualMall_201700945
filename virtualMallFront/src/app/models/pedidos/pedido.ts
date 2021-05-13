import { Producto } from "../productos/producto";

export class Pedido {
    Cliente: string
    Anio:number
    Mes:number
    Dia:number
    Productos: ProductoPed[]
    Total: number
    constructor(cliente: string, anio:number, mes:number, dia:number, productos: ProductoPed[], total: number){
        this.Cliente = cliente
        this.Anio=anio
        this.Mes=mes
        this.Dia=dia
        this.Productos=productos
        this.Total=total
    }
}
export class ProductoPed{
    Tienda: string
    Departamento: string
    Calificacion: string
    Nombre:string
    Codigo: number
    Descripcion:string
    Precio: number
    Cantidad: number
    Imagen:string

    constructor(tienda: string, departamento: string, calificiacion:string, codigo: number, precio: number, nombre:string, desc:string, img:string){
        this.Tienda=tienda
        this.Departamento=departamento
        this.Calificacion=calificiacion
        this.Codigo=codigo
        this.Cantidad=1
        this.Precio= precio
        this.Nombre=nombre
        this.Descripcion=desc
        this.Imagen=img
    }
}
