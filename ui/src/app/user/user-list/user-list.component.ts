import { Component, OnDestroy, OnInit } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Item, UserService } from '../user.service';
import { TranslateService } from '@ngx-translate/core';
import { NOTIFICATION_EXIST_TIME, RESOURCE_USAGE_REFRESH_PERIOD } from '@tide-config/const';
import { LoginService } from 'src/app/login/login.service';
import { roleTypes, roleTypes4Org } from '@tide-config/cloudPlatform';

@Component({
  selector: 'tide-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss'],
})
export class UserListComponent implements OnInit, OnDestroy {

  constructor(
    public userService: UserService,
    public readonly translate: TranslateService,
    public readonly loginService: LoginService,
  ) {
    this.roleTypes = loginService.inSiteAdminView ? roleTypes: roleTypes4Org;

  }

  readonly vo = {
    alertType: '',
    alertText: '',
  };

  async delete(id: string) {
    await this.userService.removeItem(id).then(() => {
      this.vo.alertText = `Successfully delete User with id ${id}`;
      this.vo.alertType = 'success';
    }, (error) => {
      this.vo.alertType = 'danger';
      this.vo.alertText = error;
    }).then(() => {
      this.resetAlert();
    });
    this.refreshList();
  }

  async resetAlert(time?: number) {
    window.setTimeout(() => {
      this.vo.alertText = '';
    }, time || NOTIFICATION_EXIST_TIME);
  }

  
  list$: Observable<Item[]> = of([]);
  opened = false;
  refreshInterval: number;
  // selected: Observable<Item[]> = of([])
  orgList: Object = {};
  UpdateOpened = false;
  UserId = 1;
  updateName: string;
  updateOrg: string;
  updateRole: string;
  updateEmail: string;
  updatePhone: string;
  roleTypes: any;



  async save() {
    await this.refreshList();
  }

  cancel() {
    this.opened = false;
    this.UpdateOpened = false;
  }

  async ngOnInit() {
    await this.refreshList();
  }

  async refreshList() {
    this.list$ = this.loginService.inSiteAdminView() ? of(await this.userService.getUserList()) : of(await this.userService.getUserList4Org(localStorage.getItem('orgName')));
    this.refreshInterval = window.setInterval(async () => {
      this.list$ = this.loginService.inSiteAdminView() ? of(await this.userService.getUserList()) : of(await this.userService.getUserList4Org(localStorage.getItem('orgName')));
    }, RESOURCE_USAGE_REFRESH_PERIOD);
    if (this.loginService.inSiteAdminView()){
      this.orgList = Object(await this.userService.getOrgList());
    }
  }

  async displayDetail(id: number, name: string, org: string, role:string, email: string, phone: string){
    this.UserId = id;
    this.updateName = name;
    this.updateOrg = org;
    this.updateRole = role;
    this.updateEmail = email;
    this.updatePhone = phone;
    this.UpdateOpened = true;
  }

  ngOnDestroy(): void {
    window.clearInterval(this.refreshInterval);
  }

}
