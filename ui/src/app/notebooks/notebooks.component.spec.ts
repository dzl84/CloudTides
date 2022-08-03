import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NotebooksComponent } from './notebooks.component';

describe('NotebooksComponent', () => {
  let component: NotebooksComponent;
  let fixture: ComponentFixture<NotebooksComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NotebooksComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NotebooksComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
