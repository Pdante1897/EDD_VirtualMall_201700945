import { Component, OnInit } from '@angular/core';
import { UploadService } from 'src/app/services/upload/upload.service';

@Component({
  selector: 'app-encriptar',
  templateUrl: './encriptar.component.html',
  styleUrls: ['./encriptar.component.css']
})
export class EncriptarComponent implements OnInit {
  key:string =""
  constructor(private service: UploadService) { }
  
  ngOnInit(): void {
  }
  validarKey(a: any){
    if(this.key=="HolaMundo!EDD_AunSale_2021"|| this.key==localStorage.getItem("key")){
      this.getUsuariosBcrypt()
      this.getUsuariosSha256()
      window.alert("Si se pudo!");
    }else{
      window.alert("Contrasenia erronea!");
    }
  }
  getUsuariosBcrypt(){
    console.log('apachurrado')

    this.service.getUsuariosBcrypt().subscribe(Response => {null});
  }
  getUsuariosSha256(){
    console.log('apachurrado')

    this.service.getUsuariosSha256().subscribe(Response => {null});
  }
}
