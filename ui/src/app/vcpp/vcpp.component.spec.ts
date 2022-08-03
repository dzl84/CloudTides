import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VcppComponent } from './vcpp.component';

describe('VcppComponent', () => {
  let component: VcppComponent;
  let fixture: ComponentFixture<VcppComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VcppComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VcppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
