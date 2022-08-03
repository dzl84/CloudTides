import { Component, Input, OnInit, ViewChild, ElementRef } from '@angular/core';
import { ClarityIcons, successStandardIcon, exclamationCircleIcon,exclamationTriangleIcon,infoCircleIcon  } from '@cds/core/icon';
// import { AppService } from '../../app.service'
ClarityIcons.addIcons(successStandardIcon);
ClarityIcons.addIcons(exclamationCircleIcon);
ClarityIcons.addIcons(exclamationTriangleIcon);
ClarityIcons.addIcons(infoCircleIcon);
@Component({
  selector: 'app-message',
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.css']
})
export class MessageComponent implements OnInit {

  constructor() {
  }
  public classType: string[] = []
  public classCloseType: string[] = []
  public msgContent = ''
  ngOnInit(): void {
    this.classType = ['upc-message-' + this.msgType]
    this.classCloseType = ['close-' + this.msgType]
  }
  @Input() msgType: 'success' | 'info' | 'warning' | 'hide' | 'error'='info'
  @ViewChild('msg') msg!: ElementRef;
  public distory () {
    this.msg.nativeElement.style.display = 'none'
  }
}
