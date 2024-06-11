import { TestBed } from '@angular/core/testing';

import { MupvozilaService } from './mupvozila.service';

describe('MupvozilaService', () => {
  let service: MupvozilaService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MupvozilaService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
