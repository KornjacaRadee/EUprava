import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ConfigService } from './services/config.service';

interface License {
  user_jmbg: string;
  category: string;
  issuing_date: string;
  valid_until_date: string;
  address: string;
  points: number;
  is_valid: boolean;
}

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


  getAllVehicles(): Observable<any> {
    return this.http.get(`${this.configService._getAllVehicles}`);
  }

  getAllRegistrations(): Observable<any> {
    return this.http.get(`${this.configService._mup_vozila_url}/registrations`);
  }


  getLicenseById(id: string): Observable<any> {
    return this.http.get(`${this.configService._getLicenceById}` + id);
  }

  getVehicleById(id: string): Observable<any> {
    return this.http.get(`${this.configService._getVehicleById}` + id);
  }

  // getVehiclesByUserId(userId: string): Observable<any> {
  //   return this.http.get(`${this.configService._mup_vozila_url}/cars/user/${userId}`);
  // }

  registerVehicle(registration: any): Observable<any> {
    return this.http.post(`${this.configService._mup_vozila_url}/vehicles/register`, registration);
  }

  deleteVehicle(carId: string): Observable<any> {
    return this.http.delete(`${this.configService._mup_vozila_url}/cars/${carId}`);
  }

  deleteRegistration(registrationId: string): Observable<any> {
    return this.http.delete(`${this.configService._mup_vozila_url}/registrations/${registrationId}`);
  }

  deleteLicense(licenseId: string): Observable<any> {
    return this.http.delete(`${this.configService._mup_vozila_url}/licenses/${licenseId}`);
  }

  issueLicense(license: any): Observable<any> {
    return this.http.post<any>(this.configService._create_licences, license);
  }

  getAllLicenses(): Observable<any> {
    return this.http.get<any>(this.configService._getAllLicences);
  }

  //  getLicensesByUserJMBG(jmbg: string): Observable<any> {
  //    return this.http.get<any>(`${this.configService._getLicenceByJMBG}/${jmbg}`);
  //  }

  updateCar(car: any): Observable<any> {
    return this.http.put(`${this.configService._update_vehicle}${car.id}`, car);
  }

  updateRegistration(registration: any): Observable<any> {
    return this.http.put(`${this.configService._mup_vozila_url}/registrations/${registration.id}`, registration);
  }

  updateLicense(license: any): Observable<any> {
    return this.http.put(`${this.configService._mup_vozila_url}/licenses/${license.id}`, license);
  }

  // getVehiclesByUserId(userId: string): Observable<any> {
  //   console.log('Requesting vehicles for user:', userId);
  //   return this.http.get(`${this.configService._mup_vozila_url}/cars/user/${userId}`);
  // }

  getLicensesByUserJMBG(jmbg: string): Observable<any> {
    console.log('Requesting licenses for JMBG:', jmbg);
    return this.http.get(`${this.configService._getLicenceByJMBG}/${jmbg}`);
  }

  getVehiclesByUserId(userId: string): Observable<any> {
    console.log('Requesting vehicles for user:', userId);
    return this.http.get(`${this.configService._mup_vozila_url}/cars/user/${userId}`);
  }

  getLicensesByUserId(userId: string): Observable<any> {
    console.log('Requesting licenses for user:', userId);
    return this.http.get<any>(`${this.configService._mup_vozila_url}/licenses/user/${userId}`);
  }

  getAccidentsByDriver(driverId: string): Observable<any> {
    return this.http.get(`${this.configService._mup_vozila_url}/nesrece/vozac/${driverId}`);
  }
}
