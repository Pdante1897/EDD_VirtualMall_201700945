import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {baseURL} from '../../apiURL/baseURL';
import { Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class PdfsService {

  constructor(private http: HttpClient) { 

  }
  getPDFs():Observable<any>{

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'getPDFs', httpOptions);
  }
}
