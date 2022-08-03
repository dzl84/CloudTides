import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserComponent } from './user.component';
import { UserService } from './user.service';
import { UserListComponent } from './user-list/user-list.component';
import { UserCardComponent } from './user-card/user-card.component';
import { UserDialogComponent } from './user-dialog/user-dialog.component';

const routes: Routes = [
  {
    path: '',
    component: UserComponent,
    children: [
      {
        path: '',
        component: UserListComponent,
      },
    ],
  },
];

export const declarations = [
  UserComponent,
  UserListComponent,
  UserDialogComponent,
  UserCardComponent,
];

export const providers = [
  UserService,
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class UserRoutingModule {}
