import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaPrekrsajaComponent } from './statistika-prekrsaja.component';

describe('StatistikaPrekrsajaComponent', () => {
  let component: StatistikaPrekrsajaComponent;
  let fixture: ComponentFixture<StatistikaPrekrsajaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaPrekrsajaComponent]
    });
    fixture = TestBed.createComponent(StatistikaPrekrsajaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
