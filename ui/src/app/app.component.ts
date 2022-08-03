import { Component, OnInit, OnChanges, SimpleChanges } from '@angular/core';
import { Router } from '@angular/router';
import { PRODUCT_NAME } from '@tide-config/const';

import { LoginService, UserInfo } from './login/login.service';
import { TranslateService } from '@ngx-translate/core';
import { I18nService } from '@tide-shared/service/i18n';
import { Observable, Subject, timer } from 'rxjs';
import { RegisterService } from './register/register.service';
import { ResetService } from './reset/reset.service';

import { Location } from '@angular/common';
import { NavigationEnd } from '@angular/router';

// Testing 6.29
import { environment } from '@tide-environments/environment';
import { ORG_PATH, USER_PATH } from "@tide-config/path";

@Component({
  selector: 'tide-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  path = '';


  constructor(
    readonly loginService: LoginService,
    readonly resetService: ResetService,
    readonly registerService: RegisterService,
    private readonly router: Router,
    translate: TranslateService,
    public readonly i18nService: I18nService,
    private location: Location
  ) {
    // translate.addLangs(['en', 'zh-CN']);
    // // this language will be used as a fallback when a translation isn't found in the current language
    // translate.setDefaultLang('zh-CN');

    

    this.router.events.subscribe((val) => {
      this.path = this.location.path();
      // if (this.path == '/login') {
      //   window.location.reload();
      // }
    });

  }
  // useLanguage(language: string): void {
  //   this.translate.use(language);
  // }

  // redirectTo(uri:string){
  //   this.router.navigateByUrl('/', {skipLocationChange: true}).then(()=>
  //   this.router.navigate([uri]));
  //   }

  redirectToLogin() {
    this.router.navigate(['/login'])
    .then(() => {
      window.location.reload()
    })
  }

  routeToVcpp() {
    this.router.navigate(['/vcpp'])
  }

  reloadCurrentPage() {
    window.location.reload();
    }
  readonly vo = {
    title: PRODUCT_NAME,
  };

  subject = new Subject();

  signOut() {
    this.loginService.logout();
  }

  cloudtides_logout() {
    this.loginService.cloudtides_logout()
  }

  cloudtides_reset_code() {
    this.resetService.cloud_reset_code(this.loginService.session.username);
    console.log("here!");
    this.router.navigate(['/cloudtides/reset']);
    
  }

  ngOnInit(): void {
    // 前端埋点测试
    // window.onload = () => {
    //   this.getAllNodes(document.body.children[0])
    // }


    
    //Testing for 6-29

    console.log(environment.apiPrefix + USER_PATH ); 


  }
  getAllNodes(doms){
    // if (doms.children.length > 0) {
    //   for(let d=0; d < doms.children.length; d++) {        
    //     if (doms.children[d].children.length > 0) {
    //       this.getAllNodes(doms.children[d])
    //     } else if (doms.children[d].innerText !== '') {
    //       doms.children[d].timer = 0
    //       doms.children[d].onmouseenter = function (e) {
    //         doms.children[d].timer = Date.now()
    //       }
    //       doms.children[d].onmouseleave = function (e) {
    //         const timeDiff = Date.now() - doms.children[d].timer
    //         console.log(doms.children[d].innerText, '停留', timeDiff/1000 + '秒'); 
    //         doms.children[d].timer = Date.now()
    //       }
    //       doms.children[d].onclick = function (e) {
    //         console.log(e.target.innerText, '点击'); 
    //       }
    //     }
    //   }
    // }
    const header = document.querySelector('.header')
    const nav = document.querySelector('.subnav')
    const container = document.querySelector('.main-container')
    console.log(header.children, nav.children, container.children);
    console.log(textNodesUnder(header).filter(el => el.data!=='\n            '));
    
    function textNodesUnder(node){
      var all = [];
      for (node=node.firstChild;node;node=node.nextSibling){
        if (node.nodeType==3) all.push(node);
        else all = all.concat(textNodesUnder(node));
      }
      return all;
    }    
    // header.addEventListener('mouseleave', (e) => {
    //   header[timer] = Date.now()
    //   header[obj] = e.target
    // })
  };
}
