import { Component, OnInit } from '@angular/core';
import { ArrTienda, Tienda } from 'src/app/models/tiendas/tienda';
import { TiendasService } from 'src/app/services/tiendas/tiendas.service';
import { FormControl } from '@angular/forms';


@Component({
  selector: 'app-tiendas',
  templateUrl: './tiendas.component.html',
  styleUrls: ['./tiendas.component.css']
})
export class TiendasComponent implements OnInit {
  lista_tiendas: ArrTienda = new ArrTienda()
  listaAux: Tienda[] = []
  constructor(private tiendaservise:TiendasService) {
    
   }

  ngOnInit(): void {
    this.tiendaservise.getTiendas().subscribe((dataList: any)=>{
      this.lista_tiendas = dataList 
      console.log(this.lista_tiendas)
      console.log(Object.getOwnPropertyDescriptors(this.lista_tiendas))
    console.log(Object.is(this.lista_tiendas,ArrTienda))
      console.log(dataList)
      this.setearLista()
      console.log(this.listaAux)
    },(err)=>{
      console.log("no se pudo")

    })
    
  }
  
  setearLista(){
    

    for (let i = 0; i < this.lista_tiendas.Tiendas.length; i++) {
      var tienda: Tienda
      
      tienda = new Tienda(this.lista_tiendas.Tiendas[i].Departamento,this.lista_tiendas.Tiendas[i].Nombre, this.lista_tiendas.Tiendas[i].Descripcion, this.lista_tiendas.Tiendas[i].Contacto, this.lista_tiendas.Tiendas[i].Calificacion, this.lista_tiendas.Tiendas[i].Logo)
      this.listaAux.push(tienda);
      console.log(i)
    }
  }

  



}
