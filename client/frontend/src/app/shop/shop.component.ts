import { Component } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { ItemComponent } from '../item/item.component';

@Component({
  selector: 'app-shop',
  imports: [MatGridListModule, ItemComponent],
  templateUrl: './shop.component.html',
  styleUrls: ['./shop.component.css']
})
export class ShopComponent {

}
