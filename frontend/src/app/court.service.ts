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

  getUserEntities(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_user_entities}`+id);
  }

  getEntity(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_entity}`+id);
  }

  getUserWarrants(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_user_warrant}`+id);
  }

  getWarrant(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_warrant}`+id);
  }

  getUserHearings(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_user_hearing}`+id);
  }

  getHearing(id: string): Observable<any> {
    return this.http.get(`${this.configService._get_hearing}`+id);
  }
}
