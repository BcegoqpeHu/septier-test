import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'notAvailable'
})
export class NotAvailablePipe implements PipeTransform {

  transform(value: any, args?: any): any {
    return value ? value : 'N/A';
  }
}
