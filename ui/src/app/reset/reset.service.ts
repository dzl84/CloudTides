import { UserInfo } from "./../login/login.service";
import { LoginService } from "src/app/login/login.service";
import { Inject, Injectable } from "@angular/core";
import { HttpClient, HttpErrorResponse } from "@angular/common/http";
import { RESET_API_URL, RESET_PATH, RESET_VERIFY_API_URL } from "@tide-config/path";
import { environment } from "@tide-environments/environment";
import { tap } from "rxjs/operators";
import { DOCUMENT } from "@angular/common";
import { LoginComponent } from "../login/login.component";

@Injectable()
export class ResetService {
  constructor(
    private readonly http: HttpClient,
    // public readonly loginService: LoginService,
    @Inject(DOCUMENT) private readonly document: Document
  ) {}

  reset(username = "", password = "", newPassword = "", verificationCode = "") {
    return this.http
      .post<ResetResult>(environment.apiPrefix + RESET_API_URL, { username, password, newPassword, verificationCode })
      .pipe(tap((val) => {}));
  }

  cloud_reset_code(message = "") {
    console.log("reset_code entered");
    // return this.http.post<VerificationResult>(environment.apiPrefix + RESET_VERIFY_API_URL, {message}).pipe(
    //   tap(() => {}),
    // );
    return this.http
      .post<VerificationResult>(environment.apiPrefix + RESET_VERIFY_API_URL, { message })
      .toPromise()
      .then(
        () => {
          return Promise.resolve();
        },
        (errResp) => {
          return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
        }
      );
  }

  inResetPage() {
    return this.document.location.pathname === "/cloudtides" + RESET_PATH;
  }
}

export interface ResetResult {
  userinfo: {
    username: string;
    password: string;
    pwReset: string;
  };
}

export interface VerificationResult {
  verify: {
    code: string;
  };
}
