import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-registracija',
  templateUrl: './statistika-registracija.component.html',
  styleUrls: ['./statistika-registracija.component.css'],
})
export class StatistikaRegistracijaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaRegistrovanihVozila();
  }

  fetchStatistikaRegistrovanihVozila(): void {
    this.statistikaService.getStatistikaRegistrovanihVozila().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika registrovanih vozila:', error);
      }
    );
  }
}
