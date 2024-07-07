import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaNesrecaComponent } from './statistika-nesreca.component';

describe('StatistikaNesrecaComponent', () => {
  let component: StatistikaNesrecaComponent;
  let fixture: ComponentFixture<StatistikaNesrecaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaNesrecaComponent]
    });
    fixture = TestBed.createComponent(StatistikaNesrecaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
