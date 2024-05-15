import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MupvozilaHomeComponent } from './mupvozila-home.component';

describe('MupvozilaHomeComponent', () => {
  let component: MupvozilaHomeComponent;
  let fixture: ComponentFixture<MupvozilaHomeComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MupvozilaHomeComponent]
    });
    fixture = TestBed.createComponent(MupvozilaHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
