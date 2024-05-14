import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient,HttpHeaders } from '@angular/common/http';
import { ConfigService } from './services/config.service';
import { Observable,tap } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class CourtService {

  constructor(
    private http: HttpClient,
    private configService:ConfigService,
    private router: Router

  ) {

  }
  createEntity(entity: any): Observable<any> {
    return this.http.post(`${this.configService._create_entity}`, entity);
  }

  getUserEntities(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_user_entities}`+id);
  }

  getEntity(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_entity}`+id);
  }

  getUserWarrants(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_user_warrant}`+id);
  }

  createWarrant(warrant: any): Observable<any> {
    return this.http.post(`${this.configService._create_warrant}`, warrant);
  }

  getWarrant(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_warrant}`+id);
  }

  getUserHearings(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_user_hearing}`+id);
  }

  createHearing(hearing: any): Observable<any> {
    return this.http.post(`${this.configService._create_hearing}`, hearing);
  }

  getHearing(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_hearing}`+id);
  }
}
