import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {


  _api_url: string;
  _auth_url: string;
  _register_url: string;
  _login_url: string;



  constructor() {
    this._api_url = 'http://localhost';
    this._auth_url =this._api_url + ':8082';


    this._register_url = this._auth_url + '/register';

    this._login_url = this._auth_url + '/login';





  }
}
