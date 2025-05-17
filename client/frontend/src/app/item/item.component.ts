import { Component } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { PaymentComponent } from '../payment/payment.component';

export interface Tile {
  color: string;
  cols: number;
  rows: number;
  text: string;
}

@Component({
  selector: 'app-item',
  imports: [MatGridListModule, PaymentComponent],
  templateUrl: './item.component.html',
  styleUrls: ['./item.component.css']
})
export class ItemComponent {

  tiles: Tile[] = [
    {text: 'One', cols: 1, rows: 1, color: 'white'},
  ];

}
