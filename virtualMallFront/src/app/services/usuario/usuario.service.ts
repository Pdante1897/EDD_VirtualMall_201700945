import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { baseURL } from 'src/app/apiURL/baseURL';

@Injectable({
  providedIn: 'root'
})
export class UsuarioService {

  constructor(private http: HttpClient) { }
  getUser(Dpi: any):Observable<any>{
    var file = JSON.stringify(Dpi)
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json', 
      }),
    };
    return this.http.post<any>(baseURL + 'buscarUser', Dpi );
  }
}
