import { Component } from '@angular/core';
import { CourtService } from '../court.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-court-home',
  templateUrl: './court-home.component.html',
  styleUrls: ['./court-home.component.css']
})
export class CourtHomeComponent {

  entities: any[] = [];
  warrants: any[] = [];

  constructor(private courtService: CourtService, private authService: AuthService) { }

  ngOnInit(): void {
    this.getUserEntities();
    this.getUserWarrants();
  }

  getUserEntities(){
    let id = this.authService.getUserId();
    this.courtService.getUserEntities(id).subscribe((data: any) => {
      this.entities = data;
    });
  }

  getUserWarrants(){
    let id = this.authService.getUserId();
    this.courtService.getUserWarrants(id).subscribe((data: any) => {
      this.warrants = data;
    });
  }
}
