import { Component } from '@angular/core';
import { CgEdgeUsersService, User } from './cg-edge-users.service'
import { Router } from '@angular/router';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'CG-EDGE-CONF';

  

  constructor(private CgEdgeUsersService: CgEdgeUsersService,
              private router: Router) {
  
  }

  public ngOnInit(): void {
  
  }

  logout() {
    this.CgEdgeUsersService.logout();
    this.router.navigate(['/Login']).then(() => {window.location.reload();});
  }

  isAuthenticated() {
    return this.CgEdgeUsersService.isAuthenticated();
  }

  openWebPage() {
    window.open('https://github.com/cadugrillo/cg-edge-configurator', '_blank');
  }
}

