import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
// import { map } from 'rxjs/operators';
import { environment } from "@tide-environments/environment";
import { ORG_PATH, USER_PATH } from "@tide-config/path";
import { LoginService } from "../login/login.service";
import { stringify } from "querystring";

@Injectable()
export class UserService {
  constructor(private readonly http: HttpClient, private readonly loginService: LoginService) {}

  async getUserList() {
    const tempList = await this.http
      .get<ItemUser[]>(environment.apiPrefix + USER_PATH, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      })
      .toPromise();

    //Testing: Print user list url
    console.log("User information got from ", environment.apiPrefix + USER_PATH);
    const List: ItemUser[] = [];
    for (const tem of tempList) {
      const TempItem: ItemUser = {
        id: tem.id,
        name: tem.name,
        orgName: tem.orgName,
        role: tem.role,
        email: tem.email,
        phone: tem.phone,
      };
      List.push(TempItem);
    }
    return List;
  }


  async getUserList4Org(orgName: string) {
    const tempList = await this.http
      .get<ItemUser[]>(environment.apiPrefix + USER_PATH + `/` + orgName, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      })
      .toPromise();

    // Testing print url:
    console.log("UserList for org information got from ", environment.apiPrefix + USER_PATH + `/` + orgName);
    // End testing    

    const List: ItemUser[] = [];
    for (const tem of tempList) {
      const TempItem: ItemUser = {
        id: tem.id,
        name: tem.name,
        orgName: tem.orgName,
        role: tem.role,
        email: tem.email,
        phone: tem.phone,
      };
      List.push(TempItem);
    }
    return List;
  }

  async getOrgList() {
    // console.log("getOrgList");
    const OrgList = await this.http
      .get<ItemOrg[]>(environment.apiPrefix + ORG_PATH, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      })
      .toPromise();

    // Testing print url:
    console.log("User Org information got from ", environment.apiPrefix + ORG_PATH);
    // End testing   

    const OrgObject: Object = {};
    OrgObject["SITE"] = "SITE";
    for (let item of OrgList) {
      OrgObject[item.name] = item.name;
    }
    return OrgObject;
  }

  addUser(payload: ItemAddUser) {
    // console.log(payload);
    const body = {
      ...payload,
    };
    return this.http
      .post<any>(environment.apiPrefix + USER_PATH, body, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      })
      .toPromise()
      .then(
        () => {
          return Promise.resolve();
        },
        (errResp) => {
          // console.log(errResp);
          return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
        }
      );
  }

  editItemUser(id: number, payload: ItemUpdateUser) {
    const body = {
      ...payload,
    };
    return this.http
      .put<any>(environment.apiPrefix + USER_PATH + `/` + id, body, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      })
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

  async removeItem(id: string) {
    await this.http
      .delete<any>(environment.apiPrefix + USER_PATH + `/` + id, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      })
      .toPromise()
      .then(
        () => {
          return Promise.resolve();
        },
        (errResp) => {
          return Promise.reject(`${errResp.message}`);
        }
      );
  }
}

export interface ItemUser {
  id: number;
  name: string;
  orgName: string;
  role: string;
  email: string;
  phone: string;
}
export interface ItemAddUser {
  name: string;
  orgName: string;
  role: string;
  email: string;
  phone: string;
}

export interface ItemUpdateUser {
  name: string;
  role: string;
  email: string;
  phone: string;
}

interface ItemOrg {
  id: number;
  name: string;
  currentCPU: number;
  totalCPU: number;
  currentRAM: number;
  totalRAM: number;
  currentDisk: number;
  totalDisk: number;
}

export type Item = ItemUser;
