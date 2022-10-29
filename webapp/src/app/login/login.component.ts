import { Component, OnInit } from '@angular/core';
import { CgEdgeUsersService, User } from '../cg-edge-users.service'
import { Router } from '@angular/router';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../message-popup/message-popup.component';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  User: User = new User();
  isValidUser!: boolean;
  

  constructor(private CgEdgeUsersService: CgEdgeUsersService,
              private router: Router, public dialog: MatDialog) { }

  ngOnInit(): void {

  }

  login() {
    this.CgEdgeUsersService.Login(this.User).subscribe((data) => {
      this.isValidUser = (data as boolean);
      if (!this.isValidUser) {
        this.dialog.open(MessagePopupComponent, {data: {title: "Login", text: "Invalid Username or Password!"}});
      }
      this.router.navigate(['/Dashboard']);
    });
    
  }
}
