import { ModuleWithProviders, NgModule } from '@angular/core';

import { BaseModule } from './base.module';
import { sharedComponents } from './component';
import { sharedPipes } from './pipe';
import { TranslateModule } from '@ngx-translate/core';
import { WebSocketService } from './service/web-socket.service'
@NgModule({
  imports: [
    BaseModule
  ],
  exports: [
    BaseModule,
    TranslateModule,
    ...sharedComponents,
    ...sharedPipes,
  ],
  declarations: [
    ...sharedComponents,
    ...sharedPipes,
  ],
})
export class SharedModule {

  static forRoot(): ModuleWithProviders {
    return {
      ngModule: SharedModule,
      providers: [
        WebSocketService
      ],
    };
  }

  static forChild(): ModuleWithProviders {
    return {
      ngModule: SharedModule,
      providers: [],
    };
  }
}
