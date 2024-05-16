import { Component, OnInit } from '@angular/core';
import { MupvozilaService } from '../mupvozila.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-mupvozila-home',
  templateUrl: './mupvozila-home.component.html',
  styleUrls: ['./mupvozila-home.component.css']
})
export class MupvozilaHomeComponent implements OnInit {

  vehicles: any[] = [];
  registrations: any[] = [];
  licenses: any[] = [];

  constructor(private mupvozilaService: MupvozilaService, private authService: AuthService) { }

  ngOnInit(): void {
    this.getUserVehicles();
    this.getUserRegistrations();
    this.getUserLicenses();
  }

  getUserVehicles() {
    let id = this.authService.getUserId();
    this.mupvozilaService.getVehicleById(id).subscribe((data: any) => {
      this.vehicles = data;
    });
  }

  getUserRegistrations() {
    let id = this.authService.getUserId();
    this.mupvozilaService.getVehicleById(id).subscribe((data: any) => {
      this.registrations = data;
    });
  }

  getUserLicenses() {
    let id = this.authService.getUserId();
    this.mupvozilaService.getLicenseById(id).subscribe((data: any) => {
      this.licenses = data;
    });
  }
}