import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit'
import { environment } from '@tide-environments/environment';

@Component({
  selector: 'tide-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {

  constructor(private http:HttpClient) { }
  ngOnInit(): void {
  }
  users:any = []
}