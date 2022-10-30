import { Component, OnInit } from '@angular/core';
import { VesslContainersService, ContainersRepo, Template } from '../../services/vessl-containers.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent} from '../message-popup/message-popup.component';
import { WaitPopupComponent } from '../wait-popup/wait-popup.component';

@Component({
  selector: 'app-app-repository',
  templateUrl: './app-repository.component.html',
  styleUrls: ['./app-repository.component.css']
})
export class AppRepositoryComponent implements OnInit {

  containersRepo!: ContainersRepo

  constructor(private VesslContainerService: VesslContainersService,
              private VesslUsersService: VesslUsersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getContainersRepo();
  }

  getContainersRepo() {
    this.VesslContainerService.getContainersRepo(this.VesslUsersService.CurrentUser.ID).subscribe((data) => {
      this.containersRepo = (data as ContainersRepo);
    });
  }

  installContainer(AppTemplate: Template) {
    this.dialog.open(WaitPopupComponent, {});
    this.VesslContainerService.installContainer(AppTemplate).subscribe((data) => {
      this.dialog.closeAll();
      this.dialog.open(MessagePopupComponent, {data: {title: "App Installation", text: data}});
    });
  }

  getInfo(info_url: string) {
    window.open(info_url,'_blank')
  }
}
