import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ComentariosProdComponent } from './comentarios-prod.component';

describe('ComentariosProdComponent', () => {
  let component: ComentariosProdComponent;
  let fixture: ComponentFixture<ComentariosProdComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ComentariosProdComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ComentariosProdComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
