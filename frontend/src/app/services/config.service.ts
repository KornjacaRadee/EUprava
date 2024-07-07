import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  _api_url: string;
  _auth_url: string;
  _register_url: string;
  _get_user: string;

  _login_url: string;
  _court_url: string;

  _create_entity: string;
  _create_request: string;
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

  _mup_vozila_url: string;
  _create_licences: string;
  _create_vehicles: string;
  _getAllLicences: string;
  _getAllVehicles: string;
  _getLicenceById: string;
  _getVehicleById: string;
  _getLicenceByJMBG: string;
  _update_vehicle: string;

  _create_nesreca: string;
  _get_nesreca: string;
  _delete_nesreca: string;
  _court_urll: string;

  _create_statistic_prekrsaj: string;
  _create_statistic_nesreca: string;
  _create_statistic_vozacka: string;
  _create_statistic_registracija: string;
  _create_statistic_auta: string;
  _create_statistic_pretres: string;
  _crate_statistic_saslusanje: string;
  _create_statistic_pravni_zahtev: string;
  _create_statistic_procenat_np: string;
  _crate_statistic_procenat_vr: string;
  _create_statistic_procenat_ra: string;
  _create_statistic_procenat_pn: string;
  _crate_statistic_procenat_rv: string;
  _statistic_url: string;

  constructor() {
    this._api_url = 'http://localhost';
    this._auth_url = this._api_url + ':8000';
    this._court_url = this._api_url + ':8000/api/law_court';

    this._court_urll = this._api_url + ':8084';

    this._mup_vozila_url = this._api_url + ':8081';

    this._statistic_url = this._api_url + ':8085';

    //AUTH ROUTES

    this._register_url = this._auth_url + '/api/auth/register';

    this._login_url = this._auth_url + '/api/auth/login';

    this._get_user = this._auth_url + '/api/auth/users/';

    //COURT ROUTES

    this._create_request = this._court_url + '/legal_requests';
    this._create_entity = this._court_url + '/legal_entities';
    this._get_user_entities = this._court_url + '/legal_entities/user/';
    this._get_entity = this._court_url + '/legal_entities/';

    this._create_warrant = this._court_url + '/house_search_warrants';
    this._get_user_warrant = this._court_url + '/house_search_warrants/user/';
    this._get_warrant = this._court_url + '/house_search_warrants/';

    this._create_hearing = this._court_url + '/hearings';
    this._get_user_hearing = this._court_url + '/hearings/user/';
    this._get_hearing = this._court_url + '/hearings/';

    this._create_prekrsaj = this._court_urll + '/prekrsaj/new';
    this._delete_prekrsaj = this._court_urll + '/prekrsaj/';
    this._get_prekrsaj = this._court_urll + '/prekrsaj';

    this._create_nesreca = this._court_urll + '/nesreca/new';
    this._delete_nesreca = this._court_urll + '/nesreca/';
    this._get_nesreca = this._court_urll + '/nesreca';

    // MUP VOZILA ROUTES

    this._create_licences = this._mup_vozila_url + '/licenses';
    this._create_vehicles = this._mup_vozila_url + '/cars';
    this._getAllLicences = this._mup_vozila_url + '/getAllLicenses';
    this._getAllVehicles = this._mup_vozila_url + '/getAllCars';
    this._getLicenceById = this._mup_vozila_url + '/getLicenseById/user/';
    this._getVehicleById = this._mup_vozila_url + '/getVehicleById/user/';

    this._getLicenceByJMBG = this._mup_vozila_url + '/licenses/user';
    this._update_vehicle = this._mup_vozila_url + '/cars/';

    //STATISTIKA ROUTES

    this._create_statistic_prekrsaj =
      this._statistic_url + '/statistikaPrekrsaja';
    this._create_statistic_nesreca = this._statistic_url + '/statistikaNesreca';
    this._create_statistic_vozacka =
      this._statistic_url + '/statistikaVozackihDozvola';
    this._create_statistic_registracija =
      this._statistic_url + '/statistikaRegistrovanihVozila';
    this._create_statistic_auta = this._statistic_url + '/statistikaAuta';
    this._create_statistic_pretres =
      this._statistic_url + '/statistikaNalogaZaPretres';
    this._crate_statistic_saslusanje =
      this._statistic_url + '/statistikaSaslusanja';
    this._create_statistic_pravni_zahtev =
      this._statistic_url + '/statistikaPravnogZahteva';
    this._create_statistic_procenat_np =
      this._statistic_url + '/calculatePercentageNP';
    this._crate_statistic_procenat_vr =
      this._statistic_url + '/calculatePercentageVR';
    this._create_statistic_procenat_ra =
      this._statistic_url + '/calculatePercentageRA';
    this._create_statistic_procenat_pn =
      this._statistic_url + '/calculatePercentagePN';
    this._crate_statistic_procenat_rv =
      this._statistic_url + '/calculatePercentageRV';
  }
}
