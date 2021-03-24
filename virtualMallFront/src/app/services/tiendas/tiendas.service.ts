import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {baseURL} from '../../apiURL/baseURL';

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

   
}
