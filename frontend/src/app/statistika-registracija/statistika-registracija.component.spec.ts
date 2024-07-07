import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaRegistracijaComponent } from './statistika-registracija.component';

describe('StatistikaRegistracijaComponent', () => {
  let component: StatistikaRegistracijaComponent;
  let fixture: ComponentFixture<StatistikaRegistracijaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaRegistracijaComponent]
    });
    fixture = TestBed.createComponent(StatistikaRegistracijaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
