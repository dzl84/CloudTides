import { Component, OnInit } from '@angular/core';
import { AppComponent } from '../app.component';

@Component({
  selector: 'tide-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.scss']
})
export class LandingComponent implements OnInit {

  constructor(
    readonly appcomponent: AppComponent,
  ) { }

  goCTwiki(){
    window.location.href = 'https://github.com/cloudtides/CloudTides/wiki'
  }

  ngOnInit(): void {
  }

}
