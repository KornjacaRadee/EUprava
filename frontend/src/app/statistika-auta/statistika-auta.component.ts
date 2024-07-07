import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-auta',
  templateUrl: './statistika-auta.component.html',
  styleUrls: ['./statistika-auta.component.css'],
})
export class StatistikaAutaComponent implements OnInit {
  statistika: any[] = [];
  procenat: number | null = null;

  constructor(private statistikaService: StatisticsService) {}
  ngOnInit(): void {
    this.fetchStatistikaAuta();
    this.fetchStatistikaProcenatRA();
  }
  fetchStatistikaAuta(): void {
    this.statistikaService.getStatistikaAuta().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika auta:', error);
      }
    );
  }

  fetchStatistikaProcenatRA(): void {
    this.statistikaService.getStatistikaProcenatRA().subscribe(
      (data) => {
        console.log('Statistika procenat RA:', data);
        this.procenat = data.percentage;
      },
      (error) => {
        console.error('Error fetching statistika procenata:', error);
        this.procenat = null; // Postavljamo na null u slučaju greške
      }
    );
  }
}
