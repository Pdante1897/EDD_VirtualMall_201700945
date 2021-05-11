import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Comentario, Comentarios, SubComentarios } from 'src/app/models/comentarios/comentarios';
import { ComentariosService } from 'src/app/services/comentarios/comentarios.service';

@Component({
  selector: 'app-comentarios',
  templateUrl: './comentarios.component.html',
  styleUrls: ['./comentarios.component.css']
})
export class ComentariosComponent implements OnInit {
  comentariosE: Comentarios = new Comentarios
  listaCom: Comentario[]=[]
  cadena : string = ""
  nombre : string = ""
  listaV: Comentario[]=[]
  comNull: Comentario= new Comentario(0,0,"", this.listaV)
  constructor(private router: Router, private route: ActivatedRoute, private comentariosService: ComentariosService) { }

  ngOnInit(): void {
    var nombre = (this.route.snapshot.paramMap.get('nombre') || '')
    var departamento = (this.route.snapshot.paramMap.get('departamento')|| '')
    var calificacion = (this.route.snapshot.paramMap.get('calificacion')|| '')
    this.nombre= nombre+" -"+departamento+" -"+calificacion
    this.comentariosService.getComentarios(nombre, departamento, calificacion).subscribe((dataList: any)=>{
      this.comentariosE = dataList 
      console.log(this.comentariosE)
      console.log(Object.getOwnPropertyDescriptors(this.comentariosE))
      console.log(Object.is(this.comentariosE,Comentarios))
      console.log(dataList)
      if (this.comentariosE!=null) {
        this.setearLista()

      }
      console.log(this.listaCom)
    },(err)=>{
      console.log("no se pudo")

    })
  }
  setearLista(){
    

    for (let i = 0; i < this.comentariosE.Comentarios.length; i++) {
      var com: Comentario     
      com = new Comentario(this.comentariosE.Comentarios[i].Id,this.comentariosE.Comentarios[i].Dpi,this.comentariosE.Comentarios[i].Cadena,this.comentariosE.Comentarios[i].SubComentarios)
      this.listaCom.push(com);
      console.log(i)
    }
  }
  comentar(cadena: string){
    var nombre = (this.route.snapshot.paramMap.get('nombre') || '')
    var departamento = (this.route.snapshot.paramMap.get('departamento')|| '')
    var calificacion = (this.route.snapshot.paramMap.get('calificacion')|| '')
    var id= parseInt(localStorage.getItem("dpi")||"0")
    var sub: Comentario[]=[]

    try {
      var comentarioN = new Comentario(this.comentariosE.Comentarios.length, id, cadena, sub) 
    } catch (error) {
      var comentarioN = new Comentario(0, id, cadena, sub) 
    }
       
    
   
    console.log(comentarioN)
    var file = JSON.stringify(comentarioN)
    console.log(file)
    this.comentariosService.postComentario(nombre, departamento, calificacion, file).subscribe((dataList: any)=>{
      location.reload()
    },(err)=>{
    console.log("no se pudo")
    })
  }


  comentarSub(com:Comentario, com2:Comentario, com3:Comentario, com4:Comentario, com5:Comentario,com6:Comentario ,cadena : string){
    var nombre = (this.route.snapshot.paramMap.get('nombre') || '')
    var departamento = (this.route.snapshot.paramMap.get('departamento')|| '')
    var calificacion = (this.route.snapshot.paramMap.get('calificacion')|| '')

    var id= parseInt(localStorage.getItem("dpi")||"0")
    var sub: Comentario[]=[]
    var subCom : SubComentarios = new SubComentarios()
    subCom.Comentarios.push(com)
    if (com2!=this.comNull) {
      subCom.Comentarios.push(com2)

    }    
    if (com3!=this.comNull) {
      subCom.Comentarios.push(com3)

    }if (com4!=this.comNull) {
      subCom.Comentarios.push(com4)

    }
    if (com5!=this.comNull) {
      subCom.Comentarios.push(com5)
    }
    if (com6!=this.comNull) {
        subCom.Comentarios.push(com6)
  
    }

      var comentarioN = new Comentario(0, id, cadena, sub) 
    subCom.Comentarios.push(comentarioN)
       
    
    console.log(subCom)
    var file = JSON.stringify(subCom)
    console.log(file)
    this.comentariosService.postSubComentario(nombre, departamento, calificacion, file).subscribe((dataList: any)=>{
      location.reload()
    },(err)=>{
    console.log("no se pudo")
    })
  }
}
