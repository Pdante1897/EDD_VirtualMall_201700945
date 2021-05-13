import { conditionallyCreateMapObjectLiteral } from '@angular/compiler/src/render3/view/util';
import { Component, OnInit } from '@angular/core';
import { UploadService } from 'src/app/services/upload/upload.service';

@Component({
  selector: 'app-cargar',
  templateUrl: './cargar.component.html',
  styleUrls: ['./cargar.component.css']
})
export class CargarComponent implements OnInit {
  archivo: any
  cargarTiendas= "cargartienda"
  cargarUsuarios = "cargarUsuarios"
  cargarPedidos = "cargarpedido"
  cargarInventarios = "cargarinventario"
  cargarGrafos = "cargarGrafos"
  key: string = ""
  tiempo: string = ""

  constructor(private service: UploadService) { }
  
  ngOnInit(): void {
  }
  
  getTiendas(){
    console.log('apachurrado')

    this.service.getTiendas().subscribe(Response => {null});
  }
  getUsuarios(){
    console.log('apachurrado')

    this.service.getUsuarios().subscribe(Response => {null});
  }
  getInventarios(){
    console.log('apachurrado')

    this.service.getInventarios().subscribe(Response => {null});
  }
  getMatriz(){
    console.log('apachurrado')

    this.service.getMatriz().subscribe(Response => {null});
  }


  leerArchivo(e: any) {
    var archivo = e.target.files[0];
    if (!archivo) {
      return;
    }
    var lector = new FileReader();
    lector.onload = function(e: any) {
      var contenido = e.target.result;
    };
    lector.readAsText(archivo);
    this.archivo= archivo
  }
  
  subirArchivo(archivo: any, peticion: string) {
    console.log(this.archivo)
    this.service.uploadFile(this.archivo, peticion).subscribe(Response => {this.archivo});
  }

  setKey(a:any){
    localStorage.setItem("key",this.key)
    console.log(localStorage.getItem("key"))

  }
  setTiempo(a:any){
    console.log(this.tiempo)
    this.service.postTiempo(this.tiempo).subscribe(Response => {null});
  }
  
//ng serve --live-reload false
  
}


