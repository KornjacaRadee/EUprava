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

  private _getUserByJmbgUrl = this._courtUrl + '/users/jmbg/{jmbg}';
  private _getCarByLicensePlateUrl = this._courtUrl + '/cars/plate/{license_plate}';

  private _getLicensesByUserJMBGUrl = this._courtUrl + '/licenses/user/{jmbg}';

  private _getNesreceByVozacUrl = this._courtUrl + '/nesrece/vozac/{vozac}';

  private _getAllCarsUrl = this._courtUrl + '/cars';

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

  getUserByJMBG(jmbg: string): Observable<any> { // Dodato
    const url = `${this._getUserByJmbgUrl.replace('{jmbg}', jmbg)}`;
    return this.http.get(url);
  }

  getCarByLicensePlate(licensePlate: string): Observable<any> {
    const url = `${this._getCarByLicensePlateUrl.replace('{license_plate}', licensePlate)}`;
    return this.http.get(url);
  }

  getLicensesByUserJMBG(jmbg: string): Observable<any> {
    const url = `${this._getLicensesByUserJMBGUrl.replace('{jmbg}', jmbg)}`;
    return this.http.get(url);
  }

  getAllCars(): Observable<any> {
    return this.http.get(this._getAllCarsUrl);
  }

  getNesreceByVozac(vozac: string): Observable<any> {
    const url = `${this._getNesreceByVozacUrl.replace('{vozac}', vozac)}`;
    return this.http.get(url);
  }


}
