import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
import { HttpEvent, HttpHeaderResponse, HttpResponse, HttpEventType } from '@angular/common/http';
@Component({
  selector: 'tide-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

  constructor(public nd: NotebooksService, private router:Router) {
  }

  ngOnInit(): void {
    this.nd.getApplictionList()
  }
  fileList:FileType[] = []
  progress = 0
  selected= 'selected'
  check = false
  sureDeleteFalg = false
  background = false
  file :FileModel
  filterSearchValue: string = ''
  token = ''
  errorTip = ''
  currentInstanc: AppModel
  get readyupload () {
    if (this.file && this.file.size > 0 && this.file.name !== '') {
      return true
    } else {
      return false
    }
  }
  toJupyter (app: AppModel) {
    const link = app.link.split('?')[0]    
    let form = document.createElement('form')
    form.action='http://'+ link
    form.method='get'
    form.target = '_blank'
    const body = document.documentElement
    body.appendChild(form)
    const input = document.createElement('input')
    input.type = 'hidden'
    input.name = 'token'
    input.value = app.token
    form.appendChild(input)
    setTimeout(() => {
      form.submit()
      form = null
    }, 1000)
  }
  toGromacs (app: AppModel) {
    const link = 'http://' + app.link  
    // let form = document.createElement('form')
    // form.action= link
    // form.method='get'
    // form.target = '_blank'
    // const body = document.documentElement
    // body.appendChild(form)
    //   const input = document.createElement('input')
    //   input.type = 'hidden'
    //   input.name = 'token'
    //   input.value = app.token
    //   form.appendChild(input)
    // setTimeout(() => {
    //   form.submit()
    //   form = null
    // }, 1000)
    window.open(`${link}/token/${app.token}`,'_blank','width=990,height=504,menubar=no,toolbar=no,status=no,scrollbars=yes')
  }
  modifyApp (app: AppModel) {
    this.nd.createInstanceFlag = true
    this.nd.createInstanceTitle = 'HOME.APPLICATION.Modify'
    this.nd.modifiable = false
    console.log('app', app);
    
    this.nd.instanceForm.setValue({
      instanceName: app.instanceName,
      port: app.port,
      appType: app.appType,
      sshHost: app.sshHost,
      sshPassword: app.sshPassword,
      sshPort: app.sshPort,
      sshUser: app.sshUser,
      token: app.token
    })
    this.nd.instanceForm.get('appType').disable({onlySelf: false, emitEvent: false})
  }
  deleteApp (app: AppModel) {
    this.sureDeleteFalg = true
    this.background = true
    this.token = app.token
  }
  getCurrentAppLogs (app: AppModel) {
    this.nd.getAppLogs(app.token)
    this.router.navigate(['/cloudtides/notebooks/log'])
  }
  sure () {
    this.nd.deleteApp(this.token).subscribe(data => {
      localStorage.removeItem(this.token)
      this.sureDeleteFalg = false
      this.background = false
      this.nd.getApplictionList()
    })
  }
  cancel () {
    this.sureDeleteFalg = false
    this.check = false
    this.background = false
    this.token = ''
    this.progress = 0
    this.file = null
    this.errorTip = ''
  }
  checkApp (item: AppModel) {
    this.currentInstanc = item
    this.check = true
    this.background = true
    this.token = item.token
    this.getFileList()
  }
  onDrop (file: FileModel[]) {
    this.file = file[0]
    this.errorTip = ''
    this.nd.uploadData(file[0], this.token).subscribe(
      (data: HttpEvent<FileType> | HttpHeaderResponse | HttpResponse<any>) => {
        if (data.type === HttpEventType.UploadProgress) {
          if (data.total && data.total > 0) {
            this.progress = data.loaded / data.total * 100
          }
        }
        if (data.type === HttpEventType.Response) {
          if (data.status === 200) {
            console.log('上传成功');
            this.getFileList()
          } else {
            this.progress = 0
            console.log('上传失败');
            this.errorTip = '上传失败'
          }
        }
      },
      err => {
        console.log('上传错误',err);
        this.errorTip = err.message
      }
  );
  }
  changeFile(event:Event) {
    this.errorTip = ''
    this.file = (event.target as HTMLInputElement)?.files?.[0]
    this.nd.uploadData((event.target as HTMLInputElement)?.files?.[0],this.token).subscribe(
      (data: HttpEvent<FileType> | HttpHeaderResponse | HttpResponse<any>) => {
        if (data.type === HttpEventType.UploadProgress) {
          if (data.total && data.total > 0) {
            this.progress = data.loaded / data.total * 100
          }
        }
        if (data.type === HttpEventType.Response) {
          if (data.status === 200) {
            console.log('上传成功');
            this.getFileList()
          } else {
            this.progress = 0
            console.log('上传失败');
            this.errorTip = '上传失败'
          }
        }
      },
      err => {
        console.log('上传错误',err);
        this.errorTip = err.message
      }
  );
  }
  getFileList () {
    this.nd.getFileList(this.token).subscribe(
      (data: FileType[]) => {
        this.fileList = data.map(el => {
          el.createtime = Number(el.createtime) * 1000
          return el
        })
      }
    )
  }
  downFile (item:FileType) {
    window.location.href =window.location.origin+ '/' + item.downlink
  }
}
interface FileType {
  createtime: string | number,
  downlink: string,
  filename: string
  filesize: string
}
interface AppModel {
  instanceName: string
  link: string
  token: string
  logo: string
  port: string,
  appType: string,
  sshHost: string,
  sshPassword: string
  sshPort: string
  sshUser: string
  extra: null | Base64
}
interface Base64 {
  base64: string 
}
interface ExtraModel {
  appType: string,
  sshHost: string,
  sshPassword: string
  sshPort: string
  sshUser: string
  cmd: string
}
interface FileModel {
  lastModified: number
  name: string
  size: number
  type: string
}