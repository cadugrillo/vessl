import { Component, OnInit } from '@angular/core';
import { VesslUsersService, Users, User } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../../popups/message-popup/message-popup.component';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  Users!: Users

  constructor(private VesslUsersService: VesslUsersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getUsers();
  }

  getUsers() {
    this.VesslUsersService.getUsers().subscribe((data) => {
      this.Users = (data as Users);
    });
  }

  updateUser(User: User) {
    this.VesslUsersService.updateUser(User).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Update Users", text: data}});
      this.getUsers();
    });
  }

  addUser() {
    this.VesslUsersService.addUser().subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Add User", text: data}});
      this.getUsers();
    });
  }

  deleteUser(Id: string) {
    this.VesslUsersService.deleteUser(Id).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Delete User", text: data}});
      this.getUsers();
    });
  }

  AppsRepEditEnabled() {
    if (this.VesslUsersService.CurrentUser.Username == 'master') {
      return true
    }
    return false
  }

}
