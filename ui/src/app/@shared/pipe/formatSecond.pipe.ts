import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'formatSecond'
})
export class FormatSecondPipe implements PipeTransform {

  transform(value: number): string {
    const d = Math.floor(value / 86400);
    let remainderDay =  value % (86400)
    const h = Math.floor(remainderDay / 3600);
    let remainderHour =  remainderDay % (3600)
    const m = Math.floor((remainderHour/60));
    let s =  remainderHour % (60)
    // const s = Math.floor((remainderMinute / 60));
    // let remainderSecond =  remainderHour % (60)
    return (d > 0 ? `${d}天` : '') + (h > 0 ? `${h}时` : '') + (m > 0 ? `${m}分` : '') + (s > 0 ? `${s}秒` : '');
  }
}