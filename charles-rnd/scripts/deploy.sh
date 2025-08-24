# Deploy contracts in `test-contracts` to Paseo Asset Hub

SURI="$(jq -r .secretPhrase ~/.polkadot/address.json)"
PUBKEY="$(jq -r .ss58Address ~/.polkadot/address.json)"
# WS=wss://testnet-passet-hub.polkadot.io
WS=wss://testnet-passet-hub.polkadot.io

npx @polkadot/api-cli query.system.account $PUBKEY
  --ws wss://passet-hub-paseo.ibp.network

cd test-contracts

cargo contract upload \
  --url $WS \
  --suri "$SURI" \
  --storage-deposit-limit 1000000000 \
  -x

ADDR=$(cargo contract instantiate -y \
  --constructor new \
  --url $WS \
  --suri "$SURI" \
  --salt 0x00 -x \
  2>&1 | tee /dev/tty \
  | sed -nE 's/.*Contract (0x[0-9a-fA-F]{40}).*/\1/p' | tail -1)

echo "Contract address: $ADDR"
