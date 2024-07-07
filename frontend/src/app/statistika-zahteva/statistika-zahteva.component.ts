import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-zahteva',
  templateUrl: './statistika-zahteva.component.html',
  styleUrls: ['./statistika-zahteva.component.css'],
})
export class StatistikaZahtevaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaZahteva();
  }

  fetchStatistikaZahteva(): void {
    this.statistikaService.getStatistikaPravnogZahteva().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika pravnog zahteva:', error);
      }
    );
  }
}
