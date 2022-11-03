import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { VesslContainersService, ContainersRepo, Template } from '../../services/vessl-containers.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
@Component({
  selector: 'app-app-repository',
  templateUrl: './app-repository.component.html',
  styleUrls: ['./app-repository.component.css']
})
export class AppRepositoryComponent implements OnInit {

  containersRepo!: ContainersRepo

  constructor(private VesslContainerService: VesslContainersService,
              private VesslUsersService: VesslUsersService,
              private router: Router, public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getContainersRepo();
  }

  getContainersRepo() {
    this.VesslContainerService.getContainersRepo(this.VesslUsersService.CurrentUser.ID).subscribe((data) => {
      this.containersRepo = (data as ContainersRepo);
    });
  }

  installContainer(AppTemplate: Template) {
    this.VesslContainerService.setTemplateToInstall(AppTemplate);
    this.router.navigate(['/App-Launcher']);
  }

  getInfo(info_url: string) {
    window.open(info_url,'_blank')
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }
}
