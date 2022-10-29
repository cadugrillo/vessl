import { Component, OnInit } from '@angular/core';
import { CgEdgeContainersService, ContainersRepo, Template } from '../cg-edge-containers.service';
import { CgEdgeUsersService } from '../cg-edge-users.service';
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

  constructor(private CgEdgeContainerService: CgEdgeContainersService,
              private CgEdgeUsersService: CgEdgeUsersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getContainersRepo();
  }

  getContainersRepo() {
    this.CgEdgeContainerService.getContainersRepo(this.CgEdgeUsersService.CurrentUser.ID).subscribe((data) => {
      this.containersRepo = (data as ContainersRepo);
    });
  }

  installContainer(AppTemplate: Template) {
    this.dialog.open(WaitPopupComponent, {});
    this.CgEdgeContainerService.installContainer(AppTemplate).subscribe((data) => {
      this.dialog.closeAll();
      this.dialog.open(MessagePopupComponent, {data: {title: "App Installation", text: data}});
    });
  }

  getInfo(info_url: string) {
    window.open(info_url,'_blank')
  }
}
