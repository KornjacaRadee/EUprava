import { Component, OnInit } from '@angular/core';
import { StatisticsService } from '../statistics.service';

@Component({
  selector: 'app-statistika-vozacka',
  templateUrl: './statistika-vozacka.component.html',
  styleUrls: ['./statistika-vozacka.component.css'],
})
export class StatistikaVozackaComponent implements OnInit {
  statistika: any[] = [];
  procenat: number | null = null;

  constructor(private statistikaService: StatisticsService) {}

  ngOnInit(): void {
    this.fetchStatistikaVozackihDozvola();
    this.fetchStatistikaProcenatVR();
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
  fetchStatistikaProcenatVR(): void {
    this.statistikaService.getStatistikaProcenatVR().subscribe(
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
