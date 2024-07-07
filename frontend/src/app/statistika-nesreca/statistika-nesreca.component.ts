import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-nesreca',
  templateUrl: './statistika-nesreca.component.html',
  styleUrls: ['./statistika-nesreca.component.css'],
})
export class StatistikaNesrecaComponent implements OnInit {
  statistika: any[] = [];
  procenat: number | null = null;

  constructor(private statistikaService: StatisticsService) {}
  ngOnInit(): void {
    this.fetchStatistikaNesreca();
    this.fetchStatistikaProcenatNP();
  }

  fetchStatistikaNesreca(): void {
    this.statistikaService.getStatistikaNesreca().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika nesreca:', error);
      }
    );
  }

  fetchStatistikaProcenatNP(): void {
    this.statistikaService.getStatistikaProcenatNP().subscribe(
      (data) => {
        console.log('Statistika procenat NP:', data);
        this.procenat = data.percentage; // Pretpostavljamo da se procenat nalazi u polju 'percentage' objekta
      },
      (error) => {
        console.error('Error fetching statistika procenata:', error);
        this.procenat = null; // Postavljamo na null u slučaju greške
      }
    );
  }
}
