import { Component, OnInit } from '@angular/core';
import { VesslVolumesService, VolumeList } from '../../services/vessl-volumes.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../message-popup/message-popup.component';

@Component({
  selector: 'app-volumes',
  templateUrl: './volumes.component.html',
  styleUrls: ['./volumes.component.css']
})
export class VolumesComponent implements OnInit {

  volumes!: VolumeList

  constructor(private VesslVolumesService: VesslVolumesService,
    public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getVolumes();
  }

  getVolumes() {
    this.VesslVolumesService.getVolumes().subscribe((data) => {
      this.volumes = (data as VolumeList);
    });
  }

  removeVolume(Id: string) {
    this.VesslVolumesService.removeVolume(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove Volume", text: data}});
      this.getVolumes();
    });
  }

  hideVolume(volumeName: string) {
    return !volumeName.includes("vessl");
   }

}