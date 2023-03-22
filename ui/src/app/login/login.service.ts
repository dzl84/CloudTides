import { ResetService } from './../reset/reset.service';
import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { DOCUMENT } from '@angular/common';
import { BehaviorSubject } from 'rxjs';
import { tap } from 'rxjs/operators';
import { environment } from '@tide-environments/environment';
import { LOGIN_API_URL, LOGIN_PATH, PROFILE_API_URL } from '@tide-config/path';
import { LOCAL_STORAGE_KEY } from '@tide-config/const';
import { Router } from '@angular/router';
import { RegisterService } from '../register/register.service';

@Injectable()
export class LoginService {

  constructor(
    private readonly http: HttpClient,
    private readonly router: Router,
    private readonly registerService: RegisterService,
    private readonly resetService: ResetService,
    @Inject(DOCUMENT) private readonly document: Document,
  ) {
    this.loginNavigate();
  }

  readonly session$ = new BehaviorSubject<UserInfo>({} as any);

  login(
    username = '',
    password = '',
  ) {
    return this.http.post<ServerUserInfo>(environment.apiPrefix + LOGIN_API_URL, { username, password }).pipe(
      tap(serverUserInfo => {
        this.storeToken(serverUserInfo.token);
        this.session$.next({ ...serverUserInfo.userInfo});
        this.storeRole(this.session.role);
        this.storeOrgName(this.session.orgName);
        this.storePwReset(this.session.pwReset);
        this.storeUsername(this.session.username);
        this.storeAuthority(this.session.authority)
        // this.storePwReset(this.session.pwReset);
        // console.log(localStorage.getItem("role"));
        // console.log(this.session.pwReset);
        
        // if (this.session.pwReset == true) {
        //   console.log("inside");
        //   console.log(this.session.pwReset);
        // }
        // console.log("outside");
      }),
    );
  }

  async loginNavigate() {
    if (this.hasLoggedIn) {
      
      // if (!this.session.pwReset) {
      //   this.router.navigate(['cloudtides/reset'])
      // }
      this.current().subscribe(
        () => {},
        async error => {
          await this.logout();
        });
    } else {
      if (!this.inLoginPage() && !this.registerService.inRegisterPage() && !this.resetService.inResetPage()) {
        await this.logout();
      }
      // else if (!this.inLoginPage() && !this.registerService.inRegisterPage() && !this.session.pwReset) {
      //   this.router.navigate(['cloudtides/reset'])
      // }
    }
  }

  current() {
    return this.http.get<any>(environment.apiPrefix + PROFILE_API_URL, {
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
    }).pipe(
      tap(returnMessage => {
        const userInfo = returnMessage.results;
        this.session$.next(userInfo);
      }),
    );
  }

  async logout() {
    console.log('logout')
    this.removeToken();
    this.removeInfo();

    // await this.router.navigate(['home']);
    await this.router.navigate(['login'])
    .then(() => {
      console.log('home navigate is called')
      // window.location.reload();
    });
    // await location.href('home')
  }

  async cloudtides_logout() {
    console.log('cloudtides logout')
    this.removeToken();
    // await this.router.navigate(['home']);
    await this.router.navigate(['login'])
    .then(() => {
      console.log('home navigate is called')
      window.location.reload();
    });
    // await location.href('home')
  }

  storeToken(token: string) {
    localStorage.setItem(LOCAL_STORAGE_KEY.TOKEN, token);
  }

  storeRole(role: string) {
    localStorage.setItem('role', role);
  }

  storeOrgName(org: string){
    localStorage.setItem('orgName', org);
  }

  storeUsername(username: string){
    localStorage.setItem('username', username);
  }


  storePwReset(pwReset: string){
    localStorage.setItem('pwReset', pwReset);
  }

  storeAuthority(authority: string){
    localStorage.setItem('authority', authority);
  }

  removeInfo() {
    localStorage.removeItem('role');
    localStorage.removeItem('orgName');
    localStorage.removeItem('pwReset');
    localStorage.removeItem('username');
  }

  removeToken() {
    localStorage.removeItem(LOCAL_STORAGE_KEY.TOKEN);
  }

  get session() {
    return this.session$.value;
  }

  get hasLoggedIn() {
    return localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN) !== null;
  }

  get token() {
    return localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN);
  }

  inLoginPage() {
    return this.document.location.pathname === LOGIN_PATH;
  }

  inAdminView() {
    // return true;
    return this.session.priority === 'High';
  }

  inSiteAdminView() {
    return localStorage.getItem('role') === 'SITE_ADMIN';
  }

  inOrgAdminView() {
    return localStorage.getItem('role') === 'ORG_ADMIN';
  }

  inUserView() {
    return localStorage.getItem('role') === 'USER';
  }

}

export interface UserInfo {
  username: string;
  priority: string;
  firstName: string;
  lastName: string;
  country: string;
  city: string;
  companyName: string;
  position: boolean;
  email: string,
  phone: string,
  role: string,
  pwReset: string,
  orgName: string,
  authority: string
}

export interface ServerUserInfo {
  token: string;
  userInfo: UserInfo;
}

