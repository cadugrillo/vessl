import { Component, OnDestroy, OnInit } from '@angular/core';
import { VesslSystemService, HostStats } from '../../services/vessl-system.service';
import { VesslContainersService, Template } from '../../services/vessl-containers.service';
import { Subscription, timer } from 'rxjs';
import { switchMap } from 'rxjs/operators';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit, OnDestroy {

  subscription !: Subscription;
  HostStats: HostStats = new HostStats();
  RamTotal!: string
  RamUsed!: string
  RamUsedBar!: number
  RamAvailable!: string
  RamAvailableBar!: number
  RamFree!: string
  RamFreeBar!: number
  DiskTotal!: string
  DiskAvailable!: string
  DiskAvailableBar!: number

  constructor(private VesslSystemService: VesslSystemService,
              private VesslContainerService: VesslContainersService) { }

  ngOnInit(): void {
    this.subscription = timer(0,5000).pipe(
      switchMap(() => this.VesslSystemService.getHostStats())).subscribe((data) => {
        this.getHostStats();
      });     
  }

  getHostStats() {
    this.VesslSystemService.getHostStats().subscribe((data) => {
      this.HostStats = (data as HostStats);
      this.RamTotal = (this.HostStats.RamTotal / 1000).toFixed(2);
      this.RamUsed = (this.HostStats.RamUsed / 1000).toFixed(2);
      this.RamAvailable = (this.HostStats.RamAvailable / 1000).toFixed(2);
      this.RamFree = (this.HostStats.RamFree / 1000).toFixed(2);
      this.RamUsedBar = (this.HostStats.RamUsed/this.HostStats.RamTotal) * 100;
      this.RamAvailableBar = (this.HostStats.RamAvailable/this.HostStats.RamTotal) * 100;
      this.RamFreeBar = (this.HostStats.RamFree/this.HostStats.RamTotal) * 100;
      this.DiskTotal = (this.HostStats.DiskTotal / 1000).toFixed(2);
      this.DiskAvailable = (this.HostStats.DiskAvailable / 1000).toFixed(2);
      this.DiskAvailableBar = (this.HostStats.DiskAvailable/this.HostStats.DiskTotal) * 100;
    });
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
}

}
