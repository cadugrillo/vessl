import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, Router } from '@angular/router';
import { CgEdgeUsersService } from './cg-edge-users.service'

@Injectable({
  providedIn: 'root'
})

export class AuthGuardService implements CanActivate {

  constructor(private CgEdgeUsersService: CgEdgeUsersService,
    private router: Router) { }


  // canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<boolean> {
   
  //   return new Promise((resolve, reject) => {
      
  //     return resolve (this.CgEdgeUsersService.isAuthenticated())

  //   });
  // }

  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
   
    return this.CgEdgeUsersService.isAuthenticated()
    
  }


}
