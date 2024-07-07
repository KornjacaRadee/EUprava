import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MupvozilaCommunicationComponent } from './mupvozila-communication.component';

describe('MupvozilaCommunicationComponent', () => {
  let component: MupvozilaCommunicationComponent;
  let fixture: ComponentFixture<MupvozilaCommunicationComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MupvozilaCommunicationComponent]
    });
    fixture = TestBed.createComponent(MupvozilaCommunicationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
