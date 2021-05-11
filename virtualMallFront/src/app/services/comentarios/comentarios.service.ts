import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { baseURL } from 'src/app/apiURL/baseURL';

@Injectable({
  providedIn: 'root'
})
export class ComentariosService {

  constructor(private http: HttpClient) { }
  getComentarios(nombre: string, departamento: string, calificacion: string):Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getComentarios/'+nombre.replace('%20','-')+'+'+departamento.replace('%20','-')+'+'+calificacion, httpOptions);
  }
  getComentariosProd(nombre: string, departamento: string, calificacion: string, producto: string):Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getComentariosProd/'+nombre.replace('%20','-')+'+'+departamento.replace('%20','-')+'+'+calificacion+"+"+producto, httpOptions);
  }
  postComentario(nombre: string, departamento: string, calificacion: string ,comentario: any):Observable<any>{
    var file = JSON.stringify(comentario)
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json', 
      }),
    };
    return this.http.post<any>(baseURL + 'postComentario/'+nombre.replace('%20','-')+'+'+departamento.replace('%20','-')+'+'+calificacion, comentario );
  }
  postComentarioProd(nombre: string, departamento: string, calificacion: string, producto: string ,comentario: any):Observable<any>{
    var file = JSON.stringify(comentario)
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json', 
      }),
    };
    return this.http.post<any>(baseURL + 'postComentarioProd/'+nombre.replace('%20','-')+'+'+departamento.replace('%20','-')+'+'+calificacion+"+"+producto, comentario );
  }

  postSubComentarioProd(nombre: string, departamento: string, calificacion: string, producto: string ,comentario: any):Observable<any>{
    var file = JSON.stringify(comentario)
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json', 
      }),
    };
    return this.http.post<any>(baseURL + 'postSubComentarioProd/'+nombre.replace('%20','-')+'+'+departamento.replace('%20','-')+'+'+calificacion+"+"+producto, comentario );
  }

  postSubComentario(nombre: string, departamento: string, calificacion: string,comentario: any):Observable<any>{
    var file = JSON.stringify(comentario)
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json', 
      }),
    };
    return this.http.post<any>(baseURL + 'postSubComentario/'+nombre.replace('%20','-')+'+'+departamento.replace('%20','-')+'+'+calificacion, comentario );
  }
}
