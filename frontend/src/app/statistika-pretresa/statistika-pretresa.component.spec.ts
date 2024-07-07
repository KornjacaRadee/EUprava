import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StatistikaPretresaComponent } from './statistika-pretresa.component';

describe('StatistikaPretresaComponent', () => {
  let component: StatistikaPretresaComponent;
  let fixture: ComponentFixture<StatistikaPretresaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StatistikaPretresaComponent]
    });
    fixture = TestBed.createComponent(StatistikaPretresaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
