import { Component, OnDestroy, OnInit } from '@angular/core';
import { LoginService } from '../login/login.service';
import { Router } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
import { I18nService } from '@tide-shared/service/i18n';
import { EMPTY, Subject } from 'rxjs';
import { catchError, switchMap, tap } from 'rxjs/operators';
import { RegisterService } from '../register/register.service';
import { FormBuilder, Validators, FormGroup, FormControl, ValidatorFn } from '@angular/forms';
import { LOGIN_PATH } from '@tide-config/path';
import { ResetService } from './reset.service';

function passwordMatchValidator(password: string): ValidatorFn {
  return (control: FormControl) => {
    if (!control || !control.parent) {
      return null;
    }
    return control.parent.get(password).value === control.value ? null : { mismatch: true };
  };
}

function passwordUnmatchValidator(password: string): ValidatorFn {
  return (control: FormControl) => {
    if (!control || !control.parent) {
      return null;
    }
    return control.parent.get(password).value === control.value ? {same: true} : null;
  };
}

@Component({
  selector: 'tide-reset',
  templateUrl: './reset.component.html',
  styleUrls: ['./reset.component.scss']
})
export class ResetComponent implements OnInit {

  resetForm: FormGroup;

  constructor(
    public readonly loginService: LoginService,
    public readonly registerService: RegisterService,
    public readonly resetService: ResetService,
    private readonly router: Router,
    public readonly translate: TranslateService,
    public readonly i18nService: I18nService,
    private fb: FormBuilder,
  ) {
    this.resetForm = this.fb.group({
      username: [
        localStorage.getItem("username"), [
          Validators.required,
          Validators.minLength(4),
          Validators.maxLength(12),
        ],
      ],
      password: [
        '', [
          Validators.required,
          Validators.minLength(6),
          Validators.maxLength(16),
        ]],
      newPassword: [
        '', [
          Validators.required,
          Validators.minLength(6),            
          Validators.maxLength(16),
          passwordUnmatchValidator('password'),
        ]],
      confirmPassword: [
        '', [
          Validators.required,
          Validators.minLength(6),
          Validators.maxLength(16),
          passwordMatchValidator('newPassword'),
        ]],
      verificationCode: [
        '', [
          Validators.required,
        ]],
      // email: [
      //   '', [
      //     Validators.required,
      //     Validators.email,
      //   ],
      // ],
      // phone: [
      //   '', [
      //     Validators.required,
      //     Validators.minLength(4),
      //   ],
      // ],
    });
   }

   readonly vo = {
     submitting: false,
     resetError: '',
   };

   private readonly submit$ = new Subject<ResetForm>();

   private readonly submit$$ = this.submit$.asObservable()
   .pipe(
     tap(() => {
       this.vo.submitting = true;
       this.vo.resetError = '';
     }),
     switchMap(({ username, password, newPassword, verificationCode }) => {
       return this.resetService
         .reset(username, password, newPassword, verificationCode)
         .pipe(
           tap(() => {
             this.vo.submitting = false;
           }),
           catchError((error, source) => {
             this.vo.submitting = false;
             this.vo.resetError = error.message;

             return EMPTY as typeof source;
           }),
         );
     }),
   )
   .subscribe(res => {
      this.router.navigate([LOGIN_PATH]);
    })
  ;

  onSubmit() {
    this.submit$.next(this.resetForm.value);
  }

  ngOnInit() {
    // if (this.loginService.session.priority !== 'High') {
    //   // this.document.location.href = '/';
    //   this.router.navigate(['/']);
    // }
  }

  ngOnDestroy() {
    this.submit$$.unsubscribe();
  }

}

interface ResetForm {
  username: string;
  password: string;
  newPassword: string;
  verificationCode: string;
}
