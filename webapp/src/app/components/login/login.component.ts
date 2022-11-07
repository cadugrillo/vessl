import { Component, OnInit } from '@angular/core';
import { VesslUsersService, User } from '../../services/vessl-users.service'
import { Router } from '@angular/router';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../../popups/message-popup/message-popup.component';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  User: User = new User();
  isValidUser!: boolean;
  

  constructor(private VesslUsersService: VesslUsersService,
              private router: Router, public dialog: MatDialog) { }

  ngOnInit(): void {

  }

  login() {
    this.VesslUsersService.Login(this.User).subscribe((user) => {
      if (user.Username == "invalid" || user.Username == "") {
        this.dialog.open(MessagePopupComponent, {data: {title: "Login", text: "Invalid Username or Password."+user.Password}});
      }
      if (user.Username == "locked") {
        this.dialog.open(MessagePopupComponent, {data: {title: "Login", text: "System locked. Wait 5 minutes before next attempt"}});
      }
      this.router.navigate(['/Apps']);
    });
    
  }
}
