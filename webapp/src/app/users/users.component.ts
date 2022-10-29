import { Component, OnInit } from '@angular/core';
import { CgEdgeUsersService, Users, User } from '../cg-edge-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../message-popup/message-popup.component';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  Users!: Users

  constructor(private CgEdgeUsersService: CgEdgeUsersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getUsers();
  }

  getUsers() {
    this.CgEdgeUsersService.getUsers().subscribe((data) => {
      this.Users = (data as Users);
    });
  }

  updateUser(User: User) {
    this.CgEdgeUsersService.updateUser(User).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Update Users", text: data}});
      this.getUsers();
    });
  }

  addUser() {
    this.CgEdgeUsersService.addUser().subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Add User", text: data}});
      this.getUsers();
    });
  }

  deleteUser(Id: string) {
    this.CgEdgeUsersService.deleteUser(Id).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Delete User", text: data}});
      this.getUsers();
    });
  }

  AppsRepEditEnabled() {
    if (this.CgEdgeUsersService.CurrentUser.Username == 'master') {
      return true
    }
    return false
  }

}
