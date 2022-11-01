import { Component, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';

export interface DialogData {
  Title: string;
  CpuPct: number
  MemUsage: number
  MemLimit: number
  MemPct: number
}

@Component({
  selector: 'app-stats-popup',
  templateUrl: './stats-popup.component.html',
  styleUrls: ['./stats-popup.component.css']
})

export class StatsPopupComponent {

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData) { }

 

}
