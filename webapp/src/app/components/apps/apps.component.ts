import { Component, OnInit } from '@angular/core';
import { VesslNetworksService, Network } from '../../services/vessl-networks.service';
import { VesslContainersService, Container, ContainerStats, Template } from '../../services/vessl-containers.service';
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
  networkName: string = "vessl-default"
  emptyTemplate: Template = new Template();

  constructor(private VesslContainerService: VesslContainersService,
              private VesslNetworksService: VesslNetworksService,
              private VesslUsersService: VesslUsersService,
              public dialog: MatDialog, private router: Router) { }

  ngOnInit(): void {
    this.getNetworks();
    this.getContainers(this.networkName);
    this.emptyTemplate.ports = [];
    this.emptyTemplate.env = [];
    this.emptyTemplate.cmd = [];
    this.emptyTemplate.volumes = [];
    this.emptyTemplate.restart_policy = "always";
    this.emptyTemplate.network = "vessl-default";
    this.VesslContainerService.setTemplateToInstall(this.emptyTemplate);
  }

  getContainers(NetworkName: string) {
    this.VesslContainerService.getContainers(NetworkName).subscribe((data) => {
      this.containers = (data as Container[]);
    });
  }
  
  startContainer(Id: string) {
    this.VesslContainerService.startContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Start App", text: data}});
      this.getContainers(this.networkName);
    });
  }

  stopContainer(Id: string) {
    this.VesslContainerService.stopContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Stop App", text: data}});
      this.getContainers(this.networkName);
    });
  }

  restartContainer(Id: string) {
    this.VesslContainerService.restartContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Restart App", text: data}});
      this.getContainers(this.networkName);
    });
  }

  removeContainer(Id: string) {
    this.VesslContainerService.removeContainer(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove App", text: data}});
      this.getContainers(this.networkName);
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
      return saveAs(new Blob([data as string], { type: 'TXT' }), 'messages_'+Id.substring(0,11)+'.log');
    });
  }

  openConfig(Container: Container) {
        window.open('http://' + window.location.hostname + ':' + Container.Ports[0].PublicPort,'_blank');
  }

  getNetworks() {
    this.VesslNetworksService.getNetworks().subscribe((data) => {
      this.networks = (data as Network[]);
    });
  }

  roleUser() {
    return this.VesslUsersService.user();
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }

  hideContainer(ContainerName: string) {
    return !ContainerName.includes("vessl");
   }

}
