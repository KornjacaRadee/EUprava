import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaAutaComponent } from './statistika-auta.component';

describe('StatistikaAutaComponent', () => {
  let component: StatistikaAutaComponent;
  let fixture: ComponentFixture<StatistikaAutaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaAutaComponent]
    });
    fixture = TestBed.createComponent(StatistikaAutaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
