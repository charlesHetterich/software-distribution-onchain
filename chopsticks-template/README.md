## Chopsticks Template

This is a template for simulating forked Kusama Asset Hub and Crust Shadow networks, primarily for testing purposes.

### Setup

```bash
npm install
```

### Run Networks

Run the following command to start the networks with XCM support:

```bash
npx @acala-network/chopsticks xcm \
--r kusama \
--p kusama-asset-hub \
--p configs/crust-shadow.yml
```


### Access Networks via Explorer

After the networks are running, the ports for each network will be shown in the terminal. Use these ports to access the networks via the PolkadotJS Explorer.

```bash
[14:54:49.816] INFO: Loading config file https://raw.githubusercontent.com/AcalaNetwork/chopsticks/master/configs/kusama-asset-hub.yml
    app: "chopsticks"
[14:55:39.673] INFO: Kusama Asset Hub RPC listening on http://[::]:8000 and ws://[::]:8000
    app: "chopsticks"
        chopsticks::executor  TRACE: [1] Calling Metadata_metadata
        chopsticks::executor  TRACE: [1] Completed Metadata_metadata
[14:55:41.171] INFO: Crust Shadow RPC listening on http://[::]:8001 and ws://[::]:8001
```

From the above output, we can see that the Kusama Asset Hub is running on port 8000 and the Crust Shadow is running on port 8001.