import { Component, OnInit } from '@angular/core';
import { VesslNetworksService, Network } from '../../services/vessl-networks.service';
import { VesslContainersService, Container, ContainerStats } from '../../services/vessl-containers.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import { Router } from '@angular/router';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../../popups/message-popup/message-popup.component';
import { StatsPopupComponent } from '../../popups/stats-popup/stats-popup.component';
import { WaitPopupComponent } from '../wait-popup/wait-popup.component';
import { saveAs } from "file-saver";

@Component({
  selector: 'app-home',
  templateUrl: './apps.component.html',
  styleUrls: ['./apps.component.css']
})
export class AppsComponent implements OnInit {

  containers!: Container[]
  networks!: Network[]
  networkName: string = "cg-edge"

  constructor(private VesslContainerService: VesslContainersService,
              private VesslNetworksService: VesslNetworksService,
              private VesslUsersService: VesslUsersService,
              public dialog: MatDialog, private router: Router) { }

  ngOnInit(): void {
    this.getNetworks();
    this.getContainers(this.networkName);
  }

  getContainers(NetworkName: string) {
    this.VesslContainerService.getContainers(NetworkName).subscribe((data) => {
      console.log(data);
      this.containers = (data as Container[]);
    });
  }
  
  startContainer(Id: string) {
    this.VesslContainerService.startContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Start App", text: data}});
      this.getContainers("cg-edge");
    });
  }

  stopContainer(Id: string) {
    this.VesslContainerService.stopContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Stop App", text: data}});
      this.getContainers("cg-edge");
    });
  }

  restartContainer(Id: string) {
    this.VesslContainerService.restartContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Restart App", text: data}});
      this.getContainers("cg-edge");
    });
  }

  removeContainer(Id: string) {
    this.VesslContainerService.removeContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove App", text: data}});
      this.getContainers("cg-edge");
    });
  }

  getContainerStats(Id: string) {
    this.dialog.open(WaitPopupComponent, {});
    this.VesslContainerService.getContainerStats(Id).subscribe((data) => {
      this.dialog.closeAll();
      let containerStats = data as ContainerStats
      this.dialog.open(StatsPopupComponent, {data: {Title: "App Usage Statistics", CpuPct: containerStats.CpuPct,
        MemUsage: containerStats.MemUsage, MemLimit: containerStats.MemLimit, MemPct: containerStats.MemPct}});
    });
  }

  getContainerLogs(Id: string) {
    this.dialog.open(WaitPopupComponent, {});
    this.VesslContainerService.getContainersLogs(Id).subscribe((data) =>{
      this.dialog.closeAll();
      return saveAs(new Blob([data as string], { type: 'TXT' }), 'logs.txt');
    });
  }

  openConfig(Container: Container) {
    switch (Container.Names[0]) {
      case "/mqtt-cloud-connector":
        this.router.navigate(['/mqtt-cloud-connector']);
        break;
      case "/opcua-mqtt-connector":
        this.router.navigate(['/opcua-mqtt-connector']);
        break;
      default:
        window.open('http://' + window.location.hostname + ':' + Container.Ports[0].PublicPort,'_blank');
        break;
    }
  }

  getNetworks() {
    this.VesslNetworksService.getNetworks().subscribe((data) => {
      this.networks = (data as Network[]);
    });
  }

  filterEnabled() {
    if (this.VesslUsersService.CurrentUser.Username == 'master') {
      return true
    }
    return false
  }

}
