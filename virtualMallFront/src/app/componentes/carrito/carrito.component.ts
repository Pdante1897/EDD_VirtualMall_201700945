import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Pedido, ProductoPed } from 'src/app/models/pedidos/pedido';
import { Producto } from 'src/app/models/productos/producto';
import { ProductosService } from 'src/app/services/productos/productos.service';

@Component({
  selector: 'app-carrito',
  templateUrl: './carrito.component.html',
  styleUrls: ['./carrito.component.css']
})
export class CarritoComponent implements OnInit {
  constructor(private router: Router, private route: ActivatedRoute, private productoService: ProductosService) { }
  productos : ProductoPed[]=[]
  suma = 0
  date:Date = new Date()
  pedido: Pedido=new Pedido(localStorage.getItem("dpi")||"",this.date.getFullYear(),this.date.getMonth()+1,this.date.getDate(),this.productos,this.suma)

  ngOnInit(): void {
    this.productos = JSON.parse(localStorage.getItem('carrito')||"")
    this.suma=0
    for (let i = 0; i < this.productos.length; i++) {
      this.suma=this.suma+this.productos[i].Precio      
    }

    this.pedido=new Pedido(localStorage.getItem("dpi")||"",this.date.getFullYear(),this.date.getMonth()+1,this.date.getDate(),this.productos,this.suma)
  }

  btnComprar(){
    try {
      this.productos = JSON.parse(localStorage.getItem('carrito')||"")
    var archivo = JSON.stringify(this.pedido)
    console.log(archivo)

    localStorage.removeItem('carrito') 
    this.productoService.postPedido(archivo).subscribe(Response => {null});
    } catch (error) {
      alert("Debe elegir al menos 1 producto para realizar un pedido!")
    }
    
  }

}
