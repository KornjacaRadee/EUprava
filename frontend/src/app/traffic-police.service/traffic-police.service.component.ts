// traffic-police.component.ts

import { Component, OnInit } from '@angular/core';
import { TrafficPoliceService } from '../services/trafficpolice.service';

@Component({
  selector: 'app-traffic-police',
  templateUrl: '../traffic-police.service/traffic-police.service.component.html',
  styleUrls: ['../traffic-police.service/traffic-police.service.component.css']
})
export class TrafficPoliceComponent implements OnInit {
  prekrsaji: any[] = [];
  nesrece: any[] = [];
  prekrsajData: any = { lokacija: '', vozac: '', vozilo: '', opis: '' };
  nesrecaData: any = { lokacija: '', vozac: '', vozilo: '', opis: '' };

  constructor(private trafficPoliceService: TrafficPoliceService) { }

  ngOnInit(): void {
    this.getPrekrsaji();
    this.getNesrece();
  }

  getPrekrsaji(): void {
    this.trafficPoliceService.getPrekrsaji().subscribe(prekrsaji => {
      this.prekrsaji = prekrsaji;
      console.log('Prekršaji fetched:', prekrsaji);
    });
  }

  createPrekrsaj(): void {
    const formattedPrekrsaj = {
      ...this.prekrsajData,
      // Add any necessary formatting here if needed, e.g., date formatting
    };

    this.trafficPoliceService.createPrekrsaj(formattedPrekrsaj).subscribe(response => {
      console.log('Prekršaj Created:', response);
      // Refresh the list of Prekršaji
      this.getPrekrsaji();
      // Reset form fields
      this.prekrsajData = { lokacija: '', vozac: '', vozilo: '', opis: '' };
    });
  }

  deletePrekrsaj(prekrsajId: string): void {
    this.trafficPoliceService.deletePrekrsaj(prekrsajId).subscribe(() => {
      console.log('Prekršaj Deleted:', prekrsajId);
      this.getPrekrsaji();
    });
  }
  

  getNesrece(): void {
    this.trafficPoliceService.getNesrece().subscribe(nesrece => {
      this.nesrece = nesrece;
      console.log('Nesreće fetched:', nesrece);
    });
  }

  createNesreca(): void {
    const formattedNesreca = {
      ...this.nesrecaData,
      // Add any necessary formatting here if needed, e.g., date formatting
    };

    this.trafficPoliceService.createNesreca(formattedNesreca).subscribe(response => {
      console.log('Nesreca Created:', response);
      // Refresh the list of Nesreće
      this.getNesrece();
      // Reset form fields
      this.nesrecaData = { lokacija: '', vozac: '', vozilo: '', opis: '' };
    });
  }

  deleteNesreca(nesrecaId: string): void {
    this.trafficPoliceService.deleteNesreca(nesrecaId).subscribe(() => {
      console.log('Nesreca Deleted:', nesrecaId);
      this.getNesrece();
    });
  }
}