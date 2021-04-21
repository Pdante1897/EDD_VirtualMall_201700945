import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ArrProducto, Producto } from 'src/app/models/productos/producto';
import { ProductosService } from 'src/app/services/productos/productos.service';

@Component({
  selector: 'app-productos',
  templateUrl: './productos.component.html',
  styleUrls: ['./productos.component.css']
})
export class ProductosComponent implements OnInit {
  lista_productos: ArrProducto = new ArrProducto()
  listaAux: Producto[]=[]

  constructor(private router: Router, private route: ActivatedRoute, private productoService: ProductosService) {

   }

  ngOnInit(): void {
    var nombre = (this.route.snapshot.paramMap.get('nombre') || '')
    var departamento = (this.route.snapshot.paramMap.get('departamento')|| '')
    var calificacion = (this.route.snapshot.paramMap.get('calificacion')|| '')
    this.productoService.getProductos(nombre, departamento, calificacion).subscribe((dataList: any)=>{
      this.lista_productos = dataList 
      console.log(this.lista_productos)
      console.log(Object.getOwnPropertyDescriptors(this.lista_productos))
      console.log(Object.is(this.lista_productos,ArrProducto))
      console.log(dataList)
      this.setearLista()
      console.log(this.listaAux)
    },(err)=>{
      console.log("no se pudo")

    })
  }
  setearLista(){
    

    for (let i = 0; i < this.lista_productos.Productos.length; i++) {
      var producto: Producto      
      producto = new Producto(this.lista_productos.Productos[i].Nombre,this.lista_productos.Productos[i].Codigo, this.lista_productos.Productos[i].Descripcion, this.lista_productos.Productos[i].Precio, this.lista_productos.Productos[i].Cantidad, this.lista_productos.Productos[i].Imagen)
      this.listaAux.push(producto);
      console.log(i)
    }
  }

  agregarCarrito(prod: Producto){
    var productos : Producto[]=[]
    if (localStorage.getItem('carrito')!=null){
      productos = JSON.parse(localStorage.getItem('carrito')||"")
      productos.push(prod)
      localStorage.setItem('carrito', JSON.stringify(productos))
    }else{
      productos.push(prod)
      localStorage.setItem('carrito', JSON.stringify(productos))

    }

   
  }

}

