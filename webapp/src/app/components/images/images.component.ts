import { Component, OnInit } from '@angular/core';
import { VesslImagesService, Image } from '../../services/vessl-images.service';
import { VesslUsersService } from '../../services/vessl-users.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../../popups/message-popup/message-popup.component';
import { saveAs } from "file-saver";

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

  inspectImage(Id: string) {
    this.VesslImagesService.inspectImage(Id).subscribe((data) =>{
      return saveAs(new Blob([JSON.stringify(data, null, 2)], { type: 'JSON' }), 'inspect_image_'+Id.substring(0,12)+'.json');
    });
  }

  hideImage(imageName: string) {
   return !imageName.includes("vessl/vessl");
  }

  roleExpert() {
    return this.VesslUsersService.expert();
  }

}
