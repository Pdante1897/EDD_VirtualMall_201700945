import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  usuario: string = ""
  password: string = ""


  constructor(private router:Router) { }

  ngOnInit(): void {
  }

  user(dato: string){
        console.log(this.usuario)
  }
  login(usuario: string, password: string){
    console.log(this.usuario)
    var user = parseInt(usuario)
    if (user == 1234567890101 && password=="1234"){
      this.router.navigate(['admin'])
    }
  }

}


