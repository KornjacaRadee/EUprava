import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ConfigService } from './services/config.service';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class StatisticsService {
  private baseUrl = 'http://localhost:8085';

  constructor(private http: HttpClient) {}

  getStatistikaPrekrsaja(): Observable<any> {
    const url = `${this.baseUrl}/statistikaPrekrsaja`;
    return this.http.get(url);
  }

  getStatistikaNesreca(): Observable<any> {
    const url = `${this.baseUrl}/statistikaNesreca`;
    return this.http.get(url);
  }

  getStatistikaVozackihDozvola(): Observable<any> {
    const url = `${this.baseUrl}/statistikaVozackihDozvola`;
    return this.http.get(url);
  }

  getStatistikaRegistrovanihVozila(): Observable<any> {
    const url = `${this.baseUrl}/statistikaRegistrovanihVozila`;
    return this.http.get(url);
  }

  getStatistikaAuta(): Observable<any> {
    const url = `${this.baseUrl}/statistikaAuta`;
    return this.http.get(url);
  }

  getStatistikaNalogaZaPretres(): Observable<any> {
    const url = `${this.baseUrl}/statistikaNalogaZaPretres`;
    return this.http.get(url);
  }

  getStatistikaSaslusanja(): Observable<any> {
    const url = `${this.baseUrl}/statistikaSaslusanja`;
    return this.http.get(url);
  }

  getStatistikaPravnogZahteva(): Observable<any> {
    const url = `${this.baseUrl}/statistikaPravnogZahteva`;
    return this.http.get(url);
  }
}
