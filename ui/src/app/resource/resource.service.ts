import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { environment } from '@tide-environments/environment';
import { VCD_URL_PATH, VENDOR_PATH } from '@tide-config/path';
import { LoginService } from '../login/login.service';
import toFixed from 'accounting-js/lib/toFixed.js';

@Injectable()
export class ResourceService {

  constructor(
    private readonly http: HttpClient,
    private readonly loginService: LoginService,
  ) {
  }

  private prefix = `${environment.apiPrefix}/computeResource`;

  // mode variable added by azy 7/12
  mode = 'cloud'

  changeMode(new_mode: string){
    this.mode = new_mode; 
  }

  getMode(){
    return this.mode; 
  }

  async getList() {
    var destination_url : string;
    
    if(this.mode == 'cloud'){
      destination_url = environment.apiPrefix + VCD_URL_PATH
      const list = await this.http.get<Item[]>(destination_url, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      }).toPromise();
      const usage: Item[] = [];
      for (const resource of list) {
        const rawUsage = await this.http.get<ItemUsage>(`${environment.apiPrefix}/usage/${resource.id}`, {
          headers: {
            Authorization: `Bearer ${this.loginService.token}`,
          },
        }).toPromise();
        const resourceItem: Item = {
          hostname: resource.hostname,
          id: resource.id,
          vcdId: resource.vcdId,
          name: rawUsage.name,
          organization: resource.organization,
          vendor: resource.vendor,
          cpu: rawUsage.totalCPU / 1000,
          mem: rawUsage.totalRAM / 1024,
          disk: rawUsage.totalDisk / 1024,
          resType: resource.resType,
          usage: {
            'cpu%': toFixed(rawUsage.percentCPU * 100, 2),
            'mem%': toFixed(rawUsage.percentRAM * 100, 2),
            'disk%': toFixed(rawUsage.percentDisk * 100, 2),
            'cpu': toFixed(rawUsage.percentCPU * rawUsage.totalCPU / 1000, 1),
            'mem': toFixed(rawUsage.percentRAM * rawUsage.totalRAM / 1024, 1),
            'disk': toFixed(rawUsage.percentDisk * rawUsage.totalDisk / 1024, 1),
          },
        };
        usage.push(resourceItem);
      }
      return usage;

    }
    else{
      destination_url = '/api-broker/api/webserver/getHosts';
      
      const list_test = await this.http.get<{message: string, results: ItemDTO_local[]}>(destination_url, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      }).toPromise();
      

      // console.log(list_test);

      const list = list_test.results;
  

