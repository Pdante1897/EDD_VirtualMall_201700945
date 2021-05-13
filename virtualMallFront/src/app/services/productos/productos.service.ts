import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {baseURL} from '../../apiURL/baseURL';
import { Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class ProductosService {

  constructor(private http: HttpClient) {

   }

  getProductos(nombre: string, departamento: string, calificacion: string):Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getProductos/'+nombre.replace('%20','-')+'+'+departamento.replace('%20','-')+'+'+calificacion, httpOptions);
  }

  postPedido(archivo: any){
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'postPedido', archivo);
  }
}
