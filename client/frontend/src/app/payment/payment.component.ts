import { Component } from '@angular/core';
import { MetaMaskSDK } from "@metamask/sdk"
import {MatButtonModule} from '@angular/material/button';

const MMSDK = new MetaMaskSDK({
  dappMetadata: {
    name: "Example JavaScript Dapp",
    url: window.location.href,
  },
})

@Component({
  selector: 'app-payment',
  imports: [MatButtonModule],
  templateUrl: './payment.component.html',
  styleUrls: ['./payment.component.css']
})


export class PaymentComponent {

  viewResult: any | undefined;

  public async signEIP712() {

    let ethereum = MMSDK.getProvider()

    const msgParams = JSON.stringify({
      domain: {
        // This defines the network, in this case, Mainnet.
        chainId: 84532,
        // Give a user-friendly name to the specific contract you're signing for.
        name: "Ether Mail",
        // Add a verifying contract to make sure you're establishing contracts with the proper entity.
        verifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
        // This identifies the latest version.
        version: "1",
      },
      message: {
        contents: "Hello, Bob!",
        attachedMoneyInEth: 4.2,
        from: {
          name: "Cow",
          wallets: [
            "0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826",
            "0xDeaDbeefdEAdbeefdEadbEEFdeadbeEFdEaDbeeF",
          ],
        },
        to: [
          {
            name: "Bob",
            wallets: [
              "0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB",
              "0xB0BdaBea57B0BDABeA57b0bdABEA57b0BDabEa57",
              "0xB0B0b0b0b0b0B000000000000000000000000000",
            ],
          },
        ],
      },
      // This refers to the keys of the following types object.
      primaryType: "Mail",
      types: {
        // This refers to the domain the contract is hosted on.
        EIP712Domain: [
          { name: "name", type: "string" },
          { name: "version", type: "string" },
          { name: "chainId", type: "uint256" },
          { name: "verifyingContract", type: "address" },
        ],
        // Not an EIP712Domain definition.
        Group: [
          { name: "name", type: "string" },
          { name: "members", type: "Person[]" },
        ],
        // Refer to primaryType.
        Mail: [
          { name: "from", type: "Person" },
          { name: "to", type: "Person[]" },
          { name: "contents", type: "string" },
        ],
        // Not an EIP712Domain definition.
        Person: [
          { name: "name", type: "string" },
          { name: "wallets", type: "address[]" },
        ],
      },
    })
    var accounts = await MMSDK.connect();
    var from = accounts[0];
    var params = [from, msgParams]
    console.log(params)
    var method = "eth_signTypedData_v4"

    // Or window.ethereum if you don't support EIP-6963.
    if(ethereum) {
      try {
        let result: any = await ethereum.request({
          method,
          params
        })
        if(result) {
          this.viewResult = result;
        }
      }
      catch (error: any) {
        this.viewResult = error.message;
      }
    }
  }
}
