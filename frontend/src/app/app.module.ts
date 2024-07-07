import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import {HTTP_INTERCEPTORS}from '@angular/common/http';
import { ToastrModule } from 'ngx-toastr';
import { ReactiveFormsModule } from '@angular/forms';
import { CourtHomeComponent } from './court-home/court-home.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NavbarComponent } from './navbar/navbar.component';
import { CreateLawEtititiesComponent } from './create-law-etitities/create-law-etitities.component';

import { TrafficPoliceComponent } from './traffic-police.service/traffic-police.service.component';





import { MupvozilaComponent } from './mupvozila/mupvozila.component';
import { MupvozilaHomeComponent } from './mupvozila-home/mupvozila-home.component';
import { UserDashboardComponent } from './user-dashboard/user-dashboard.component';
import { TrafficPoliceDataComponent } from './traffic-police-data/traffic-police-data.component';
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
    MupvozilaCommunicationComponent


  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot(),
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    NgbModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: TokeninterceptorInterceptor,
      multi: true
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
