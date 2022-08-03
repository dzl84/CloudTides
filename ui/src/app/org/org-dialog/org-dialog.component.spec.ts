import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OrgDialogComponent } from './org-dialog.component';

describe('OrgDialogComponent', () => {
  let component: OrgDialogComponent;
  let fixture: ComponentFixture<OrgDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OrgDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OrgDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
