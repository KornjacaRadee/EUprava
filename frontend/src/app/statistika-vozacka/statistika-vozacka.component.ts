import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-vozacka',
  templateUrl: './statistika-vozacka.component.html',
  styleUrls: ['./statistika-vozacka.component.css'],
})
export class StatistikaVozackaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaVozackihDozvola();
  }

  fetchStatistikaVozackihDozvola(): void {
    this.statistikaService.getStatistikaVozackihDozvola().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika vozackih dozvola:', error);
      }
    );
  }
}
