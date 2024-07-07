import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { CourtHomeComponent } from './court-home/court-home.component';
import { NavbarComponent } from './navbar/navbar.component';
import { CreateLawEtititiesComponent } from './create-law-etitities/create-law-etitities.component';

import { TrafficPoliceComponent } from './traffic-police.service/traffic-police.service.component';

import { MupvozilaComponent } from './mupvozila/mupvozila.component';
import { MupvozilaHomeComponent } from './mupvozila-home/mupvozila-home.component';
import { TrafficPoliceDataComponent } from './traffic-police-data/traffic-police-data.component';
import { UserDashboardComponent } from './user-dashboard/user-dashboard.component';
import { StatistikaPrekrsajaComponent } from './statistika-prekrsaja/statistika-prekrsaja.component';
import { StatistikaNesrecaComponent } from './statistika-nesreca/statistika-nesreca.component';
import { StatistikaVozackaComponent } from './statistika-vozacka/statistika-vozacka.component';
import { StatistikaRegistracijaComponent } from './statistika-registracija/statistika-registracija.component';
import { StatistikaAutaComponent } from './statistika-auta/statistika-auta.component';
import { StatistikaHomeComponent } from './statistika-home/statistika-home.component';
import { StatistikaSaslusanjaComponent } from './statistika-saslusanja/statistika-saslusanja.component';
import { StatistikaPretresaComponent } from './statistika-pretresa/statistika-pretresa.component';
import { StatistikaZahtevaComponent } from './statistika-zahteva/statistika-zahteva.component';

import { MupvozilaCommunicationComponent } from './mupvozila-communication/mupvozila-communication.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full',
  },
  {
    path: 'navbar',
    component: NavbarComponent,
  },
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'register',
    component: RegisterComponent,
  },
  {
    path: 'court-home',
    component: CourtHomeComponent,
  },
  {
    path: 'create-law-etitities',
    component: CreateLawEtititiesComponent,
  },
  {
    path: 'traffic-police',
    component: TrafficPoliceComponent,
  },
  {
    path: 'mupvozila',
    component: MupvozilaComponent,
  },

  { path: 'traffic-police-data', component: TrafficPoliceDataComponent },
  {
    path: 'mupvozila-home',
    component: MupvozilaHomeComponent,
  },

  {
    path: 'user-dashboard',
    component: UserDashboardComponent,
  },
  {
    path: 'statistika-prekrsaja',
    component: StatistikaPrekrsajaComponent,
  },
  {
    path: 'statistika-nesreca',
    component: StatistikaNesrecaComponent,
  },
  {
    path: 'statistika-vozackih',
    component: StatistikaVozackaComponent,
  },
  {
    path: 'statistika-registracija',
    component: StatistikaRegistracijaComponent,
  },
  {
    path: 'statistika-auta',
    component: StatistikaAutaComponent,
  },
  {
    path: 'statistika-home',
    component: StatistikaHomeComponent,
  },
  {
    path: 'statistika-saslusanja',
    component: StatistikaSaslusanjaComponent,
  },
  {
    path: 'statistika-pretresa',
    component: StatistikaPretresaComponent,
  },
  {
    path: 'statistika-zahteva',
    component: StatistikaZahtevaComponent,
  },

  { 
    path: 'user-dashboard', 
    component: UserDashboardComponent
  },
  {
    path: 'mupvozila-communication',
    component: MupvozilaCommunicationComponent
  }



];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
