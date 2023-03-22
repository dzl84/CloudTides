import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
// import { map } from 'rxjs/operators';
import { environment } from '@tide-environments/environment';
import { ORG_PATH } from '@tide-config/path';
import { LoginService } from '../login/login.service';
import toFixed from 'accounting-js/lib/toFixed.js';

@Injectable()
export class OrgService {

  constructor(
    private readonly http: HttpClient,
    private readonly loginService: LoginService,
  ) {
  }

  async getOrgList() {
    const orgList = await this.http.get<ItemOrgSimple[]>(environment.apiPrefix + ORG_PATH, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise();
    const List: ItemOrg[] = [];
    for (const org of orgList) {
      const TempItem: ItemOrg = {
        id: org.id,
        name: org.name,
        currentCPU: org.currentCPU ? org.currentCPU : 0,
        totalCPU: org.totalCPU ? org.totalCPU : 0,
        currentRAM: org.currentRAM ? org.currentRAM : 0,
        totalRAM: org.totalRAM ? org.totalRAM : 0,
        currentDisk: org.currentDisk ? org.currentDisk : 0,
        totalDisk: org.totalDisk ? org.totalDisk : 0,
        percentCPU: !org.totalCPU || !org.currentCPU || org.totalCPU === 0 ? 0: toFixed(100 * org.currentCPU/org.totalCPU, 2),
        percentRAM: !org.totalRAM || !org.currentRAM || org.totalRAM === 0 ? 0: toFixed(100 * org.currentRAM/org.totalRAM, 2),
        percentDisk: !org.totalDisk || !org.currentDisk || org.totalDisk === 0 ? 0: toFixed(100 * org.currentDisk/org.totalDisk, 2),
      }
      List.push(TempItem)
    }
    return List
  }

  addItem(payload: ItemPayload) {
    const body = {
      ...payload,
    };
    return this.http.post<any>(environment.apiPrefix + ORG_PATH, body, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise()
    .then(() => {
      return Promise.resolve();
    }, (errResp) => {
      // console.log(errResp)
      return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
    });
  }

  async removeItem(id: string) {
    await this.http.delete<any>(environment.apiPrefix + ORG_PATH + `/` + id, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
    }, }).toPromise().then(
      () => {
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
  }

}

interface ItemOrgSimple {
  id: number;
  name: string;
  currentCPU: number;
  totalCPU: number;
  currentRAM: number;
  totalRAM: number;
  currentDisk: number;
  totalDisk: number;
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
  percentCPU: number;
  percentRAM: number;
  percentDisk: number;
}


// UI
export interface ItemPayload {
  name: string;
}


export type Item = ItemOrg;
