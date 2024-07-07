import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-registracija',
  templateUrl: './statistika-registracija.component.html',
  styleUrls: ['./statistika-registracija.component.css'],
})
export class StatistikaRegistracijaComponent implements OnInit {
  statistika: any[] = [];
  procenat: number | null = null;

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaRegistrovanihVozila();
    this.fetchStatistikaProcenatRV();
  }

  fetchStatistikaRegistrovanihVozila(): void {
    this.statistikaService.getStatistikaRegistrovanihVozila().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data];
      },
      (error) => {
        console.error('Error fetching statistika registrovanih vozila:', error);
      }
    );
  }
  fetchStatistikaProcenatRV(): void {
    this.statistikaService.getStatistikaProcenatRV().subscribe(
      (data) => {
        console.log('Statistika procenat VR:', data);
        this.procenat = data.percentage; // Pretpostavljamo da se procenat nalazi u polju 'percentage' objekta
      },
      (error) => {
        console.error('Error fetching statistika procenata:', error);
        this.procenat = null; // Postavljamo na null u slučaju greške
      }
    );
  }
}
