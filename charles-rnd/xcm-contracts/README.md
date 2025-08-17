### Contracts R&D

To compile contracts
```
cargo contract build --release
```

To deploy a contract (local test node), first launch the test node in a separate terminal with `./ink-node --dev --tmp`. Then deploy with
```
cargo contract upload --url ws://127.0.0.1:9944 --suri //Alice -x
```

Instantiate the contract (not sure what that means tbh... we just uploaded it ?)
```
export CONTRACT_ADDRESS=$(
  cargo contract instantiate -y \
    --constructor new \
    --url ws://127.0.0.1:9944 \
    --suri //Alice \
    --salt 0x00 \
    -x 2>&1 \
  | tee /dev/tty \
  | sed -nE 's/^[[:space:]]*Contract (0x[0-9a-fA-F]{40})$/\1/p' \
  | tail -1
)
```

Finally, you can call to the contract
```
cargo contract call \
  --contract $CONTRACT_ADDRESS \
  --message teleport \
  --args 2000 \
        0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d \
        1000000000000 \
  --suri //Alice \
  --url ws://127.0.0.1:9944 \
  --skip-dry-run -y
```

Because we're trying to run XCM on a local node which does not have any parachains connected, we'll get errors— you should see something like
```
thread 'main' panicked at crates/cargo-contract/src/cmd/call.rs:200:25:
Call did revert "\u{1}\0"
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
```