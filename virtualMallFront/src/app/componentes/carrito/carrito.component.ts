import { Component, OnInit } from '@angular/core';
import { Producto } from 'src/app/models/productos/producto';

@Component({
  selector: 'app-carrito',
  templateUrl: './carrito.component.html',
  styleUrls: ['./carrito.component.css']
})
export class CarritoComponent implements OnInit {

  constructor() { }
  productos : Producto[]=[]
  suma = 0
  ngOnInit(): void {
    this.productos = JSON.parse(localStorage.getItem('carrito')||"")
    this.suma=0
    for (let i = 0; i < this.productos.length; i++) {
      this.suma=this.suma+this.productos[i].Precio      
    }
  }

  btnComprar(){
    localStorage.removeItem('carrito') 
    this.productos = JSON.parse(localStorage.getItem('carrito')||"")
  }

}
