import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SnackAddFormComponent } from './snack-add-form.component';

describe('SnackAddFormComponent', () => {
  let component: SnackAddFormComponent;
  let fixture: ComponentFixture<SnackAddFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SnackAddFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SnackAddFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
