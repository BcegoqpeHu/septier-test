import {TestBed} from '@angular/core/testing';

import {WebsocketService} from './websocket.service';
import {WebsocketModule} from './websocket.module';

describe('WebsocketService', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [
      WebsocketModule.config({
        path: '/ws'
      })
    ]
  }));

  it('should be created', () => {
    const service: WebsocketService = TestBed.get(WebsocketService);
    expect(service).toBeTruthy();
  });
});
