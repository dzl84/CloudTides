import { NgModule } from '@angular/core';

import { SharedModule } from '@tide-shared/shared.module';

import { OrgRoutingModule, declarations, providers } from './org-routing.module';
import { OrgListComponent } from './org-list/org-list.component';

@NgModule({
  declarations: [
    ...declarations,
    OrgListComponent,
  ],
  providers: [
    ...providers,
  ],
  imports: [
    SharedModule,
    OrgRoutingModule,
  ],
})
export class OrgModule {}
