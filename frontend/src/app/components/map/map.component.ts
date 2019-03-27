/// <reference types="@types/googlemaps" />
import {Component, OnInit, ViewChild} from '@angular/core';
import {WebsocketService} from '../../websocket';
import {WsMessageModel} from '../../models/ws-message.model';
import {GpsExtModel} from '../../models/gps-ext.model';

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.styl']
})
export class MapComponent implements OnInit {

  @ViewChild('gmap') googleMapElement: any;
  map: google.maps.Map;
  marker: google.maps.Marker;

  constructor(private wsService: WebsocketService) {
  }

  ngOnInit() {
    this.initGoogleMap();
    this.wsService.on<WsMessageModel>('notification')
      .subscribe((notification) => this.refreshMap(notification.GpsExt));
  }

  private initGoogleMap() {
    const mapProp = {
      center: new google.maps.LatLng(53.195509, 45.012651),
      zoom: 16,
      mapTypeId: google.maps.MapTypeId.ROADMAP
    };
    this.map = new google.maps.Map(this.googleMapElement.nativeElement, mapProp);
  }

  private refreshMap(gpsExt: GpsExtModel) {
    const point = new google.maps.LatLng(gpsExt.Lat, gpsExt.Lon);
    this.refreshMarker(point);
    this.map.setCenter(point);
  }

  private refreshMarker(point: google.maps.LatLng) {
    if (this.marker) {
      this.marker.setMap(null);
      this.marker = undefined;
    }
    this.marker = new google.maps.Marker({
      position: point,
      map: this.map,
      title: 'Lat: ' + point.lat() + ', Lng: ' + point.lng()
    });
  }
}
