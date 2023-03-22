import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TranslateService } from '@ngx-translate/core';
import { Observable, of } from 'rxjs';
import { OrgListComponent } from '../org-list/org-list.component';
import { OrgService } from '../org.service';

@Component({
  selector: 'tide-org-dialog',
  templateUrl: './org-dialog.component.html',
  styleUrls: ['./org-dialog.component.scss']
})
export class OrgDialogComponent implements OnInit {

  constructor(
    private readonly fb: FormBuilder,
    public readonly translate: TranslateService,
    private readonly orgService: OrgService,
    private readonly  orgList: OrgListComponent,
  ) {
    this.orgForm = this.fb.group({
      name: ['', Validators.required],
    })
  }

  @Input() opened = false;
  @Output() save = new EventEmitter();
  @Output() cancel = new EventEmitter();

  orgForm: FormGroup;

  readonly vo = {
    serverError: '',
    spinning: false,
  };

  ngOnInit(): void {
  }

  onCancel() {
    this.close();
  }

  async onSave() {
    const { value } = this.orgForm;
    this.resetModal();
    this.vo.spinning = true;
    await this.orgService.addItem(value).then(() => {
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

}
