import { NgModule } from '@angular/core';

import { SharedModule } from '@tide-shared/shared.module';

import { UserRoutingModule, declarations, providers } from './user-routing.module';
import { UserListComponent } from './user-list/user-list.component';
@NgModule({
  declarations: [
    ...declarations,
    UserListComponent,
  ],
  providers: [
    ...providers,
  ],
  imports: [
    SharedModule,
    UserRoutingModule,
  ],
})
export class UserModule {}
