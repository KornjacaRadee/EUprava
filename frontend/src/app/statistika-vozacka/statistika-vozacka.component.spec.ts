import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaVozackaComponent } from './statistika-vozacka.component';

describe('StatistikaVozackaComponent', () => {
  let component: StatistikaVozackaComponent;
  let fixture: ComponentFixture<StatistikaVozackaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaVozackaComponent]
    });
    fixture = TestBed.createComponent(StatistikaVozackaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
