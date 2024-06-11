import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MupvozilaComponent } from './mupvozila.component';

describe('MupvozilaComponent', () => {
  let component: MupvozilaComponent;
  let fixture: ComponentFixture<MupvozilaComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MupvozilaComponent]
    });
    fixture = TestBed.createComponent(MupvozilaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
