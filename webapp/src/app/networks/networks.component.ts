import { Component, OnInit } from '@angular/core';
import { CgEdgeNetworksService, Network } from '../cg-edge-networks.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../message-popup/message-popup.component';
import { WaitPopupComponent } from '../wait-popup/wait-popup.component';

@Component({
  selector: 'app-networks',
  templateUrl: './networks.component.html',
  styleUrls: ['./networks.component.css']
})
export class NetworksComponent implements OnInit {

  networks!: Network[]
  networkName!: string

  constructor(private CgEdgeNetworksService: CgEdgeNetworksService,
    public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getNetworks();
  }

  getNetworks() {
    this.CgEdgeNetworksService.getNetworks().subscribe((data) => {
      this.networks = (data as Network[]);
    });
  }

  createNetwork() {
    this.dialog.open(WaitPopupComponent, {});
    this.CgEdgeNetworksService.createNetwork(this.networkName).subscribe((data) => {
      this.dialog.closeAll();
      this.dialog.open(MessagePopupComponent, {data: {title: "Network Creation", text: data}});
      this.getNetworks();
    });
  }

  removeNetwork(Id: string) {
    this.CgEdgeNetworksService.removeNetwork(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove Network", text: data}});
      this.getNetworks();
    });
  }

}
