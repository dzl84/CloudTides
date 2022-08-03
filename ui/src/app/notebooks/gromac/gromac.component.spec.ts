import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GromacComponent } from './gromac.component';

describe('GromacComponent', () => {
  let component: GromacComponent;
  let fixture: ComponentFixture<GromacComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GromacComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GromacComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
