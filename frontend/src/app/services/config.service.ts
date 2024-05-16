import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {


  _api_url: string;
  _auth_url: string;
  _register_url: string;


  _login_url: string;
  _court_url: string;

  _create_entity: string;
  _get_user_entities: string;
  _get_entity: string;

  _create_warrant: string;
  _get_user_warrant: string;
  _get_warrant: string;

  _create_hearing: string;
  _get_user_hearing: string;
  _get_hearing: string;

  _create_prekrsaj: string;
  _get_prekrsaj: string;
  _delete_prekrsaj: string;

  _create_nesreca: string;
  _get_nesreca: string;
  _delete_nesreca: string;
  _court_urll: string;
  

  

  constructor() {
    this._api_url = 'http://localhost';
    this._auth_url =this._api_url + ':8082';
    this._court_url =this._api_url + ':8083';
    this._court_urll = this._api_url + ':8084';

//AUTH ROUTES

    this._register_url = this._auth_url + '/register';

    this._login_url = this._auth_url + '/login';

//COURT ROUTES

    this._create_entity = this._court_url + '/legal_entities';
    this._get_user_entities = this._court_url + '/legal_entities/user/';
    this._get_entity = this._court_url + '/legal_entities/';

    this._create_warrant = this._court_url + '/house_search_warrants';
    this._get_user_warrant = this._court_url + '/house_search_warrants/user/';
    this._get_warrant = this._court_url + '/house_search_warrants/';

    this._create_hearing = this._court_url + '/hearings'
    this._get_user_hearing = this._court_url + '/hearings/user/';
    this._get_hearing = this._court_url + '/hearings/';

    this._create_prekrsaj = this._court_urll + '/prekrsaj/new'; 
    this._delete_prekrsaj = this._court_urll + '/prekrsaj/'
    this._get_prekrsaj = this._court_urll + '/prekrsaj';

    this._create_nesreca = this._court_urll + '/nesreca/new';
    this._delete_nesreca = this._court_urll + '/nesreca/'
    this._get_nesreca = this._court_urll + '/nesreca';










  }
}
