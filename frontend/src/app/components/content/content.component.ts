import {Component, OnInit} from '@angular/core';
import {WebsocketService} from '../../websocket';
import {WsMessageModel} from '../../models/ws-message.model';
import {PointModel} from '../../models/point.model';

@Component({
  selector: 'app-content',
  templateUrl: './content.component.html',
  styleUrls: ['./content.component.styl']
})
export class ContentComponent implements OnInit {

  points: PointModel[] = new Array({} as PointModel);

  constructor(private wsService: WebsocketService) {
  }

  ngOnInit() {
    this.wsService.on<WsMessageModel>('notification')
      .subscribe((notification) => this.points = notification.Points);
  }
}
