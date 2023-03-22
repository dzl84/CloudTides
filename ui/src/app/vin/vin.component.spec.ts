import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VinComponent } from './vin.component';

describe('VinComponent', () => {
  let component: VinComponent;
  let fixture: ComponentFixture<VinComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VinComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VinComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
