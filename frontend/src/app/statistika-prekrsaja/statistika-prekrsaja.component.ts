import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-prekrsaja',
  templateUrl: './statistika-prekrsaja.component.html',
  styleUrls: ['./statistika-prekrsaja.component.css'],
})
export class StatistikaPrekrsajaComponent implements OnInit {
  statistika: any[] = [];
  procenat: number | null = null;

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaPrekrsaja();
    this.fetchStatistikaProcenatPN();
  }

  fetchStatistikaPrekrsaja(): void {
    this.statistikaService.getStatistikaPrekrsaja().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika prekršaja:', error);
      }
    );
  }

  fetchStatistikaProcenatPN(): void {
    this.statistikaService.getStatistikaProcenatPN().subscribe(
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
