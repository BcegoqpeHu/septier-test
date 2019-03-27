import {async, ComponentFixture, TestBed} from '@angular/core/testing';

import {ContentComponent} from './content.component';
import {MapComponent} from '../map/map.component';
import {NotAvailablePipe} from '../../pipes/not-available.pipe';
import {WebsocketModule} from '../../websocket';

describe('ContentComponent', () => {
  let component: ContentComponent;
  let fixture: ComponentFixture<ContentComponent>;

  beforeEach(async(() => {
    TestBed
      .configureTestingModule({
        imports: [
          WebsocketModule.config({
            path: '/ws'
          })
        ],
        declarations: [
          ContentComponent,
          MapComponent,
          NotAvailablePipe
        ]
      })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
