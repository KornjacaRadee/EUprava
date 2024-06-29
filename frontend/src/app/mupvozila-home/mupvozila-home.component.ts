import { Component, OnInit } from '@angular/core';
import { MupvozilaService } from '../mupvozila.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-mupvozila-home',
  templateUrl: './mupvozila-home.component.html',
  styleUrls: ['./mupvozila-home.component.css']
})
export class MupvozilaHomeComponent implements OnInit {
  selectedForm: string = '';
  vehicles: any[] = [];
  registrations: any[] = [];
  licenses: any[] = [];
  carIdToLicensePlateMap: { [key: string]: string } = {};
  newLicense: any = {
    user_jmbg: '',
    category: '',
    issuing_date: '',
    valid_until_date: '',
    address: '',
    points: 0,
    is_valid: true
  };
  selectedCar: any = null;
  editingRegistration: any = null;
  editingLicense: any = null;

  constructor(private mupvozilaService: MupvozilaService, private authService: AuthService) { }

  ngOnInit(): void {
   // this.getCurrentUserIdVehicles();
    this.getUserRegistrations();
    this.getUserLicenses();
    this.getAllVehicles();
    this.getAllRegistrations();
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
      this.createCarIdToLicensePlateMap(); // Create map after fetching vehicles
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

  getAllRegistrations() {
    this.mupvozilaService.getAllRegistrations().subscribe((data: any) => {
      this.registrations = data;
    });
  }

  deleteRegistration(registrationId: string) {
    this.mupvozilaService.deleteRegistration(registrationId).subscribe(() => {
      this.registrations = this.registrations.filter(registration => registration.id !== registrationId);
    });
  }

  deleteLicense(licenseId: string) {
    this.mupvozilaService.deleteLicense(licenseId).subscribe(() => {
      this.licenses = this.licenses.filter(license => license.id !== licenseId);
    });
  }

  createCarIdToLicensePlateMap() {
    this.carIdToLicensePlateMap = {};
    for (const vehicle of this.vehicles) {
      this.carIdToLicensePlateMap[vehicle.id] = vehicle.license_plate;
    }
  }

  getCarLicensePlate(carId: string): string {
    return this.carIdToLicensePlateMap[carId] || carId;
  }

  issueLicense() {
    this.mupvozilaService.issueLicense(this.newLicense).subscribe((response: any) => {
      console.log('License issued:', response);
      this.newLicense = {
        user_jmbg: '',
        category: '',
        issuing_date: '',
        valid_until_date: '',
        address: '',
        points: 0,
        is_valid: true
      };
      this.getUserLicenses(); // Refresh the licenses list
    });
  }


  selectCarForEdit(car: any) {
    this.selectedCar = { ...car };
  }

  updateCar() {
    if (this.selectedCar) {
      this.mupvozilaService.updateCar(this.selectedCar).subscribe(response => {
        console.log('Car updated:', response);
        this.selectedCar = null;
        this.getAllVehicles(); // Refresh the vehicles list
      });
    }
  }


  editRegistration(registration: any) {
    this.editingRegistration = { ...registration };
  }

  formatDatetimeLocal(datetime: string): string {
    if (!datetime) {
      return '';
    }
    const date = new Date(datetime);
    if (isNaN(date.getTime())) {
      console.error('Invalid datetime:', datetime);
      return '';
    }
    return date.toISOString();
  }

  updateRegistration() {
    if (this.editingRegistration) {
      const formattedRegistration = {
        ...this.editingRegistration,
        issuing_date: this.formatDatetimeLocal(this.editingRegistration.issuing_date),
        valid_until_date: this.formatDatetimeLocal(this.editingRegistration.valid_until_date)
      };

      this.mupvozilaService.updateRegistration(formattedRegistration).subscribe(response => {
        console.log('Registration updated:', response);
        this.editingRegistration = null;
        this.getAllRegistrations(); // Refresh the registrations list
      });
    }
  }

  editLicense(license: any) {
    this.editingLicense = { ...license };
  }

  updateLicense() {
    if (this.editingLicense) {
      const formattedLicense = {
        ...this.editingLicense,
        issuing_date: this.formatDatetimeLocal(this.editingLicense.issuing_date),
        valid_until_date: this.formatDatetimeLocal(this.editingLicense.valid_until_date)
      };

      this.mupvozilaService.updateLicense(formattedLicense).subscribe(response => {
        console.log('License updated:', response);
        this.editingLicense = null;
        this.getUserLicenses(); // Refresh the licenses list
      });
    }
  }
}