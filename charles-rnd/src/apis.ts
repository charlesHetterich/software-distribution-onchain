// `dot` is the name we gave to `npx papi add`
import * as DESC from "@polkadot-api/descriptors";
import { para2107 } from "@polkadot-api/descriptors";
import { createClient } from "polkadot-api";
import { getInkClient } from "polkadot-api/ink";
import { createReviveSdk } from "@polkadot-api/sdk-ink";
import { getSmProvider } from "polkadot-api/sm-provider";
import { getWsProvider } from "polkadot-api/ws-provider/web";
import { withPolkadotSdkCompat } from "polkadot-api/polkadot-sdk-compat";

import * as SPECS from "polkadot-api/chains";
import { start } from "polkadot-api/smoldot";

const DEBUG = process.env.DEBUG === "true";
const relayIDs = ["ksmcc3", "paseo"] as const;
const paraIDs = ["ksmcc3_asset_hub", "paseo_asset_hub"] as const;

async function initAPIs() {
    const smoldot = start({ logCallback: DEBUG ? undefined : () => {} });
    const kusama = await smoldot.addChain({ chainSpec: SPECS.ksmcc3 });
    const paseo = await smoldot.addChain({ chainSpec: SPECS.paseo });

    const kusamaAssetHub = await smoldot.addChain({
        chainSpec: SPECS.ksmcc3_asset_hub,
        potentialRelayChains: [kusama, paseo],
    });
    const paseoAssetHub = await smoldot.addChain({
        chainSpec: SPECS.paseo_asset_hub,
        potentialRelayChains: [kusama, paseo],
    });

    const paseoAssetHubSC = createClient(
        withPolkadotSdkCompat(
            getWsProvider("wss://passet-hub-paseo.ibp.network")
        )
    );

    return {
        kusama: createClient(getSmProvider(kusama)).getTypedApi(DESC.ksmcc3),
        kusamaAssetHub: createClient(getSmProvider(kusamaAssetHub)).getTypedApi(
            DESC.ksmcc3_asset_hub
        ),
        paseo: createClient(getSmProvider(paseo)).getTypedApi(DESC.paseo),
        paseoAssetHub: createClient(getSmProvider(paseoAssetHub)).getTypedApi(
            DESC.paseo_asset_hub
        ),
        paseoAssetHubSC: paseoAssetHubSC.getTypedApi(DESC.passethub),
        crust: null, // TODO!
    };
}

export const API = await initAPIs();
