import { Component, OnInit } from '@angular/core';
import { CourtService } from '../court.service';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-law-etitities',
  templateUrl: './create-law-etitities.component.html',
  styleUrls: ['./create-law-etitities.component.css']
})
export class CreateLawEtititiesComponent implements OnInit{
  selectedForm: string = '';
  entity: any = { title: '', description: '', issueDate: '', dueToDate: '', userId: '',licence: "",category: "" };
  warrant: any = { title: '', description: '', issueDate: '', dueToDate: '', jmbg: '', address: '' };
  hearing: any = { title: '', description: '', scheduledAt: '', duration: 0, legalEntityId: '' };

  constructor(private courtService: CourtService, private authService: AuthService,private router: Router) { }

  ngOnInit(): void {
    if(!this.isCourtWorker()){
      this.router.navigate(['/court-home']);
    }
  }

  isCourtWorker(): boolean{
    if(this.authService.getUserRole() == "courtWorker"){
      return true
    }else{
      return false
    }
  }
  toggleCategoryOptions() {
    if (!this.entity.licence) {
      this.entity.category = ''; // Reset category if license is unchecked
    }
  }

  openForm(formType: string) {
    this.selectedForm = formType;
  }

  formatDatetimeLocal(datetime: string): string {
    // Convert the local datetime to a format acceptable by the backend (ISO 8601 with 'Z' as timezone indicator)
    return new Date(datetime).toISOString();
  }

  submitEntity() {
    const formattedEntity = {
      ...this.entity,
      issueDate: this.formatDatetimeLocal(this.entity.issueDate),
      dueToDate: this.formatDatetimeLocal(this.entity.dueToDate)
    };
    if(this.entity.licence != true){
      this.entity.licence = false;
    }

    this.courtService.createEntity(formattedEntity).subscribe(response => {
      console.log('Legal Entity Created:', response);
      // Reset form
      this.entity = { title: '', description: '', issueDate: '', dueToDate: '', jmbg: '' };
      this.selectedForm = '';
    });
  }

  submitWarrant() {
    const formattedWarrant = {
      ...this.warrant,
      issueDate: this.formatDatetimeLocal(this.warrant.issueDate),
      dueToDate: this.formatDatetimeLocal(this.warrant.dueToDate)
    };

    this.courtService.createWarrant(formattedWarrant).subscribe(response => {
      console.log('Search Warrant Created:', response);
      // Reset form
      this.warrant = { title: '', description: '', issueDate: '', dueToDate: '', userId: '', address: '' };
      this.selectedForm = '';
    });
  }

  submitHearing() {
    const formattedHearing = {
      ...this.hearing,
      scheduledAt: this.formatDatetimeLocal(this.hearing.scheduledAt)
    };

    this.courtService.createHearing(formattedHearing).subscribe(response => {
      console.log('Court Hearing Created:', response);
      // Reset form
      this.hearing = { title: '', description: '', scheduledAt: '', duration: 0, legalEntityId: '' };
      this.selectedForm = '';
    });
  }
}
