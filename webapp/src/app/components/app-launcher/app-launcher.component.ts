import { Component, OnDestroy, OnInit } from '@angular/core';
import { VesslNetworksService, Network } from '../../services/vessl-networks.service';
import { VesslContainersService, Template } from '../../services/vessl-containers.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent} from '../../popups/message-popup/message-popup.component';
import { WaitPopupComponent } from '../wait-popup/wait-popup.component';

@Component({
  selector: 'app-app-launcher',
  templateUrl: './app-launcher.component.html',
  styleUrls: ['./app-launcher.component.css']
})
export class AppLauncherComponent implements OnInit, OnDestroy {

  appTemplate!: Template;
  newPort!: string
  newEnv!: string
  newVolume!: string
  networks!: Network[]
  networkName: string = "cg-edge"

  Sources: string[] = ['no', 'on-failure', 'always', 'unless-stopped'];

  constructor(private VesslContainerService: VesslContainersService,
              private VesslNetworksService: VesslNetworksService,
              private VesslUsersService: VesslUsersService,
    public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getNetworks();
    this.getTemplateToInstall();
  }

  ngOnDestroy(): void {
    this.appTemplate = new Template();
    this.appTemplate.ports = [];
    this.appTemplate.env = [];
    this.appTemplate.volumes = [];
    this.appTemplate.restart_policy = "always"
    this.VesslContainerService.setTemplateToInstall(this.appTemplate);
  }

  trackByFn(index: any, item: any) {
    return index;
 }

  installContainer() {
    this.dialog.open(WaitPopupComponent, {});
    this.appTemplate.network = this.networkName
    this.VesslContainerService.installContainer(this.appTemplate).subscribe((data) => {
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
    this.VesslNetworksService.getNetworks().subscribe((data) => {
      this.networks = (data as Network[]);
    });
  }

  getTemplateToInstall() {
    this.VesslContainerService.getTemplateToInstall().subscribe((template) => {
      this.appTemplate = template;
    });
  }

  filterEnabled() {
    if (this.VesslUsersService.CurrentUser.Username == 'master') {
      return true
    }
    return false
  }
}
