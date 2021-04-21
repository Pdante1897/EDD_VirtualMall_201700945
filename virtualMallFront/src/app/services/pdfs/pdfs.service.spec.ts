import { TestBed } from '@angular/core/testing';

import { PdfsService } from './pdfs.service';

describe('PdfsService', () => {
  let service: PdfsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PdfsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
