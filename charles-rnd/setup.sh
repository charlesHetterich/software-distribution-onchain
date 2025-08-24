# install subkey
cargo install subkey

pnpm install

# Add PAPI descriptors
npx papi add ksmcc3 -n ksmcc3
npx papi add ksmcc3_asset_hub -n ksmcc3_asset_hub
npx papi add paseo -n paseo
npx papi add paseo_asset_hub -n paseo_asset_hub
npx papi add passethub -w wss://passet-hub-paseo.ibp.network # Development SC Paseo Assethub

# Crust web sockets
# wss://rpc.crust.network
# wss://rpc.crustnetwork.xyz
# wss://rpc.crustnetwork.cc
#
# Shadow network (Kusama)
# wss://rpc-shadow.crust.network/
# 
# Testnet
# wss://rpc-rocky.crust.network
# npx papi add crust -w wss://rpc.crust.network