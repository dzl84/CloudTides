import { Component, EventEmitter, OnInit, Input, Output } from '@angular/core';
import { Item, ItemUpdateUser, UserService } from '../user.service';
import { TranslateService } from '@ngx-translate/core';
import { NOTIFICATION_EXIST_TIME } from '@tide-config/const';
import { LoginService } from '../../login/login.service';
import { UserListComponent } from '../user-list/user-list.component';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { roleTypes4Org } from '@tide-config/cloudPlatform';
@Component({
  selector: 'tide-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.scss'],
})
export class UserCardComponent implements OnInit {

  constructor(
    private readonly fb: FormBuilder,
    public readonly translate: TranslateService,
    public readonly userService: UserService,
    public readonly loginService: LoginService,
    public readonly userList: UserListComponent,
  ) {
    this.userUpdateForm = this.fb.group({
      id: [],
      name: [userList.updateName, Validators.required],
      orgName: [userList.updateOrg],
      role: [userList.updateRole, Validators.required],
      email: [
        userList.updateEmail, [
          Validators.required,
          Validators.email,
        ]],
      phone: [
        userList.updatePhone, [
          Validators.required,
          Validators.pattern("[0-9 ]{11}"),
        ]],
      //templateID: ['', Validators.required],
    })
    this.roleTypeList = Object.keys(roleTypes4Org);
    this.roleType = roleTypes4Org;
  }

  @Input() userid: number;
  @Input() opened = false;
  @Output() save = new EventEmitter();
  @Output() cancel = new EventEmitter();

  userUpdateForm: FormGroup;
  roleTypeList: string[];
  roleType: any;

  updateSITE() {
    return this.userList.updateOrg === "SITE";
  }
  
  readonly vo = {
    serverError: '',
    spinning: false,
  };

  ngOnInit(): void {
    this.userUpdateForm.get('orgName').disable();
  }

  onCancel() {
    this.close();
  }

  async onSave() {
    const { value } = this.userUpdateForm;
    const payload = this.changeID(value)
    this.resetModal();
    this.vo.spinning = true;
    await this.userService.editItemUser(this.userid, payload).then(() => {
      this.save.emit('');
      this.close();
      this.vo.spinning = false;
    }, (error) => {
      this.vo.serverError = error;
      this.vo.spinning = false;
    });
  }

  private close() {
    this.cancel.emit();
  }
  private resetModal() {
    this.vo.serverError = '';
    this.vo.spinning = false;
  }

  private changeID (payload: ItemUpdateUser) {
    const result : ItemUpdateUser = {
      name: payload.name,
      role: payload.role,
      email: payload.email,
      phone: payload.phone,
    }
    return result;
  }

}

// export interface ItemUpdateUser {
//   name: string;
//   role: string;
//   email: string;
//   phone: string;
// }

