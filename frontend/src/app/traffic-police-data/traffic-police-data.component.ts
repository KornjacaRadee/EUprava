import { Component } from '@angular/core';
import { TrafficPoliceService } from '../services/trafficpolice.service';

@Component({
  selector: 'app-traffic-police-data',
  templateUrl: './traffic-police-data.component.html',
  styleUrls: ['./traffic-police-data.component.css']
})
export class TrafficPoliceDataComponent {
  jmbg: string = '';
  licensePlate: string = '';
  userData: any;
  carData: any;
  licensesData: any;
  allCarsData: any;

  constructor(private trafficPoliceService: TrafficPoliceService) { }

  getUserByJMBG() {
    this.trafficPoliceService.getUserByJMBG(this.jmbg).subscribe(
      data => this.userData = data,
      error => console.error('Error: ', error)
    );
  }

  getCarByLicensePlate() {
    this.trafficPoliceService.getCarByLicensePlate(this.licensePlate).subscribe(data => {
      this.carData = data;
    });
  }
  getLicensesByUserJMBG() {
    this.trafficPoliceService.getLicensesByUserJMBG(this.jmbg).subscribe(data => {
      this.licensesData = data;
    });
  }

  getAllCars() {
    this.trafficPoliceService.getAllCars().subscribe(data => {
      this.allCarsData = data;
    });
  }

  
}
