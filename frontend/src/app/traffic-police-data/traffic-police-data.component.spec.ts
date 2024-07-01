import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TrafficPoliceDataComponent } from './traffic-police-data.component';

describe('TrafficPoliceDataComponent', () => {
  let component: TrafficPoliceDataComponent;
  let fixture: ComponentFixture<TrafficPoliceDataComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TrafficPoliceDataComponent]
    });
    fixture = TestBed.createComponent(TrafficPoliceDataComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
