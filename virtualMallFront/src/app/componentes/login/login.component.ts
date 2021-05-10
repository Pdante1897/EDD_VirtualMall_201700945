import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BusquedaU, Usuario } from 'src/app/models/usuario/usuario';
import { UsuarioService } from 'src/app/services/usuario/usuario.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  usuario: string = ""
  password: string = ""


  constructor(private userservice:UsuarioService, private router:Router) { }

  ngOnInit(): void {
    
  }

  user(dato: string){
        console.log(this.usuario)
  }
  login(user: string, password: string){
    console.log(this.usuario)
    var Dpi = parseInt(user)
    var usuario = new BusquedaU(Dpi)
    console.log(usuario)
    var busqueda : Usuario
    
    var file = JSON.stringify(usuario)
    console.log(file)
    this.userservice.getUser(file).subscribe((dataList: any)=>{
      busqueda = dataList 
      console.log(busqueda)
      if (busqueda.Dpi==Dpi && busqueda.Password == password) {
        if (busqueda.Cuenta == "Admin") {
          this.router.navigate(['admin'])
          localStorage.setItem("user", "Admin")
          localStorage.setItem("dpi", busqueda.Dpi.toString())
          alert("Bienvenido Admin")

        }else{
          this.router.navigate(['/'])
          localStorage.setItem("user", busqueda.Nombre)
          localStorage.setItem("dpi", busqueda.Dpi.toString())
          alert("Bienvenido Usuario")
        }
      }
      else{
        alert("no se pudo")

      }
    },(err)=>{
    console.log("no se pudo")

    })
    
    
    if (Dpi == 1234567890101 && password=="1234"){
      localStorage.setItem("dpi", Dpi.toString())
      this.router.navigate(['admin'])
      localStorage.setItem("user", "Admin")
      alert("Bienvenido Admin")

    }
  }

}


