import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TrafficPoliceService {
  private _courtUrl = 'http://localhost:8084'; 

  private _createPrekrsajUrl = this._courtUrl + '/prekrsaj/new';
  private _deletePrekrsajUrl = this._courtUrl + '/prekrsaj/{id}';
  private _getPrekrsajUrl = this._courtUrl + '/prekrsaji';

  private _createNesrecaUrl = this._courtUrl + '/nesreca/new';
  private _deleteNesrecaUrl = this._courtUrl + '/nesreca/{id}';
  private _getNesrecaUrl = this._courtUrl + '/nesreca';

  constructor(private http: HttpClient) { }

  createPrekrsaj(prekrsajData: any): Observable<any> {
    return this.http.post(this._createPrekrsajUrl, prekrsajData);
  }

  deletePrekrsaj(prekrsajId: string): Observable<any> {
    const url = `${this._deletePrekrsajUrl.replace('{id}', prekrsajId)}`;
    return this.http.delete(url);
  }

  getPrekrsaji(): Observable<any> {
    return this.http.get(this._getPrekrsajUrl);
  }

  createNesreca(nesrecaData: any): Observable<any> {
    return this.http.post(this._createNesrecaUrl, nesrecaData);
  }

  deleteNesreca(nesrecaId: string): Observable<any> {
    const url = `${this._deleteNesrecaUrl.replace('{id}', nesrecaId)}`;
    return this.http.delete(url);
  }

  getNesrece(): Observable<any> {
    return this.http.get(this._getNesrecaUrl);
  }
}
