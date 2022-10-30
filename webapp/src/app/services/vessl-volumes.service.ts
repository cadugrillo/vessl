import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { VesslUsersService } from './vessl-users.service';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class VesslVolumesService {

  constructor(private httpClient: HttpClient,
              private VesslUsersService: VesslUsersService) {}

httpOptions = {
headers: new HttpHeaders({
"Authorization": this.VesslUsersService.getCurrentUserApiKey()
})
};

  getVolumes() {
    return this.httpClient.get(environment.gateway + '/volumes/json', this.httpOptions);
  }

  removeVolume(Id: string) {
    return this.httpClient.post(environment.gateway + '/volumes/'+ Id + '/remove', '', this.httpOptions);
  }
}

export class VolumeList {
  Volumes!: Volume[]
  Warnings!: string[]
}

class Volume {
  Name!: string
  Driver!: string
  Mountpoint!: string
  CreatedAt!: string
  //Status!: Map<string,string>
  Labels!: Map<string,string>
  Scope!: string
  Options!: Map<string,string>
  UsageData!: UsageData
}

class UsageData {
  Size!: number
  RefCount!: number
}
