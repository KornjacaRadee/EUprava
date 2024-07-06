import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaZahtevaComponent } from './statistika-zahteva.component';

describe('StatistikaZahtevaComponent', () => {
  let component: StatistikaZahtevaComponent;
  let fixture: ComponentFixture<StatistikaZahtevaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaZahtevaComponent]
    });
    fixture = TestBed.createComponent(StatistikaZahtevaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
