import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { VesslUsersService } from './vessl-users.service';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class VesslImagesService {

  constructor(private httpClient: HttpClient,
              private VesslUsersService: VesslUsersService) {}

  httpOptions = {
    headers: new HttpHeaders({
      "Authorization": this.VesslUsersService.getCurrentUserApiKey()
    })
  };

  getImages() {
    return this.httpClient.get(environment.gateway + '/images/json', this.httpOptions);
  }

  removeImage(Id: string) {
    return this.httpClient.post(environment.gateway + '/images/'+ Id + '/remove', '', this.httpOptions);
  }

  inspectImage(Id: string) {
    return this.httpClient.post(environment.gateway + '/images/'+ Id + '/inspect', '', this.httpOptions);
  }
}

export class Image {
  Id!: string
  ParentId!: string
  RepoTags!: string[]
  RepoDigests!: string[]
  Created!: EpochTimeStamp
  Size!: number
  VirtualSize!: number
  SharedSize!: number
  Labels!: {}
  Containers!: number
}