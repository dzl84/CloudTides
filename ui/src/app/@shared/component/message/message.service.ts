import { Injectable,Injector } from '@angular/core';
import { ComponentLoaderFactory } from '../loaderFactory'
import { ComponentLoader } from '../components-loader'
import { MessageComponent } from './message.component'
@Injectable({
  providedIn: 'root'
})
export class MessageService {

  constructor(
    private _clf: ComponentLoaderFactory,
    private _injector: Injector,
  ) { 
    this._loader=this._clf.create<MessageComponent>();
  }
  public ref:any
  private _loader:ComponentLoader<MessageComponent>
  private createMessage (t: string, msgContent: string, duration = 10000) {
    this._loader.attch(MessageComponent).to('body')
    const opts = {
      msgType: t,
      msgContent: msgContent
    }
    this.ref = this._loader.create(opts)
    this.ref.changeDetectorRef.markForCheck()
    this.ref.changeDetectorRef.detectChanges()
    setTimeout(() => {
      if (this.ref!=='') {
        try {
          this._loader.remove(this.ref)
        } catch (error) {
          
        }
      }
    }, duration)
  }
  public info(msgContent: string,duration?: number) {
    this.createMessage('info',msgContent,duration);
  }
  public success(msgContent: string,duration?: number) {
    this.createMessage('success',msgContent,duration);
  }
  public error(msgContent: string,duration?: number) {
    this.createMessage('error',msgContent,duration);
  }
  public warning(msgContent: string,duration?: number) {
    this.createMessage('warning',msgContent,duration);
  }
  public close () {
    if (this.ref !== '') {
      try {
        this._loader.remove(this.ref)
      } catch (error) {
        
      }
      this.ref = ''  
    }
  }
}
