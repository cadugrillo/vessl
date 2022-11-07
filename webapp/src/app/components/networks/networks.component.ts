import { Component, OnInit } from '@angular/core';
import { VesslNetworksService, Network } from '../../services/vessl-networks.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../../popups/message-popup/message-popup.component';
import { WaitPopupComponent } from '../wait-popup/wait-popup.component';

@Component({
  selector: 'app-networks',
  templateUrl: './networks.component.html',
  styleUrls: ['./networks.component.css']
})
export class NetworksComponent implements OnInit {

  networks!: Network[]
  networkName!: string

  constructor(private VesslNetworksService: VesslNetworksService,
    private VesslUsersService: VesslUsersService,
    public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getNetworks();
  }

  getNetworks() {
    this.VesslNetworksService.getNetworks().subscribe((data) => {
      this.networks = (data as Network[]);
    });
  }

  createNetwork() {
    this.dialog.open(WaitPopupComponent, {});
    this.VesslNetworksService.createNetwork(this.networkName).subscribe((data) => {
      this.dialog.closeAll();
      this.dialog.open(MessagePopupComponent, {data: {title: "Network Creation", text: data}});
      this.getNetworks();
    });
  }

  removeNetwork(Id: string) {
    this.VesslNetworksService.removeNetwork(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove Network", text: data}});
      this.getNetworks();
    });
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }

}
