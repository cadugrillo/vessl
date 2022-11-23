import { Component, OnInit } from '@angular/core';
import { VesslImagesService, Image } from '../../services/vessl-images.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../../popups/message-popup/message-popup.component';

@Component({
  selector: 'app-images',
  templateUrl: './images.component.html',
  styleUrls: ['./images.component.css']
})
export class ImagesComponent implements OnInit {

  images!: Image[]

  constructor(private VesslImagesService: VesslImagesService,
              private VesslUsersService: VesslUsersService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getImages();
  }

  getImages() {
    this.VesslImagesService.getImages().subscribe((data) => {
      this.images = (data as Image[]);
    });
  }

  removeImage(Id: string) {
    this.VesslImagesService.removeImage(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove Image", text: data}});
      this.getImages();
    });
  }

  hideImage(imageName: string) {
   return !imageName.includes("vessl/vessl");
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }

}
