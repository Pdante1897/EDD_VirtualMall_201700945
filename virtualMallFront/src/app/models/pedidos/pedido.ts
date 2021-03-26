import { Producto } from "../productos/producto";

export class Pedido {
    Productos: Producto[]
    Total: number
    constructor(productos: Producto[], total: number){
        this.Productos=productos
        this.Total=total
    }
}