      const usage: ItemDTO_local[] = [];
      if (list == null){
        return usage;
      }
      for (const resource of list) {
        // TODO: Modify url here! 
        var rawUsage: ItemUsage = {
          currentCPU : 0,
          totalCPU : 1000,
          currentRAM : 0,
          totalRAM : 1024,
          currentDisk : 0,
          totalDisk : 1024,
          percentCPU : 0,
          percentRAM : 0,
          percentDisk : 0,
          name: ''
        }
        //
        // Get machine usage one by one
        let para = new HttpParams().set("hostname", resource.hostname)

        await this.http.get<ItemUsage>('/api-broker/api/webserver/hostInfo',{
          headers: {
          },
          params: para
        }).toPromise().then((resp) => {
          rawUsage = resp;
        }, (errResp) => {
          console.log('Get item usage failed. Usage set to default. ')
          rawUsage.name = 'Failed to get';
        });
        const resourceItem: ItemDTO_local = {
          hostname: resource.hostname,
          // id: resource.id,
          datacenter: resource.datacenter,
          cluster: resource.cluster,
          ip: resource.ip,
          port: resource.port,
          sshkey: resource.sshkey,
          name: rawUsage.name,
          cpu: rawUsage.totalCPU / 1000,
          mem: rawUsage.totalRAM / 1073741824,
          disk: NaN,
          // resType: resource.resType,
          usage: {
            'cpu%': toFixed(rawUsage.percentCPU , 2),
            'mem%': toFixed(rawUsage.percentRAM , 2),
            'disk%': NaN,
            'cpu': NaN,
            'mem': toFixed(rawUsage.currentRAM/1073741824, 1),
            'disk': NaN,
            // disk usage is not return in current Prometheus
          },
        };
        usage.push(resourceItem);
        
      }
      return usage;

    }
  }

  async getVendorList() {
    const VendorList = await this.http.get<ItemVendor[]>(environment.apiPrefix + VENDOR_PATH, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise();
    const VendorObject : Object = {};
    for (let item of VendorList){
      VendorObject[item.name] = item.url
    }
    return VendorObject
  }

  addItem(payload: ItemPayload) {
    const body = {
      ...payload,
      policy: 0,
    };
    if (this.mode == 'local'){
      body.policy = 1
    }

    var destination_url : string;
    if(this.mode == 'cloud'){
      destination_url = environment.apiPrefix + VCD_URL_PATH

      return this.http.post<any>(destination_url, body, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      }).toPromise().then(() => {
        return Promise.resolve();
      }, (errResp) => {
        console.log('addItem failed')
        return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
      });

    }
    else{
      destination_url = '/api-broker/api/webserver/addHost';

      return this.http.post<any>(destination_url, body, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      }).toPromise().then(() => {
        return Promise.resolve();
      }, (errResp) => {
        console.log('addItem failed')
        return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
      });

    }

  }

  editItem(id: string, payload: ItemPayload) {
    return this.http.put<ItemDTO>(`${this.prefix}/${id}`, payload).pipe(
      map(mapItem),
    );
  }

  async removeItem(id: string) {
    if (this.mode === 'cloud'){
      await this.http.delete<any>(environment.apiPrefix + `/resource/vcd/` + id, {
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
    else{
      await this.http.post<any>('/api-broker/api/webserver/deleteHost',{hostname: id} , {
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

  async contributeResource(id: string): Promise<ContributeResp> {
    let response = null;
    await this.http.put(environment.apiPrefix + `/resource/contribute/${id}`, null, {
      headers: {
      Authorization: `Bearer ${this.loginService.token}`,
    }, }).toPromise().then(
      (resp) => {
        response = resp;
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
    return response;
  }

  async activateResource(id: string): Promise<ActivateResp> {
    let response = null;
    await this.http.put<ActivateResp>(environment.apiPrefix + `/resource/activate/${id}`, null, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise().then(
      (resp) => {
        response = resp;
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
    return response;
  }
}

// Raw
interface ItemUsage {
  currentCPU: number;
  totalCPU: number;
  currentRAM: number;
  totalRAM: number;
  currentDisk: number;
  totalDisk: number;
  percentCPU: number;
  percentRAM: number;
  percentDisk: number;
  name: string;
}

interface ItemDTO {
  hostname: string;
  id: string;
  vcdId: string;
  name: string;
  organization: string;
  vendor: string;
  // unit: GHz
  cpu: number;
  // unit: GB
  mem: number;
  // unit: GB
  disk: number;
  resType: string;
  usage: {
    'cpu%': number;
    'mem%': number;
    'disk%': number;
    'cpu': number;
    'mem': number;
    'disk': number;
  }
}

// New interface added by AZY on 7-18 but not used
// tslint:disable-next-line:class-name
interface ItemDTO_local {
  hostname: string;
  // id: string;
  name: string;
  datacenter: string;
  cluster: string;
  ip: string;
  port: number;
  sshkey: string;
  // unit: GHz
  cpu: number;
  // unit: GB
  mem: number;
  // unit: GB
  disk: number;
  // resType: string;
  usage: {
    'cpu%': number;
    'mem%': number;
    'disk%': number;
    'cpu': number;
    'mem': number;
    'disk': number;
  }
}



interface ContributeResp {
  message: string;
  contributed: boolean;
}

interface ActivateResp {
  message: string;
  activated: boolean;
}

function mapList(raw: ItemDTO[]): Item[] {
  return raw.map(mapItem);
}

function mapItem(raw: ItemDTO): Item {
  return raw;
}

// UI
export interface ItemPayload {
  datacenter: string;
  name: string;
  org: string;
  network: string;
  catalog: string;
  username: string,
  password: string,
  resType: string,
}

interface ItemVendor {
  id: number;
  name: string;
  url: string;
  vendorType: string;
  version: string;
}

export type ItemV = ItemVendor;

export type Item = ItemDTO;

export type Item_local = ItemDTO_local;
