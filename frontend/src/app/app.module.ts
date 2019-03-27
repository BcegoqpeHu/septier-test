import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {ContentComponent} from './components/content/content.component';
import {MapComponent} from './components/map/map.component';
import {WebsocketModule} from './websocket';
import {NotAvailablePipe} from './pipes/not-available.pipe';

@NgModule({
  declarations: [
    AppComponent,
    ContentComponent,
    MapComponent,
    NotAvailablePipe
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    WebsocketModule.config({
      path: '/ws'
    })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
