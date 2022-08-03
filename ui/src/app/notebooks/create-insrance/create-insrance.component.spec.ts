import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateInsranceComponent } from './create-insrance.component';

describe('CreateInsranceComponent', () => {
  let component: CreateInsranceComponent;
  let fixture: ComponentFixture<CreateInsranceComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateInsranceComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateInsranceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
