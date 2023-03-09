import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';

import { AppComponent } from './app.component';
import { AppsComponent } from './components/apps/apps.component';
import { FormsModule } from '@angular/forms';
import { TokenInterceptor } from './token.interceptor';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatCardModule} from '@angular/material/card';
import { MessagePopupComponent } from './popups/message-popup/message-popup.component';
import {MatDialogModule} from '@angular/material/dialog';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatTabsModule} from '@angular/material/tabs';
import {MatMenuModule} from '@angular/material/menu';
import {MatListModule} from '@angular/material/list';
import {MatSelectModule} from '@angular/material/select';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import { StatsComponent } from './components/host-stats/stats.component';
import { AppRepositoryComponent } from './components/app-repository/app-repository.component';
import { UsersComponent } from './components/users/users.component';
import { SystemComponent } from './components/system/system.component';
import { SettingsComponent } from './components/host-settings/settings.component';
import { VesslContainersService } from './services/vessl-containers.service';
import { VesslSystemService } from './services/vessl-system.service';
import { WaitPopupComponent } from './components/wait-popup/wait-popup.component';
import { SysRestartPopupComponent } from './components/sys-restart-popup/sys-restart-popup.component';
import { SysShutdownPopupComponent } from './components/sys-shutdown-popup/sys-shutdown-popup.component';
import { VesslUsersService } from './services/vessl-users.service';
import { LoginComponent } from './components/login/login.component';
import { ImagesComponent } from './components/images/images.component';
import { VesslImagesService } from './services/vessl-images.service';
import { AppLauncherComponent } from './components/app-launcher/app-launcher.component';
import { VesslVolumesService } from './services/vessl-volumes.service';
import { VolumesComponent } from './components/volumes/volumes.component';
import { VesslNetworksService } from './services/vessl-networks.service';
import { NetworksComponent } from './components/networks/networks.component';
import { StatsPopupComponent } from './popups/stats-popup/stats-popup.component';
import { FilterPipe } from './pipes/filter.pipe';
import { NavMenuComponent } from './navigation/nav-menu/nav-menu.component';
import { LayoutModule } from '@angular/cdk/layout';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatSidenavModule } from '@angular/material/sidenav';



@NgModule({
  declarations: [
    AppComponent,
    AppsComponent,
    MessagePopupComponent,
    StatsComponent,
    AppRepositoryComponent,
    UsersComponent,
    SystemComponent,
    SettingsComponent,
    WaitPopupComponent,
    SysRestartPopupComponent,
    SysShutdownPopupComponent,
    LoginComponent,
    ImagesComponent,
    AppLauncherComponent,
    VolumesComponent,
    NetworksComponent,
    StatsPopupComponent,
    FilterPipe,
    NavMenuComponent,
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatCheckboxModule,
    MatDialogModule,
    MatToolbarModule,
    MatTabsModule,
    MatListModule,
    MatCardModule,
    MatMenuModule,
    MatSelectModule,
    MatProgressSpinnerModule,
    MatProgressBarModule,
    LayoutModule,
    MatSidenavModule,
    FlexLayoutModule
  ],
  providers: [CookieService, VesslContainersService, VesslSystemService, VesslUsersService,VesslImagesService,
    VesslVolumesService,VesslNetworksService,{
    provide: HTTP_INTERCEPTORS,
    useClass: TokenInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
