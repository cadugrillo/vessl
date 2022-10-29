import { Component, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';
import { CgEdgeSystemService } from '../cg-edge-system.service';


export interface DialogData {
  title: string;
  text: string;
}

@Component({
  selector: 'app-sys-restart-popup',
  templateUrl: './sys-restart-popup.component.html',
  styleUrls: ['./sys-restart-popup.component.css']
})
export class SysRestartPopupComponent {

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData,
              private CgEdgeSystemService: CgEdgeSystemService) { }


  restartHostSystem() {
    this.CgEdgeSystemService.restartHostSystem().subscribe((data) => {
      console.log(data);
    });
  }
}
