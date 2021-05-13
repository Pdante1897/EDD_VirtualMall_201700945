import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {baseURL} from '../../apiURL/baseURL';

@Injectable({
  providedIn: 'root'
})
export class UploadService {

  constructor(private http: HttpClient) { }


  uploadFile(File: Document, peticion: string): Observable<any> {
    
    var json = JSON.stringify(File);
    console.log(File);
    var headers = new HttpHeaders().set("Content-Type", "application/json");
    
    return this.http.post(baseURL + peticion, File, { headers });
  }

  getTiendas():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getArreglo');
  }

  getInventarios():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getImgInv', httpOptions);
  }
  getUsuarios():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getUsuarios', httpOptions);
  }

  getUsuariosSha256():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getUsuariosSha256', httpOptions);
  }
  getUsuariosBcrypt():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getUsuariosBcrypt', httpOptions);
  }
  
  getMatriz():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getMatriz', httpOptions);
  }

  postTiempo(tiempo : string):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'postTiempo/'+tiempo, httpOptions);
  }

}
