import { Component } from '@angular/core';
import { VesslUsersService, User } from './services/vessl-users.service'
import { Router } from '@angular/router';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'CG-EDGE-CONF';

  

  constructor(private VesslUsersService: VesslUsersService,
              private router: Router) {
  
  }

  public ngOnInit(): void {
  
  }

  logout() {
    this.VesslUsersService.logout();
    this.router.navigate(['/Login']).then(() => {window.location.reload();});
  }

  isAuthenticated() {
    return this.VesslUsersService.isAuthenticated();
  }

  openWebPage() {
    window.open('https://github.com/cadugrillo/vessl', '_blank');
  }
}

