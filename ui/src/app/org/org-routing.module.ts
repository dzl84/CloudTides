import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { OrgComponent } from './org.component';
import { OrgService } from './org.service';
import { OrgListComponent } from './org-list/org-list.component';
import { OrgDialogComponent } from './org-dialog/org-dialog.component';



const routes: Routes = [
  {
    path: '',
    component: OrgComponent,
    children: [
      {
        path: '',
        component: OrgListComponent,
      },
    ],
  },
];

export const declarations = [
  OrgComponent,
  OrgListComponent,
  OrgDialogComponent,
];

export const providers = [
  OrgService,
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class OrgRoutingModule {}
