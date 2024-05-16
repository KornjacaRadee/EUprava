import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TrafficPoliceComponent } from './traffic-police.service.component';

describe('TrafficPoliceServiceComponent', () => {
  let component: TrafficPoliceComponent;
  let fixture: ComponentFixture<TrafficPoliceComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TrafficPoliceComponent]
    });
    fixture = TestBed.createComponent(TrafficPoliceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
