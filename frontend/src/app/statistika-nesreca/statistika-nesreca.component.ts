import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-nesreca',
  templateUrl: './statistika-nesreca.component.html',
  styleUrls: ['./statistika-nesreca.component.css'],
})
export class StatistikaNesrecaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}
  ngOnInit(): void {
    this.fetchStatistikaNesreca();
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
}
