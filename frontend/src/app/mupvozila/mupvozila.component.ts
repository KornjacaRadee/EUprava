import { Component, OnInit } from '@angular/core';
import { MupvozilaService } from '../mupvozila.service';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-mupvozila',
  templateUrl: './mupvozila.component.html',
  styleUrls: ['./mupvozila.component.css']
})
export class MupvozilaComponent implements OnInit {
  selectedForm: string = '';
  registration: any = { car_id: '', name: '', issuingDate: '', validUntilDate: '' };
  license: any = { type: '', expirationDate: '', issuedBy: '', vehicleId: '' };
  car: any = { owner_id: '', make: '', model: '', year: '', license_plate: '' };
  vehicles: any[] = [];
  constructor(private mupvozilaService: MupvozilaService, private authService: AuthService, private router: Router) { }

  ngOnInit(): void {

    if (!this.isMupWorker()) {
          this.getAllVehicles();
      this.router.navigate(['/mupvozila']);
    }
    
  }


  isMupWorker(): boolean {
    if(this.authService.getUserRole() == "mupWorker"){
      return true
    }else{
      return false
    }
  }

  openForm(formType: string) {
    this.selectedForm = formType;
  }

  formatDatetimeLocal(datetime: string): string {
    // Check if the datetime string is empty or null
    if (!datetime) {
      return ''; // Or any default value you prefer
    }
  
    // Convert the local datetime to a format acceptable by the backend (ISO 8601 with 'Z' as timezone indicator)
    const date = new Date(datetime);
    if (isNaN(date.getTime())) {
      console.error('Invalid datetime:', datetime);
      return ''; // Or handle the error in any other way
    }
    
    return date.toISOString();
  }
  
    submitCar() {
    this.mupvozilaService.createVehicle(this.car).subscribe(response => {
      console.log('Car Created:', response);
      // Reset form
      this.car = { owner_id: '', make: '', model: '', year: '', license_plate: '' };
      this.selectedForm = '';
    });
  }

  submitRegistration() {
    const registration = {
      car_id: this.registration.car_id,
      name: this.authService.getUserId(),
      issuingDate: this.formatDatetimeLocal(this.registration.issuingDate),
      validUntilDate: this.formatDatetimeLocal(this.registration.validUntilDate)
    };

    this.mupvozilaService.registerVehicle(registration).subscribe(response => {
      console.log('Vehicle Registered:', response);
      // Reset form
      this.registration = { car_id: '', name: '', issuingDate: '', validUntilDate: '' };
      this.selectedForm = '';
    });
  }

  submitLicense() {
    const formattedLicense = {
      ...this.license,
      expirationDate: this.formatDatetimeLocal(this.license.expirationDate),
    };

    this.mupvozilaService.createLicense(formattedLicense).subscribe(response => {
      console.log('Vehicle License Created:', response);
      // Reset form
      this.license = { type: '', expirationDate: '', issuedBy: '', vehicleId: '' };
      this.selectedForm = '';
    });

    
  }

  getAllVehicles() {
    this.mupvozilaService.getAllVehicles().subscribe((data: any) => {
      this.vehicles = data;
    });
  }

  
}
