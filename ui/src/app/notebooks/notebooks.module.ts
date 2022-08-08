import { NgModule } from '@angular/core';

import { SharedModule } from '@tide-shared/shared.module';
import { NotebooksService } from './notebooks.service'
import { NotebooksRoutingModule, declarations, providers } from './notebooks-routing.module';
import { NotebooksComponent } from './notebooks.component';
import { ListComponent } from './list/list.component'
import { RouteReuseStrategyService } from '@tide-shared/service/route-reuse-stratety.service'
import { RouteReuseStrategy } from '@angular/router';
import { CreateInsranceComponent } from './create-insrance/create-insrance.component';
import { LogComponent } from './log/log.component';
import { HistoryComponent } from './history/history.component';
import { GromacComponent } from './gromac/gromac.component';
@NgModule({
  declarations: [
    // ...declarations
    NotebooksComponent,
    ListComponent,
    CreateInsranceComponent,
    LogComponent,
    HistoryComponent,
    GromacComponent
  ],
  providers: [
    NotebooksService,
    { provide: RouteReuseStrategy, useClass: RouteReuseStrategyService },
    ...providers,
  ],
  imports: [
    SharedModule,
    NotebooksRoutingModule
  ],
})
export class NotebooksModule {}
