import { Component, OnInit } from '@angular/core';
import { CgEdgeVolumesService, VolumeList } from '../cg-edge-volumes.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../message-popup/message-popup.component';

@Component({
  selector: 'app-volumes',
  templateUrl: './volumes.component.html',
  styleUrls: ['./volumes.component.css']
})
export class VolumesComponent implements OnInit {

  volumes!: VolumeList

  constructor(private CgEdgeVolumesService: CgEdgeVolumesService,
    public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getVolumes();
  }

  getVolumes() {
    this.CgEdgeVolumesService.getVolumes().subscribe((data) => {
      this.volumes = (data as VolumeList);
    });
  }

  removeVolume(Id: string) {
    this.CgEdgeVolumesService.removeVolume(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove Volume", text: data}});
      this.getVolumes();
    });
  }

}