import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ConfigService } from './services/config.service';

@Injectable({
  providedIn: 'root'
})
export class MupvozilaService {

  constructor(
    private http: HttpClient,
    private configService: ConfigService
  ) { }

  createLicense(license: any): Observable<any> {
    return this.http.post(`${this.configService._create_licences}`, license);
  }

  createVehicle(vehicle: any): Observable<any> {
    return this.http.post(`${this.configService._create_vehicles}`, vehicle);
  }

  getAllLicenses(): Observable<any> {
    return this.http.get(`${this.configService._getAllLicences}`);
  }

  getAllVehicles(): Observable<any> {
    return this.http.get(`${this.configService._getAllVehicles}`);
  }

  getLicenseById(id: string): Observable<any> {
    return this.http.get(`${this.configService.getLicenceById}` + id);
  }

  getVehicleById(id: string): Observable<any> {
    return this.http.get(`${this.configService.getVehicleById}` + id);
  }
}
