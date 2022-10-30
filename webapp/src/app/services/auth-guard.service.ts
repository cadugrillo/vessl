import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, Router } from '@angular/router';
import { VesslUsersService } from './vessl-users.service'

@Injectable({
  providedIn: 'root'
})

export class AuthGuardService implements CanActivate {

  constructor(private VesslUsersService: VesslUsersService,
    private router: Router) { }


  // canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<boolean> {
   
  //   return new Promise((resolve, reject) => {
      
  //     return resolve (this.VesslUsersService.isAuthenticated())

  //   });
  // }

  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
   
    return this.VesslUsersService.isAuthenticated()
    
  }


}
