import { Component, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';
import { VesslSystemService } from '../../services/vessl-system.service';


export interface DialogData {
  title: string;
  text: string;
}

@Component({
  selector: 'app-sys-shutdown-popup',
  templateUrl: './sys-shutdown-popup.component.html',
  styleUrls: ['./sys-shutdown-popup.component.css']
})
export class SysShutdownPopupComponent {

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData,
  private VesslSystemService: VesslSystemService) { }


  shutdownHostSystem() {
    this.VesslSystemService.shutdownHostSystem().subscribe((data) => {
      console.log(data);
    });
  }
}
