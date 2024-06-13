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
   // this.getCurrentUserIdVehicles();
    this.getUserRegistrations();
    this.getUserLicenses();
    this.getAllVehicles();
  }

  // getCurrentUserIdVehicles() {
  //   let id = this.authService.getUserId();
  //   this.mupvozilaService.getVehiclesByUserId(id).subscribe((data: any) => {
  //     this.vehicles = data;
  //   });
  // }

  getAllVehicles() {
    let id = this.authService.getUserId();
    this.mupvozilaService.getAllVehicles().subscribe((data: any) => {
      this.vehicles = data;
    });
  }

  getUserRegistrations() {
    let id = this.authService.getUserId();
    this.mupvozilaService.getAllVehicles().subscribe((data: any) => {
      this.registrations = data;
    });
  }

  getUserLicenses() {
    let id = this.authService.getUserId();
    this.mupvozilaService.getAllLicenses().subscribe((data: any) => {
      this.licenses = data;
    });
  }

  deleteVehicle(carId: string) {
    this.mupvozilaService.deleteVehicle(carId).subscribe(() => {
      this.vehicles = this.vehicles.filter(vehicle => vehicle.id !== carId);
    });
  }
}