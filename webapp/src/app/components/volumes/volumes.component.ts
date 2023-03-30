import { Component, OnInit } from '@angular/core';
import { VesslVolumesService, VolumeList } from '../../services/vessl-volumes.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../../popups/message-popup/message-popup.component';
import { saveAs } from "file-saver";

@Component({
  selector: 'app-volumes',
  templateUrl: './volumes.component.html',
  styleUrls: ['./volumes.component.css']
})
export class VolumesComponent implements OnInit {

  volumes!: VolumeList

  constructor(private VesslVolumesService: VesslVolumesService,
              private VesslUsersService: VesslUsersService,
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

  inspectVolume(Id: string) {
    this.VesslVolumesService.inspectVolume(Id).subscribe((data) =>{
      return saveAs(new Blob([JSON.stringify(data, null, 2)], { type: 'JSON' }), 'inspect_volume_'+Id+'.json');
    });
  }

  hideVolume(volumeName: string) {
    return !volumeName.includes("vessl");
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }

}