import { Component, OnInit } from '@angular/core';
import { CgEdgeSystemService, InterfaceSet } from '../cg-edge-system.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent} from '../message-popup/message-popup.component';

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

  constructor(private CgEdgeSystemService: CgEdgeSystemService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getHostNetwork();
  }

  getHostNetwork() {
    this.CgEdgeSystemService.getHostNetwork().subscribe((data) => {
      this.InterfaceSet = (data as InterfaceSet);
    });
  }

  setHostNetwork() {
    this.CgEdgeSystemService.setHostNetwork(this.InterfaceSet).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Network Settings", text: data}});
      this.getHostNetwork();
    });
  }
}
