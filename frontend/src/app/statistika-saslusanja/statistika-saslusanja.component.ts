import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-saslusanja',
  templateUrl: './statistika-saslusanja.component.html',
  styleUrls: ['./statistika-saslusanja.component.css'],
})
export class StatistikaSaslusanjaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaSaslusanja();
  }

  fetchStatistikaSaslusanja(): void {
    this.statistikaService.getStatistikaSaslusanja().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika saslusanja:', error);
      }
    );
  }
}
