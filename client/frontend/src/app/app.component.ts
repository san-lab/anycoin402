import {Component} from '@angular/core';
import { PaymentComponent } from './payment/payment.component';

@Component({
  selector: 'app-root',
  imports: [PaymentComponent],
  template: `
  <main>
      <header class="brand-name">
        <h1> x402 Research </h1>
      </header>
      <section class="content">
        <app-payment> </app-payment>
      </section>
    </main>
  `,
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'default';
}
