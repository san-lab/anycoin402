Some notes on the "permit" scheme

1. We need the facilitator to be gven the allowance. We need to pass the facilitator's address in PaymentRequirements.
   For now I will put it in the "Extra" map under the "facilitator" key as a hex string

2. Let us try to off-load as much of unmarshalling chores as possible to json. Hence the "Permit" type, that carries the
   whole context.

3. The Permit being issued to the facilitator is an obvious security concern for the merch

4. IT could be worth our while to use the SechemeOutput facility to redefin the return message.
