import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-auta',
  templateUrl: './statistika-auta.component.html',
  styleUrls: ['./statistika-auta.component.css'],
})
export class StatistikaAutaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}
  ngOnInit(): void {
    this.fetchStatistikaAuta();
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
}
