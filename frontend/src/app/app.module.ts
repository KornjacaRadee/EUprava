import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { ToastrModule } from 'ngx-toastr';
import { ReactiveFormsModule } from '@angular/forms';
import { CourtHomeComponent } from './court-home/court-home.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NavbarComponent } from './navbar/navbar.component';
import { CreateLawEtititiesComponent } from './create-law-etitities/create-law-etitities.component';

import { TrafficPoliceComponent } from './traffic-police.service/traffic-police.service.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatTableModule } from '@angular/material/table';
import { MupvozilaComponent } from './mupvozila/mupvozila.component';
import { MupvozilaHomeComponent } from './mupvozila-home/mupvozila-home.component';
import { UserDashboardComponent } from './user-dashboard/user-dashboard.component';
import { TrafficPoliceDataComponent } from './traffic-police-data/traffic-police-data.component';

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
import { TokeninterceptorInterceptor } from './services/interceptors/tokeninterceptor.interceptor';

@NgModule({
  declarations: [
    AppComponent,
    RegisterComponent,
    LoginComponent,
    CourtHomeComponent,
    NavbarComponent,
    CreateLawEtititiesComponent,

    TrafficPoliceComponent,

    MupvozilaComponent,
    MupvozilaHomeComponent,
    UserDashboardComponent,
    TrafficPoliceDataComponent,

    StatistikaPrekrsajaComponent,
    StatistikaNesrecaComponent,
    StatistikaVozackaComponent,
    StatistikaRegistracijaComponent,
    StatistikaAutaComponent,
    StatistikaHomeComponent,
    StatistikaSaslusanjaComponent,
    StatistikaPretresaComponent,
    StatistikaZahtevaComponent,

    MupvozilaCommunicationComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot(),
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    NgbModule,
    MatButtonModule,
    MatToolbarModule,
    MatInputModule,
    MatTableModule,
  ],

  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: TokeninterceptorInterceptor,
      multi: true,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
