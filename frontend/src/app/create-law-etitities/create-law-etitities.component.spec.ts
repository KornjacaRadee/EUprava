import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateLawEtititiesComponent } from './create-law-etitities.component';

describe('CreateLawEtititiesComponent', () => {
  let component: CreateLawEtititiesComponent;
  let fixture: ComponentFixture<CreateLawEtititiesComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateLawEtititiesComponent]
    });
    fixture = TestBed.createComponent(CreateLawEtititiesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
