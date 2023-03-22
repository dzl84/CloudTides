import { Component, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { NOTIFICATION_EXIST_TIME, VENDOR_USAGE_REFRESH_PERIOD } from '@tide-shared/config/const';
import { Observable, of } from 'rxjs';
import { LoginService } from 'src/app/login/login.service';
import { Item, OrgService } from '../org.service';

@Component({
  selector: 'tide-org-list',
  templateUrl: './org-list.component.html',
  styleUrls: ['./org-list.component.scss']
})
export class OrgListComponent implements OnInit {

  constructor(
    public orgService: OrgService,
    public readonly translate: TranslateService,
    public readonly loginService: LoginService,
  ) { }

  readonly vo = {
    alertType: '',
    alertText: '',
  };

  async delete(id: string) {
    await this.orgService.removeItem(id).then(() => {
      this.vo.alertText = `Successfully delete Orgnization with id ${id}`;
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

  async save() {
    await this.refreshList();
  }

  cancel() {
    this.opened = false;

  }

  async ngOnInit() {
    await this.refreshList();
  }


  async refreshList() {
    this.list$ = of(await this.orgService.getOrgList());
    this.refreshInterval = window.setInterval(async () => {
      this.list$ = of(await this.orgService.getOrgList());
    }, VENDOR_USAGE_REFRESH_PERIOD);
  }

  ngOnDestroy(): void {
    window.clearInterval(this.refreshInterval);
  }
}
