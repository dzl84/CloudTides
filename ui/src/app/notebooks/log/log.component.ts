import { Component, OnInit,OnDestroy } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
import { Router } from '@angular/router'
@Component({
  selector: 'tide-log',
  templateUrl: './log.component.html',
  styleUrls: ['./log.component.scss']
})
export class LogComponent implements OnInit,OnDestroy {

  constructor(public nd: NotebooksService,private router:Router) {    
  }
  ngOnDestroy(): void {
    this.nd.closeWs()
  }
  ngOnInit(): void {
    if (!this.nd.currentToken) {
      this.router.navigate(['/cloudtides/notebooks/list'])
    }
  }
  back () {
    this.nd.closeWs()
    this.router.navigate(['/cloudtides/notebooks/list'])
  }
}
