import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CountriesFlagPhone } from './countries-flag-phone.component';

describe('CountriesFlagPhone', () => {
  let component: CountriesFlagPhone;
  let fixture: ComponentFixture<CountriesFlagPhone>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [CountriesFlagPhone],
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CountriesFlagPhone);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('country create', () => {
    expect(component).toBeTruthy();
  });
});
