import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CourtHomeComponent } from './court-home.component';

describe('CourtHomeComponent', () => {
  let component: CourtHomeComponent;
  let fixture: ComponentFixture<CourtHomeComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CourtHomeComponent]
    });
    fixture = TestBed.createComponent(CourtHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
