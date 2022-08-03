import { Component, OnInit, HostListener } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router'
import { NotebooksService, HostType } from './notebooks.service'
import { FormBuilder, Validators } from '@angular/forms'
// import { MessageService } from 'src/app/@shared/component/message/message.service'
@Component({
  selector: 'tide-notebooks',
  templateUrl: './notebooks.component.html',
  styleUrls: ['./notebooks.component.scss']
})
export class NotebooksComponent implements OnInit {

  constructor(public router: Router, public nb: NotebooksService, private fb: FormBuilder) {
    this.noteBook = this.nb
    this.jupyterInstanceForm = this.nb.instanceForm
  }
  ngOnInit(): void {
    this.nb.gromacsInstanceForm.get('appType').disable({
      onlySelf: true
    })
    this.noteBook.getHostNameList().subscribe(
      data => {
        this.hostNameList = data
      }
    )
  }
  noteBook: NotebooksService
  hostNameList: HostType[] = []
  // error tip flag
  loadingFlag:boolean = false
  errorMsg = ''
  // create disable for get
  get disabled () {
    const arr = []
    if (this.noteBook.currentModel === 'jupyter') {
      for (const key in this.jupyterInstanceForm.value) {
        if (key !== 'token') {
          arr.push(this.jupyterInstanceForm.value[key])
        }
      }
    } else if (this.noteBook.currentModel !== 'jupyter') {
      for (const key in this.nb.gromacsInstanceForm.value) {
        if (key !== 'token') {
          arr.push(this.nb.gromacsInstanceForm.value[key])
        }
      }
    }
    const rsg = arr.every(el => el !== '')    
    return !rsg
  }
  jupyterInstanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    port: ['', Validators.required],
    appType: ['jupyter'],
    // sshHost: [{}]
    // token: ['']
  })
  cancel () {    
    this.noteBook.createInstanceFlag = false
    this.jupyterInstanceForm.setValue({
      instanceName: '',
      port: '',
      appType: '',
      // sshHost: ''
      // token: ''
    })
    this.nb.gromacsInstanceForm.setValue({
      instanceName: '',
      // sshHost: '',
      appType: [''],
      })
    this.errorMsg = ''
    this.loadingFlag = false

  }
  create (form) {
    this.errorMsg = ''
    this.loadingFlag = true
    const data = form.getRawValue()
    // const host = this.hostNameList.find(el => el.address === data.sshHost)
    const host = this.hostNameList[0]
    data.sshPassword = host.sshPass
    data.sshPort = host.sshPort
    data.sshUser = host.sshUser
    data.sshHost = host.address
    if (data.appType === 'jupyter') {
      const str = 'HOME.APPLICATION.Create'
      data.sshPort = +data.sshPort
      if (this.nb.createInstanceTitle === str) {
        // delete data.token
        this.nb.createNewApp(data).subscribe((data:CreateData) => {
          this.loadingFlag = false
          this.noteBook.createInstanceFlag = false
          this.jupyterInstanceForm.setValue({
            instanceName: '',
            port: '',
            appType: '',
            // sshHost: ''
          })
          // this.$msg.success('serverMessage.create200', 2000)
          window.open('http://' + data.link, "_blank")
        },
        err => {
          console.log('err', err);
          
          this.errorMsg = err.statusText
          this.loadingFlag = false
        })
      } else {
        this.nb.modifyApp(data).subscribe(data => {
          this.loadingFlag = false
          this.noteBook.createInstanceFlag = false
          this.nb.getApplictionList()
        },
        err => {
          this.errorMsg = err.msg
          this.loadingFlag = false
        })
      }
    } else if (data.appType !== 'jupyter') {
      data.sshPort = +data.sshPort
      this.nb.createNewApp(data).subscribe((data:CreateData) => {
        this.loadingFlag = false
        this.noteBook.createInstanceFlag = false
        this.nb.gromacsInstanceForm.setValue({
          instanceName: '',
          appType: '',
          // sshHost: ''
        })
        // this.$msg.success('serverMessage.create200', 2000)
        this.router.navigateByUrl('/cloudtides/notebooks/list')
      },
      err => {
        console.log('err', err);
        
        this.errorMsg = err.statusText
        this.loadingFlag = false
      })
    }
  }
}
interface CreateData {
  link: string
  token: string
}
