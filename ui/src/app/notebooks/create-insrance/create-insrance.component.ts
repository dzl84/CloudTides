import { Component, OnInit } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
import { Router } from '@angular/router'
@Component({
  selector: 'tide-create-insrance',
  templateUrl: './create-insrance.component.html',
  styleUrls: ['./create-insrance.component.scss']
})
export class CreateInsranceComponent implements OnInit {
  noteBook: NotebooksService
  constructor(private noteBooks: NotebooksService, private router: Router) {
    this.noteBook = this.noteBooks
  }

  ngOnInit(): void {
  }
  currentButton = true
  createInstanceFlag = false
  backCreateInstance () {
    this.router.navigate(['/cloudtides/notebooks/list'])
    // this.noteBook.createInstance = true
  }
  openInstanceModal (type: string) {
    this.noteBook.createInstanceFlag = true
    this.noteBook.currentModel = type
    this.noteBook.gromacsInstanceForm.get('appType').setValue(type)
    this.noteBook.createInstanceTitle = 'HOME.APPLICATION.Create'
    this.noteBook.modifiable= true
    this.noteBook.instanceForm.get('appType').setValue(type)
    this.noteBook.instanceForm.get('appType').disable({
      onlySelf: true
    })
  }
  toggleCurrentButton () {
    this.currentButton = !this.currentButton
  }
}
export interface NotebookModel {
  name: string
  region: string
  zone: string
  environment: string
  machineType: string|number
  bootDisk: string
  subnetwork:string
  externalIp:string
  permission: string
  GPU: string
}
