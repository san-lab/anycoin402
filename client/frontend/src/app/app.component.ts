import {Component} from '@angular/core';
import { PaymentComponent } from './payment/payment.component';
import { ShopComponent } from './shop/shop.component';

@Component({
  selector: 'app-root',
  imports: [PaymentComponent, ShopComponent],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'default';
}
