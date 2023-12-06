import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { CookieService } from 'ngx-cookie-service';
import { BehaviorSubject, Observable, Subject } from 'rxjs';
import { Template } from './vessl-containers.service';

@Injectable({
  providedIn: 'root'
})
export class VesslUsersService {

  CurrentUser!: User
  private authenticationSubject: BehaviorSubject<any>;

  constructor(private httpClient: HttpClient,
              private CookieService: CookieService,) {

    this.authenticationSubject = new BehaviorSubject<boolean>(false);
  }

  getUsers() {
    return this.httpClient.get(environment.gateway + '/users/json', {headers:({"Authorization": this.getCurrentUserApiKey()})});
  }

  updateUser(User: User) {
    return this.httpClient.post(environment.gateway + '/users/json', User, {headers:({"Authorization": this.getCurrentUserApiKey()})});
  }

  addUser() {
    return this.httpClient.get(environment.gateway + '/users/add', {headers:({"Authorization": this.getCurrentUserApiKey()})});
  }

  deleteUser(User: User) {
    return this.httpClient.post(environment.gateway + '/users/delete', User, {headers:({"Authorization": this.getCurrentUserApiKey()})});
  }

  validateUser(User: User) {
    return this.httpClient.post(environment.gateway + '/users/validate', User);
  }

  Login(User: User): Observable<User> {
    var subject = new Subject<any>();
    this.validateUser(User).subscribe((data) => {
        this.CurrentUser = data as User;
        if (this.CurrentUser.Username == "invalid" || this.CurrentUser.Username == "locked" || this.CurrentUser.Username == "") {
          this.authenticationSubject.next(false);
        } else {
          this.CookieService.set('vessl-token', this.CurrentUser.ApiKey);
          this.CurrentUser.ApiKey = '';
          this.authenticationSubject.next(true);
        }
        subject.next(this.CurrentUser)  
    });
    return subject.asObservable();
  }


  logout() {
    this.CurrentUser = {} as User
    this.CookieService.delete('vessl-token');
    this.authenticationSubject.next(false);
  }

  isAuthenticated(): boolean {
    return this.authenticationSubject.value
  }

  getCurrentUser(): User {
    return this.CurrentUser;
  }

  getCurrentUserApiKey(): string {
    return this.CookieService.get('vessl-token')
    //return this.CurrentUser.ApiKey;
  }

  deleteCookie() {
    this.CookieService.delete('vessl-token');
  }

  ////////ROLES///////////
  user() {
    if (this.CurrentUser.Role == 'User' || this.CurrentUser.Role == 'Vessl-User' || this.CurrentUser.Role == 'Expert' ||this.CurrentUser.Role == 'Admin') {
      return true
    }
    return false
  }

  vessl_user() {
    if (this.CurrentUser.Role == 'Vessl-User' || this.CurrentUser.Role == 'Expert' ||this.CurrentUser.Role == 'Admin') {
      return true
    }
    return false
  }

  expert() {
    if (this.CurrentUser.Role == 'Expert' ||this.CurrentUser.Role == 'Admin') {
      return true
    }
    return false
  }

  admin() {
    if (this.CurrentUser.Role == 'Admin') {
      return true
    }
    return false
  }
}
export class Users {
  Users!: User[]
}

export class User {
  ID!: number
  UUID!: string
  Username!: string
  Password!: string
  Role!: string
  FullName!: string
  Email!: string
  Telephone!: string
  AppsRepositoryUrl!: string
  ApiKey!: string
  Permission!: Permission
  Template!: Template
}

class Permission {
	Apps!: boolean
	AppsRepository!: boolean
  AppLauncher!: boolean
	Images!: boolean
	Volumes!: boolean
	Networks!: boolean
  System!: boolean
  Users!: boolean
  HostSettings!: boolean
  HostStats!: boolean
}
