import {Observable} from 'rxjs';

export interface IWebsocketService {
  status: Observable<boolean>;

  on<T>(event: string): Observable<T>;

  send(event: string, data: any): void;
}

export interface WebSocketConfig {
  path: string;
  reconnectInterval?: number;
  reconnectAttempts?: number;
}

export interface IWsMessage<T> {
  event: string;
  data: T;
}
