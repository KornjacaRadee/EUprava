import { Component } from '@angular/core';
import { MupvozilaService } from '../mupvozila.service';

@Component({
  selector: 'app-mupvozila-communication',
  templateUrl: './mupvozila-communication.component.html',
  styleUrls: ['./mupvozila-communication.component.css']
})
export class MupvozilaCommunicationComponent {
  driverId: string = '';
  accidents: any[] = [];

  constructor(private mupvozilaService: MupvozilaService) { }

  ngOnInit(): void {
  }

  getAccidents() {
    if (this.driverId) {
      this.mupvozilaService.getAccidentsByDriver(this.driverId).subscribe(data => {
        this.accidents = data;
      }, error => {
        console.error('Error fetching accidents:', error);
      });
    }
  }
}