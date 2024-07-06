import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-prekrsaja',
  templateUrl: './statistika-prekrsaja.component.html',
  styleUrls: ['./statistika-prekrsaja.component.css'],
})
export class StatistikaPrekrsajaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaPrekrsaja();
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
}
