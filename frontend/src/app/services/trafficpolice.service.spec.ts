import { TestBed } from '@angular/core/testing';

import { TrafficpoliceService } from './trafficpolice.service';

describe('TrafficpoliceService', () => {
  let service: TrafficpoliceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TrafficpoliceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
