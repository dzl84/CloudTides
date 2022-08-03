import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { LoginComponent } from './login/login.component';
import { LoginService } from './login/login.service';
import { RegisterComponent } from './register/register.component';
import { ResetComponent } from './reset/reset.component';

import {
  LOGIN_PATH_NAME,
  HOME_PATH_NAME,
  RESOURCE_PATH_NAME,
  TEMPLATE_PATH_NAME,
  POLICY_PATH_NAME,
  REGISTER_PATH_NAME,
  RESET_PATH_NAME,
  VENDOR_PATH_NAME,
  VAPP_PATH_NAME,
  ORG_PATH_NAME,
  USER_PATH_NAME,
  NOTEBOOKS_PATH_NAME
} from '@tide-config/path';


import { AuthGuard } from '@tide-guard/auth.guard';
import { RegisterService } from './register/register.service';
import { ResetService } from './reset/reset.service';
import { LandingComponent } from './landing/landing.component';
import { VinComponent } from './vin/vin.component';
import { VcppComponent } from './vcpp/vcpp.component';

const routes: Routes = [
  { 
    path: '', 
    pathMatch: 'full',
    redirectTo: LOGIN_PATH_NAME
  },
  // {
  //   path: 'home',
  //   component: LandingComponent
  // },
  {
    path: 'vin',
    component: VinComponent
  },
  {
    path: 'vcpp',
    component: VcppComponent
  },
  {
    path: LOGIN_PATH_NAME,
    component: LoginComponent,
    data: {
      anonymous: true,
    } as RouterData,
  },
  {
    path: 'cloudtides',
    canActivateChild: [AuthGuard],
    children: [
      {
        path: '',
        pathMatch: 'full',
        // redirectTo: HOME_PATH_NAME,
        redirectTo: RESOURCE_PATH_NAME
      },
      // {
      //   path: LOGIN_PATH_NAME,
      //   component: LoginComponent,
      //   data: {
      //     anonymous: true,
      //   } as RouterData,
      // },
      {
        path: REGISTER_PATH_NAME,
        component: RegisterComponent,
        data: {
          anonymous: true,
        } as RouterData,
      },
      {
        path: RESET_PATH_NAME,
        component: ResetComponent,
        data: {
          anonymous: true,
        } as RouterData,
      },
      // {
      //   path: HOME_PATH_NAME,
      //   loadChildren: () => import('./home/home.module').then(m => m.HomeModule),
      // },
      {
        path: VENDOR_PATH_NAME,
        loadChildren: () => import('./vendor/vendor.module').then(m => m.VendorModule)
      },
      {
        path: VAPP_PATH_NAME,
        loadChildren: () => import('./vapp/vapp.module').then(m => m.VappModule)
      },
      {
        path: RESOURCE_PATH_NAME,
        loadChildren: () => import('./resource/resource.module').then(m => m.ResourceModule),
      },
      {
        path: POLICY_PATH_NAME,
        loadChildren: () => import('./policy/policy.module').then(m => m.PolicyModule),
      },
      {
        path: TEMPLATE_PATH_NAME,
        loadChildren: () => import('./template/template.module').then(m => m.TemplateModule),
      },
      {
        path: ORG_PATH_NAME,
        loadChildren: () => import('./org/org.module').then(m => m.OrgModule),
      },
      {
        path: USER_PATH_NAME,
        loadChildren: () => import('./user/user.module').then(m => m.UserModule),
      },
      {
        path: NOTEBOOKS_PATH_NAME,
        loadChildren: () => import('./notebooks/notebooks.module').then(m => m.NotebooksModule),
      }
    ],
  },
];

export const declarations = [
  LoginComponent,
  RegisterComponent,
  LandingComponent,
  ResetComponent,
];

export const providers = [
  AuthGuard,
  LoginService,
  RegisterService,
  ResetService,
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    scrollPositionRestoration: 'enabled',
    anchorScrolling: 'enabled',
    scrollOffset: [0, 0],
  })],
  exports: [RouterModule],
})
export class AppRoutingModule {}

export interface RouterData {
  anonymous?: boolean;
}
