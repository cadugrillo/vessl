import { Injectable, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { VesslUsersService } from './vessl-users.service';
import { environment } from '../../environments/environment';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class VesslContainersService {

  templateToInstall!: Template;

  constructor(private httpClient: HttpClient,
              private VesslUsersService: VesslUsersService) {}

  httpOptions = {
    headers: new HttpHeaders({
      "Authorization": this.VesslUsersService.getCurrentUserApiKey()
    })
  };

  getContainers(NetworkName: string) {
    return this.httpClient.get(environment.gateway + '/containers/json/' + NetworkName, this.httpOptions);
  }

  getContainersLogs(Id: string) {
    return this.httpClient.get(environment.gateway + '/containers/'+ Id + '/logs', this.httpOptions);
  }

  installContainer(AppTemplate: Template) {
    return this.httpClient.post(environment.gateway + '/containers/install', AppTemplate, this.httpOptions);
  }

  startContainer(Id: string) {
    return this.httpClient.post(environment.gateway + '/containers/'+ Id + '/start', '', this.httpOptions);
  }

  stopContainer(Id: string) {
    return this.httpClient.post(environment.gateway + '/containers/'+ Id + '/stop', '', this.httpOptions);
  }

  restartContainer(Id: string) {
    return this.httpClient.post(environment.gateway + '/containers/'+ Id + '/restart', '', this.httpOptions);
  }

  removeContainer(Id: string) {
    return this.httpClient.post(environment.gateway + '/containers/'+ Id + '/remove', '', this.httpOptions);
  }

  getContainerStats(Id: string) {
    return this.httpClient.get(environment.gateway + '/containers/'+ Id + '/stats', this.httpOptions);
  }

  getCompleteStats() {
    return this.httpClient.get(environment.gateway + '/containers/cstats', this.httpOptions);
  }

  getContainersRepo(UserId: string) {
    return this.httpClient.get(environment.gateway + '/containers/repository/' + UserId, this.httpOptions);
  }

  getDockerServerInfo() {
    return this.httpClient.get(environment.gateway + '/containers/info', this.httpOptions);
  }

  setTemplateToInstall(AppTemplate: Template) {
    this.templateToInstall = AppTemplate;
  }

  getTemplateToInstall(): Observable<Template> {
    return of(this.templateToInstall);
  }
}

export class Container {
  Id!: string
  Names!: string[]
  ImageID!: string
  Command!: string
  Created!: EpochTimeStamp
  State!: string
  Status!: string
  Ports!: Port[]
  //Labels!: any
  SizeRw!: number
  SizeRootFs!: number
  HostConfig!: HostConfig
  //NetworkSettings!: any
  //Mounts!: any
}

class Port {
  PrivatePort!: number
  PublicPort!: number
  Type!: string
}

class HostConfig {
  NetworkMode!: string
}

export class ContainersRepo {
  version!: string
  templates!: Template[]
}

export class Template {
  type!: string
	title!: string
	name!: string
	hostname!: string
	description!: string
  info_url!: string
	categories!: string[]
	platform!: string
	logo!: string
	image!: string
	restart_policy!: string
	network!: string
  env!: string[]
	ports!: string[]
	volumes!: string[]
}

export class ContainerStats {
  Id!: string
  Name!: string
  CpuPct!: number
  MemUsage!: number
  MemLimit!: number
  MemPct!: number
}

export class SystemInfo {
  Architecture!: string
  BridgeNfIp6tables!: boolean
  BridgeNfIptables!: boolean
  CPUSet!: boolean
  CPUShares!: boolean
  CgroupDriver!: string
  CgroupVersion!: string
  Containers!: number
  ContainersPaused!: number
  ContainersRunning!: number
  ContainersStopped!: number
  CpuCfsPeriod!: boolean
  CpuCfsQuota!: boolean
  Debug!: boolean
  DefaultRuntime!: string
  DockerRootDir!: string
  Driver!: string
  ExperimentalBuild!: boolean
  HttpProxy!: string
  HttpsProxy!: string
  ID!: string
  IPv4Forwarding!: boolean
  Images!: number
  IndexServerAddress!: string
  InitBinary!: string
  Isolation!: string
  KernelMemory!: boolean
  KernelMemoryTCP!: boolean
  KernelVersion!: string
  LiveRestoreEnabled!: boolean
  LoggingDriver!: string
  MemTotal!: number
  MemoryLimit!: boolean
  NCPU!: number
  NEventsListener!: number
  NFd!: number
  NGoroutines!: number
  Name!: string
  NoProxy!: string
  OSType!: string
  OSVersion!: string
  OomKillDisable!: boolean
  OperatingSystem!: string
  PidsLimit!: boolean
  ServerVersion!: string
  SwapLimit!: boolean
  SystemTime!: string
  Warnings!: string[]
}