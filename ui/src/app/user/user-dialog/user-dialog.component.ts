import { Component, OnInit, Input, Output, EventEmitter, SimpleChanges, ViewChild } from "@angular/core";
import { FormBuilder, Validators, FormGroup } from "@angular/forms";
import { UserService } from "../user.service";
import { TranslateService } from "@ngx-translate/core";
import { defaultRoleType4Site, defaultRoleType4Org, roleTypes, roleTypes4Org } from "@tide-config/cloudPlatform";
import { UserListComponent } from "../user-list/user-list.component";
import { LoginService } from "src/app/login/login.service";
import { CountriesFlagPhone } from 'src/app/@shared/component/countries-flag-phone/countries-flag-phone.component'
@Component({
  selector: "tide-user-dialog",
  templateUrl: "./user-dialog.component.html",
  styleUrls: ["./user-dialog.component.scss"],
})
export class UserDialogComponent implements OnInit {
  constructor(
    private readonly fb: FormBuilder,
    public readonly translate: TranslateService,
    public readonly userService: UserService,
    public readonly userList: UserListComponent,
    public readonly loginService: LoginService
  ) {
    this.userForm = this.fb.group({
      name: ["", [Validators.required]],
      orgName: [
        loginService.inSiteAdminView()? "":localStorage.getItem('orgName'), 
        [Validators.required]
      ],
      // role: ['', [Validators.required]],
      role: ["", [Validators.required]],
      email: ["", [Validators.required, Validators.email]],
      phone: ["", [Validators.required, Validators.pattern("[0-9 ]*")]],
    });
    if (loginService.inSiteAdminView()) {
      this.userForm.controls['orgName'].valueChanges.subscribe((newForm) => {
        if (newForm === "SITE") {
          this.roleTypeList = ["SITE_ADMIN"];
          this.userForm.patchValue({role:"SITE_ADMIN"});
        } else {
          this.roleTypeList = Object.keys(roleTypes4Org);
          this.userForm.patchValue({role:"ORG_ADMIN"});
        }
      });
    }
  

    this.orgmap = userList.orgList;
    this.orgNames = Object.keys(userList.orgList);
    this.roleTypeList = Object.keys(roleTypes4Org);
    this.roleType = this.loginService.inSiteAdminView() ? roleTypes : roleTypes4Org;

  }

  @Input() opened = false;
  @Output() save = new EventEmitter();
  @Output() cancel = new EventEmitter();
  currentCode = 'cn'
  currentAreaCode = '+86'
  userForm: FormGroup;
  orgmap: any;
  orgNames: string[];
  roleTypeList: string[];
  roleType: any;
  defaultRole: string;
  defaultOrg: string;
  validatorFlag = false
  setCurrentCode (value:string) {
    this.currentCode =value
  }
  setCurrentAreaCode (value:string) {
    this.currentAreaCode =value
  }
  validator(e:any) {
    const num = String(e.target.value)
    if (this.currentCode === 'cn' && num.length !== 11) {
      this.validatorFlag =true
    } else {
      this.validatorFlag =false
    }
  }
  readonly vo = {
    serverError: "",
    spinning: false,
  };

  ngOnInit(): void {
  }

  onCancel() {
    this.close();
  }

  // ngOnChanges(changes: any) {
  //   console.log(changes)
  // }

  async onSave() {
    const { value } = this.userForm;
    value.phone = this.currentAreaCode + value.phone
    this.resetModal();
    this.vo.spinning = true;
    await this.userService.addUser(value).then(
      () => {
        this.save.emit("");
        this.close();
        this.vo.spinning = false;
      },
      (error) => {
        this.vo.serverError = error;
        this.vo.spinning = false;
      }
    );
  }

  private close() {
    this.cancel.emit();
  }

  private resetModal() {
    this.vo.serverError = "";
    this.vo.spinning = false;
  }
}
