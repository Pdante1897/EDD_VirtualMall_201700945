import { Component, OnInit } from '@angular/core';
import { PdfsService } from 'src/app/services/pdfs/pdfs.service';
import { ArrPdf, Pdf } from 'src/app/models/pdfs/pdf';
import { DomSanitizer } from '@angular/platform-browser';
import jsPDF from 'jspdf';
import { Router } from '@angular/router';

declare var require: any
const FileSaver = require('file-saver');

const url = require('url');



@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {
  lista_PDFS: ArrPdf = new ArrPdf()
  listaAux: Pdf[] = []
  name = 'Angular 5';
 

  constructor(private pdfservice:PdfsService, private sanitizer: DomSanitizer,private router:Router) { 
    
  }

  ngOnInit(): void {
    if (localStorage.getItem("user")!="Admin") {
      this.router.navigate(['/login']);    
    }
    this.cargar()
  }

  

  cargar(){
    this.pdfservice.getPDFs().subscribe((dataList: any)=>{
      this.lista_PDFS = dataList 
      console.log(this.lista_PDFS)
      console.log(Object.getOwnPropertyDescriptors(this.lista_PDFS))
      console.log(Object.is(this.lista_PDFS,ArrPdf))
      console.log(dataList)
      this.setearLista()
      console.log(this.listaAux)
    },(err)=>{
      console.log("no se pudo")

    })
  
  }

  setearLista(){
     for (let i = 0; i < this.lista_PDFS.Archivos.length; i++) {
      var pdf: Pdf
      
      pdf = new Pdf(this.lista_PDFS.Archivos[i].Nombre, this.lista_PDFS.Archivos[i].Ruta)
      this.listaAux.push(pdf);
      console.log(i);
    }
  }

  descargar(pdfUrl: string, pdfName: string ) {
    FileSaver.saveAs("./assets/download/"+pdfName, pdfName);
  }

  

  openDoc(pdfUrl: string, startPage: number ) {
    window.open(pdfUrl + '#page=' + startPage, '_blank', '', true);
  }

}
