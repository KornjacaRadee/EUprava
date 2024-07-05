import { Component, OnInit } from '@angular/core';
import { MupvozilaService } from '../mupvozila.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-user-dashboard',
  templateUrl: './user-dashboard.component.html',
  styleUrls: ['./user-dashboard.component.css']
})
export class UserDashboardComponent implements OnInit {
  vehicles: any[] = [];
  licenses: any[] = [];

  constructor(
    private mupvozilaService: MupvozilaService,
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.getUserVehicles();
    this.getUserLicenses();
  }

  getUserVehicles(): void {
    const userId = this.authService.getUserId();
    console.log('Fetching vehicles for user:', userId);
    this.mupvozilaService.getVehiclesByUserId(userId).subscribe(
      (data: any) => {
        console.log('Vehicles retrieved:', data);
        this.vehicles = data || []; // Ensure vehicles is an array
      },
      (error) => {
        console.error('Error fetching vehicles:', error);
        this.vehicles = []; // Ensure vehicles is an array in case of error
      }
    );
  }

  getUserLicenses(): void {
    const userId = this.authService.getUserId();
    console.log('Fetching licenses for user ID:', userId);
    this.mupvozilaService.getLicensesByUserId(userId).subscribe(
      (data: any) => {
        console.log('Licenses retrieved:', data);
        this.licenses = data || []; // Ensure licenses is an array
      },
      (error) => {
        console.error('Error fetching licenses:', error);
        this.licenses = []; // Ensure licenses is an array in case of error
      }
    );
  }
}
