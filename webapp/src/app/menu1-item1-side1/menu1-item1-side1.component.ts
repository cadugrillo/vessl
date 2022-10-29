import { Component, OnInit } from '@angular/core';
import { CgEdgeUsersService, User } from '../cg-edge-users.service'

@Component({
  selector: 'menu1-item1-side1',
  templateUrl: './menu1-item1-side1.component.html',
  styleUrls: ['./menu1-item1-side1.component.css']
})
export class Menu1Item1Side1Component implements OnInit {

  currentUser!: User

  constructor(private CgEdgeUsersService: CgEdgeUsersService) {}

  ngOnInit(): void {
    
  }

  isAuthenticated() {
    return this.CgEdgeUsersService.isAuthenticated();
  }

  AppsMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.Apps
  }

  AppsRepMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.AppsRepository
  }

  AppLaunchMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.AppLauncher
  }

  ImagesMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.Images
  }

  VolumesMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.Volumes
  }

  NetworksMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.Networks
  }

  SettingsMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.Settings
  }

  UsersMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.Users
  }

  SystemMenuEnabled() {
    return this.CgEdgeUsersService.CurrentUser.Permissions.System
  }
}
