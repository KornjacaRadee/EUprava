import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { CourtHomeComponent } from './court-home/court-home.component';
import { NavbarComponent } from './navbar/navbar.component';
import { CreateLawEtititiesComponent } from './create-law-etitities/create-law-etitities.component';
import { TrafficPoliceComponent } from './traffic-police.service/traffic-police.service.component';

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
    component: TrafficPoliceComponent
  }


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
