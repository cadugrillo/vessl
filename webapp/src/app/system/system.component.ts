import { Component, OnInit } from '@angular/core';
import { CgEdgeContainersService, SystemInfo } from '../cg-edge-containers.service';
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

  constructor(private CgEdgeContainerService: CgEdgeContainersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getDockerServerInfo();
  }

  getDockerServerInfo() {
    this.CgEdgeContainerService.getDockerServerInfo().subscribe((data) => {
      this.SystemInfo = (data as SystemInfo);
    });
  }

  restartHostSystem() {
    this.dialog.open(SysRestartPopupComponent, {data: {title: "Reboot System", text: "Are you sure you want to reboot the system?"}});
  }

  shutdownHostSystem() {
    this.dialog.open(SysShutdownPopupComponent, {data: {title: "Shutdown System", text: "Are you sure you want to shutdown the system?"}});
  }
}
