import { Component, OnInit } from '@angular/core';
import { VesslSystemService, InterfaceSet } from '../../services/vessl-system.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent} from '../../popups/message-popup/message-popup.component';
import { SysRestartPopupComponent } from '../sys-restart-popup/sys-restart-popup.component';
import { SysShutdownPopupComponent } from '../sys-shutdown-popup/sys-shutdown-popup.component';

interface Source {
  value: number;
  viewValue: string;
}

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {

  InterfaceSet: InterfaceSet = new InterfaceSet();

  Sources: Source[] = [
    {value: 1, viewValue: 'DHCP'},
    {value: 2, viewValue: 'STATIC'},
    {value: 3, viewValue: 'LOOPBACK'},
    {value: 4, viewValue: 'MANUAL'},
    
  ];

  constructor(private VesslSystemService: VesslSystemService,
              private VesslUsersService: VesslUsersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getHostNetwork();
  }

  getHostNetwork() {
    this.VesslSystemService.getHostNetwork().subscribe((data) => {
      this.InterfaceSet = (data as InterfaceSet);
    });
  }

  setHostNetwork() {
    this.VesslSystemService.setHostNetwork(this.InterfaceSet).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Network Settings", text: data}});
      this.getHostNetwork();
    });
  }

  restartHostSystem() {
    this.dialog.open(SysRestartPopupComponent, {data: {title: "Reboot System", text: "Are you sure you want to reboot the system?"}});
  }

  shutdownHostSystem() {
    this.dialog.open(SysShutdownPopupComponent, {data: {title: "Shutdown System", text: "Are you sure you want to shutdown the system?"}});
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }
}
