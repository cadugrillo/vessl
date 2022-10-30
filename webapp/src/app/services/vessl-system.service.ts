import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { VesslUsersService } from './vessl-users.service';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class VesslSystemService {

  constructor(private httpClient: HttpClient,
              private VesslUsersService: VesslUsersService) {}

httpOptions = {
headers: new HttpHeaders({
"Authorization": this.VesslUsersService.getCurrentUserApiKey()
})
};

  getHostNetwork() {
    return this.httpClient.get(environment.gateway + '/system/hostnetwork', this.httpOptions);
  }

  setHostNetwork(InterfaceSet: InterfaceSet) {
    return this.httpClient.post(environment.gateway + '/system/hostnetwork', InterfaceSet, this.httpOptions);
  }

  restartHostSystem() {
    return this.httpClient.post(environment.gateway + '/system/restart', '', this.httpOptions);
  }

  shutdownHostSystem() {
    return this.httpClient.post(environment.gateway + '/system/shutdown', '', this.httpOptions);
  }

  getHostStats() {
    return this.httpClient.get(environment.gateway + '/system/hoststats', this.httpOptions);
  }

}

export class InterfaceSet {
  InterfacePath!: string
  Adapters!: Adapter[]
}

class Adapter {
  AddrFamily!: number
  AddrSource!: number
  Address!: string
  Auto!: boolean
  Broadcast!: string
  Gateway!: string
  Hotplug!: boolean
  Name!: string
  Netmask!: string
  Network!: string
}

export class HostStats {
  CpuUsage!: number[]
  RamTotal!: number
  RamUsed!: number
  RamUsedPct!: number
  RamAvailable!: number
  RamFree!: number
  DiskUsage!: number
  DiskAvailable!: number
  DiskTotal!: number
}
