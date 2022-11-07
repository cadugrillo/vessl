import { Component, OnInit } from '@angular/core';
import { VesslContainersService, SystemInfo } from '../../services/vessl-containers.service';
import {MatDialog} from '@angular/material/dialog';
import { SysRestartPopupComponent } from '../sys-restart-popup/sys-restart-popup.component';
import { SysShutdownPopupComponent } from '../sys-shutdown-popup/sys-shutdown-popup.component';

@Component({
  selector: 'app-system',
  templateUrl: './system.component.html',
  styleUrls: ['./system.component.css']
})
export class SystemComponent implements OnInit {

  SystemInfo: SystemInfo = new SystemInfo();

  constructor(private VesslContainerService: VesslContainersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getDockerServerInfo();
  }

  getDockerServerInfo() {
    this.VesslContainerService.getDockerServerInfo().subscribe((data) => {
      this.SystemInfo = (data as SystemInfo);
    });
  }
}
