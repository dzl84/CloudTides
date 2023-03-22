import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyLandingFooterComponent } from './my-landing-footer.component';

describe('MyLandingFooterComponent', () => {
  let component: MyLandingFooterComponent;
  let fixture: ComponentFixture<MyLandingFooterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyLandingFooterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyLandingFooterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
