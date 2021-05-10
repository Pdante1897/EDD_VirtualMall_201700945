import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Comentario, Comentarios } from 'src/app/models/comentarios/comentarios';
import { ComentariosService } from 'src/app/services/comentarios/comentarios.service';

@Component({
  selector: 'app-comentarios-prod',
  templateUrl: './comentarios-prod.component.html',
  styleUrls: ['./comentarios-prod.component.css']
})
export class ComentariosProdComponent implements OnInit {

  comentariosE: Comentarios = new Comentarios
  listaCom: Comentario[]=[]
  cadena : string = ""
  nombre : string = ""
  constructor(private router: Router, private route: ActivatedRoute, private comentariosService: ComentariosService) { }

  ngOnInit(): void {
    var nombre = (this.route.snapshot.paramMap.get('nombre') || '')
    var departamento = (this.route.snapshot.paramMap.get('departamento')|| '')
    var calificacion = (this.route.snapshot.paramMap.get('calificacion')|| '')
    var producto = (this.route.snapshot.paramMap.get('producto')|| '')
    var productoN = (this.route.snapshot.paramMap.get('nombreP')|| '')


    this.nombre= nombre+" -"+departamento+" -"+calificacion +" -"+productoN
    this.comentariosService.getComentariosProd(nombre, departamento, calificacion, producto).subscribe((dataList: any)=>{
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
      com = new Comentario(this.comentariosE.Comentarios[i].Id,this.comentariosE.Comentarios[i].Cadena)
      this.listaCom.push(com);
      console.log(i)
    }
  }
  comentar(cadena: string){
    var nombre = (this.route.snapshot.paramMap.get('nombre') || '')
    var departamento = (this.route.snapshot.paramMap.get('departamento')|| '')
    var calificacion = (this.route.snapshot.paramMap.get('calificacion')|| '')
    var producto = (this.route.snapshot.paramMap.get('producto')|| '')

    var id= parseInt(localStorage.getItem("dpi")||"0")
    var comentarioN = new Comentario(id, cadena) 
    console.log(comentarioN)
    var file = JSON.stringify(comentarioN)
    console.log(file)
    this.comentariosService.postComentarioProd(nombre, departamento, calificacion, producto, file).subscribe((dataList: any)=>{
      location.reload()
    },(err)=>{
    console.log("no se pudo")
    })
  }
}