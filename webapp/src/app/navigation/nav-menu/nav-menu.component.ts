import { Component } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { VesslUsersService, User } from '../../services/vessl-users.service'
import { Router } from '@angular/router';

@Component({
  selector: 'app-nav-menu',
  templateUrl: './nav-menu.component.html',
  styleUrls: ['./nav-menu.component.css']
})
export class NavMenuComponent {

  isHandset$: Observable<boolean> = this.breakpointObserver.observe(Breakpoints.Handset)
    .pipe(
      map(result => result.matches),
      shareReplay()
    );

  constructor(private breakpointObserver: BreakpointObserver,private VesslUsersService: VesslUsersService,
    private router: Router) {}

    public ngOnInit(): void {
      this.VesslUsersService.deleteCookie();
      this.router.navigate(['/Login'])
    }
  
    logout() {
      this.VesslUsersService.logout();
      this.router.navigate(['/Login']).then(() => {window.location.reload();});
    }
  
    isAuthenticated() {
      return this.VesslUsersService.isAuthenticated();
    }
  
    loggedUser() {
      if (this.isAuthenticated()) {
        return this.VesslUsersService.CurrentUser.Username;
      }
      return ''
    }
  
    openWebPage() {
      window.open('https://github.com/cadugrillo/vessl', '_blank');
    }

    openDocs() {
      window.open('https://cadugrillo.github.io/vessl-docs/', '_blank');
    }

    AppsMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.Apps
    }
  
    AppsRepMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.AppsRepository
    }
  
    AppLaunchMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.AppLauncher
    }
  
    ImagesMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.Images
    }
  
    VolumesMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.Volumes
    }
  
    NetworksMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.Networks
    }
  
    SystemMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.System
    }
  
    UsersMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.Users
    }
  
    HostSettingMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.HostSettings
    }
  
    HostStatsMenuEnabled() {
      return this.VesslUsersService.CurrentUser.Permission.HostStats
    }
  
    currentRouterUrl() {
      return this.router.url
    }

}
