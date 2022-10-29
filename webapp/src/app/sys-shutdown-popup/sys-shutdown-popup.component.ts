import { Component, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';
import { CgEdgeSystemService } from '../cg-edge-system.service';


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
  private CgEdgeSystemService: CgEdgeSystemService) { }


  shutdownHostSystem() {
    this.CgEdgeSystemService.shutdownHostSystem().subscribe((data) => {
      console.log(data);
    });
  }
}
