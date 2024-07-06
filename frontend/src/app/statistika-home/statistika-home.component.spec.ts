import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaHomeComponent } from './statistika-home.component';

describe('StatistikaHomeComponent', () => {
  let component: StatistikaHomeComponent;
  let fixture: ComponentFixture<StatistikaHomeComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaHomeComponent]
    });
    fixture = TestBed.createComponent(StatistikaHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
