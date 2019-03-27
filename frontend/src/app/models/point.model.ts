import {AngleModel} from './angle.model';

export interface PointModel {
  Id: number;
  SettingsVersion: number;
  Channel: number;
  SNR: number;
  Angles: AngleModel[];
  RSSI: number;
  Angle2a: number;
  Antenna2a: number;
  Compression: number;
  Antenna: number;
  TSi: number;
  TSf: number;
  Clocks: number[];
}
