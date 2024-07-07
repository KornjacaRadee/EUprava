import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-pretresa',
  templateUrl: './statistika-pretresa.component.html',
  styleUrls: ['./statistika-pretresa.component.css'],
})
export class StatistikaPretresaComponent implements OnInit {
  statistika: any[] = [];

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaPretresa();
  }

  fetchStatistikaPretresa(): void {
    this.statistikaService.getStatistikaNalogaZaPretres().subscribe(
      (data) => {
        this.statistika = Array.isArray(data) ? data : [data]; // Ensure data is an array
      },
      (error) => {
        console.error('Error fetching statistika naloga za pretres:', error);
      }
    );
  }
}
