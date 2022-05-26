import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder
} from '@angular/forms'

@Component({
  selector: 'tide-pool',
  templateUrl: './pool.component.html',
  styleUrls: ['./pool.component.scss']
})
export class PoolComponent implements OnInit {
  poolListLoading = false
  poolList:any[] = []
  pageSizeOptions = [10, 20, 50, 100, 500];
  actionMenuFlag = false
  createEquinixModal = false
  equinixForm!: FormGroup
  titlePool =''
  constructor(private fb: FormBuilder) {
    this.equinixForm = this.fb.group({
      name: [''],
      data_center: [''],
      project: [''],
      administrators: [''],
      server_type: [''],
      usage: [''],
      cost: [''],
      enabled: [true]
    })
  }
  ngOnInit(): void {
  }
  openCreatePoolModal () {}
  public resetEquinixModal () {
    this.createEquinixModal = false
  }

  public createEquinixPoolHandler () {
    this.createEquinixModal = false
  }
}
