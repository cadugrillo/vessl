import { Component, OnInit } from '@angular/core';
import { VesslUsersService, User } from '../../services/vessl-users.service'

@Component({
  selector: 'menu1-item1-side1',
  templateUrl: './menu1-item1-side1.component.html',
  styleUrls: ['./menu1-item1-side1.component.css']
})
export class Menu1Item1Side1Component implements OnInit {

  currentUser!: User

  constructor(private VesslUsersService: VesslUsersService) {}

  ngOnInit(): void {
    
  }

  isAuthenticated() {
    return this.VesslUsersService.isAuthenticated();
  }

  AppsMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.Apps
  }

  AppsRepMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.AppsRepository
  }

  AppLaunchMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.AppLauncher
  }

  ImagesMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.Images
  }

  VolumesMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.Volumes
  }

  NetworksMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.Networks
  }

  SystemMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.System
  }

  UsersMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.Users
  }

  HostSettingMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.HostSettings
  }

  HostStatsMenuEnabled() {
    return this.VesslUsersService.CurrentUser.Permissions.HostStats
  }
}
