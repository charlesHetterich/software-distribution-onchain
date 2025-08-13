### Contracts R&D

To compile contracts
```
cargo contract build --release```

To deploy a contract (local test node), first launch the test node in a separate terminal with `./ink-node --dev --tmp`. Then deploy with
```
cargo contract upload --url ws://127.0.0.1:9944 --suri //Alice
```

Instantiate the contract (not sure what that means tbh... we just uploaded it ?)
```
cargo contract instantiate \
  --constructor new \
  --url ws://127.0.0.1:9944 \
  --suri //Alice \
  --salt 0x00 \
  -x
```

Finally, you can call to the contract
```
# TODO! figure out a functional contract call to test out
cargo contract call \
  --message flip \
  --url ws://127.0.0.1:9944 \
  --suri //Alice \
  --contract <CONTRACT_ADDRESS>


cargo contract call \
  --contract <CONTRACT_ADDRESS> \
  --message teleport \
  --args <para_id> 0x<32-bytes-beneficiary> <amount> \
  --suri //Alice \
  --url ws://127.0.0.1:9944 \
  -x

```

