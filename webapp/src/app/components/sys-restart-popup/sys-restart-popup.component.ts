import { Component, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';
import { VesslSystemService } from '../../services/vessl-system.service';


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
              private VesslSystemService: VesslSystemService) { }


  restartHostSystem() {
    this.VesslSystemService.restartHostSystem().subscribe((data) => {
      console.log(data);
    });
  }
}
