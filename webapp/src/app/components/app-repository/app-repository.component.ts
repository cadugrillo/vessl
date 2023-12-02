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
  categories: string[] = ["All"]
  categoryName: string = 'All'
  selectedCategoryName: string = 'All'

  constructor(private VesslContainerService: VesslContainersService,
              private VesslUsersService: VesslUsersService,
              private router: Router, public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getContainersRepo();
  }

  getContainersRepo() {
    this.VesslContainerService.getContainersRepo(this.VesslUsersService.CurrentUser.UUID).subscribe((data) => {
      this.containersRepo = (data as ContainersRepo);
      for (var i=0; i < this.containersRepo.templates.length; i++) {
        for (var j=0; j < this.containersRepo.templates[i].categories.length; j++) {
          if (!this.categories.includes(this.containersRepo.templates[i].categories[j], 0)) {
            this.categories.push(this.containersRepo.templates[i].categories[j]);
          }
        } 
      }
    });
  }

  installContainer(AppTemplate: Template) {
    this.VesslContainerService.setTemplateToInstall(AppTemplate);
    this.router.navigate(['/App-Launcher']);
  }

  getInfo(info_url: string) {
    window.open(info_url,'_blank')
  }

  filterCategory(categoryName: string) {
    this.selectedCategoryName = categoryName;
  }

  filterTemplate(template: Template) {
    return template.categories.includes(this.selectedCategoryName, 0) || this.selectedCategoryName=='All'
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }

  roleVessl_User() {
    return this.VesslUsersService.vessl_user();
  }
}
