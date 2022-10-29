import { Component, OnInit } from '@angular/core';
import { CgEdgeNetworksService, Network } from '../cg-edge-networks.service';
import { CgEdgeContainersService, Template } from '../cg-edge-containers.service';
import { CgEdgeUsersService } from '../cg-edge-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent} from '../message-popup/message-popup.component';
import { WaitPopupComponent } from '../wait-popup/wait-popup.component';

@Component({
  selector: 'app-app-launcher',
  templateUrl: './app-launcher.component.html',
  styleUrls: ['./app-launcher.component.css']
})
export class AppLauncherComponent implements OnInit {

  appTemplate: Template = new Template();
  newPort!: string
  newEnv!: string
  newVolume!: string
  networks!: Network[]
  networkName: string = "cg-edge"

  Sources: string[] = ['no', 'on-failure', 'always', 'unless-stopped'];

  constructor(private CgEdgeContainerService: CgEdgeContainersService,
              private CgEdgeNetworksService: CgEdgeNetworksService,
              private CgEdgeUsersService: CgEdgeUsersService,
    public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getNetworks();
    this.appTemplate.ports = [];
    this.appTemplate.env = [];
    this.appTemplate.volumes = [];
    this.appTemplate.restart_policy = "always"
  }

  trackByFn(index: any, item: any) {
    return index;
 }

  installContainer() {
    this.dialog.open(WaitPopupComponent, {});
    this.appTemplate.network = this.networkName
    this.CgEdgeContainerService.installContainer(this.appTemplate).subscribe((data) => {
      this.dialog.closeAll();
      this.dialog.open(MessagePopupComponent, {data: {title: "App Installation", text: data}});
    });
  }

  addPort() {
    this.newPort = "";
    this.appTemplate.ports.push(this.newPort);
  }

  deletePort() {
    this.appTemplate.ports.splice(-1);
  }

  addEnv() {
    this.newEnv = "";
    this.appTemplate.env.push(this.newEnv);
  }

  deleteEnv() {
    this.appTemplate.env.splice(-1);
  }

  addVolume() {
    this.newVolume = "";
    this.appTemplate.volumes.push(this.newVolume);
  }

  deleteVolume() {
    this.appTemplate.volumes.splice(-1);
  }

  getNetworks() {
    this.CgEdgeNetworksService.getNetworks().subscribe((data) => {
      this.networks = (data as Network[]);
    });
  }

  filterEnabled() {
    if (this.CgEdgeUsersService.CurrentUser.Username == 'master') {
      return true
    }
    return false
  }
}
