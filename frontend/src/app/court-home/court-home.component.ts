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
  currentUser: any;

  constructor(private courtService: CourtService, private authService: AuthService) { }

  ngOnInit(): void {
    this.getUserEntities();
    this.getUserWarrants();
    this.getUser();
  }

  getUserEntities(){
    let id = this.authService.getUserId();
    this.courtService.getUserEntities(id).subscribe((data: any) => {
      this.entities = data;
    });
  }


  entityRequest() {
    console.log(this.currentUser)
    const entity = {
      title: "Request for Document",
      userJMBG: this.currentUser.jmbg,
      userID: this.currentUser.id
    };

    this.courtService.createRequest(entity).subscribe((response: any) => {
      console.log('Request sent successfully', response);
      this.getUserEntities();
    }, (error: any) => {
      console.error('Error sending request', error);

    });
  }

  getUser(){
    let id = this.authService.getUserId();
    this.authService.getUser(id).subscribe((data: any) => {
      this.currentUser = data;
    });
  }

  getUserWarrants(){
    let id = this.authService.getUserId();
    this.courtService.getUserWarrants(id).subscribe((data: any) => {
      this.warrants = data;
    });
  }
}
