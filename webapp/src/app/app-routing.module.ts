import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AppLauncherComponent } from './components/app-launcher/app-launcher.component';
import { AppRepositoryComponent } from './components/app-repository/app-repository.component';
import { AppsComponent } from './components/apps/apps.component';
import { AuthGuardService } from './services/auth-guard.service';
import { DashboardComponent } from './components/host-stats/dashboard.component';
import { ImagesComponent } from './components/images/images.component';
import { LoginComponent } from './components/login/login.component';
import { NetworksComponent } from './components/networks/networks.component';
import { SettingsComponent } from './components/host-settings/settings.component';
import { SystemComponent } from './components/system/system.component';
import { UsersComponent } from './components/users/users.component';
import { VolumesComponent } from './components/volumes/volumes.component';




const routes: Routes = [

  { path: '', redirectTo: 'Login', pathMatch: 'full' },
  { path: 'Login', component: LoginComponent},
  { path: 'Dashboard', component: DashboardComponent, canActivate: [AuthGuardService]},
  { path: 'Apps', component: AppsComponent, canActivate: [AuthGuardService]},
  { path: 'App-Repository', component: AppRepositoryComponent, canActivate: [AuthGuardService]},
  { path: 'App-Launcher', component: AppLauncherComponent, canActivate: [AuthGuardService]},
  { path: 'Images', component: ImagesComponent, canActivate: [AuthGuardService]},
  { path: 'Volumes', component: VolumesComponent, canActivate: [AuthGuardService]},
  { path: 'Networks', component: NetworksComponent, canActivate: [AuthGuardService]},
  { path: 'Users', component: UsersComponent, canActivate: [AuthGuardService]},
  { path: 'System', component: SystemComponent, canActivate: [AuthGuardService]},
  { path: 'Settings', component: SettingsComponent, canActivate: [AuthGuardService]},
  { path: '**', redirectTo: 'Login'},
 

];

@NgModule({
  imports: [RouterModule.forRoot(routes, {onSameUrlNavigation: 'reload'})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
