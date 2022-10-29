import { Component, OnInit } from '@angular/core';
import { CgEdgeImagesService, Image } from '../cg-edge-images.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent } from '../message-popup/message-popup.component';

@Component({
  selector: 'app-images',
  templateUrl: './images.component.html',
  styleUrls: ['./images.component.css']
})
export class ImagesComponent implements OnInit {

  images!: Image[]

  constructor(private CgEdgeImagesService: CgEdgeImagesService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getImages();
  }

  getImages() {
    this.CgEdgeImagesService.getImages().subscribe((data) => {
      this.images = (data as Image[]);
    });
  }

  removeImage(Id: string) {
    this.CgEdgeImagesService.removeImage(Id).subscribe((data) =>{
      this.dialog.open(MessagePopupComponent, {data: {title: "Remove Image", text: data}});
      this.getImages();
    });
  }

}
