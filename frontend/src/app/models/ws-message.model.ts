import {GpsExtModel} from './gps-ext.model';
import {PointModel} from './point.model';

export interface WsMessageModel {
  GpsExt: GpsExtModel;
  Points: PointModel[];
}
