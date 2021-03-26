import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {baseURL} from '../../apiURL/baseURL';
import { Observable } from "rxjs";
import { Tienda } from 'src/app/models/tiendas/tienda';


@Injectable({
  providedIn: 'root'
})
export class TiendasService {

  constructor(private http: HttpClient) {
    const httpOptions ={
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    }
   }

   getTiendas():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getTiendas', httpOptions);
  }
}
