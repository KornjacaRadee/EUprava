import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaSaslusanjaComponent } from './statistika-saslusanja.component';

describe('StatistikaSaslusanjaComponent', () => {
  let component: StatistikaSaslusanjaComponent;
  let fixture: ComponentFixture<StatistikaSaslusanjaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaSaslusanjaComponent]
    });
    fixture = TestBed.createComponent(StatistikaSaslusanjaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
