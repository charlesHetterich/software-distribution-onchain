// `dot` is the name we gave to `npx papi add`
import * as DESC from "@polkadot-api/descriptors";
import { para2107 } from "@polkadot-api/descriptors";
import { createClient } from "polkadot-api";
import { getInkClient } from "polkadot-api/ink";
import { createReviveSdk } from "@polkadot-api/sdk-ink";
import { getSmProvider } from "polkadot-api/sm-provider";
import * as SPECS from "polkadot-api/chains";
// import {para} from "polkadot-api/chains";
import { start } from "polkadot-api/smoldot";
import * as fs from "fs";

async function initAPIs() {
    const relayID = "ksmcc3";
    const paraIDs = ["ksmcc3_asset_hub"] as const;

    const smoldot = start();
    const relay = await smoldot.addChain({ chainSpec: SPECS.ksmcc3 });
    const relayClient = createClient(getSmProvider(relay));

    const AHChain = await smoldot.addChain({
        chainSpec: SPECS.ksmcc3_asset_hub,
        potentialRelayChains: [relay],
    });
    const AHClient = createClient(getSmProvider(AHChain));

    return {
        kusama: relayClient.getTypedApi(DESC.ksmcc3),
        kusamaAssetHub: AHClient.getTypedApi(DESC.ksmcc3_asset_hub),
        crust: null, // TODO!
    };
}

async function main() {
    const APIs = await initAPIs();

    APIs.kusama.event.System.Remarked.watch().subscribe((event) => {
        console.log("Kusama System.Remarked event:", event);
    });

    APIs.kusamaAssetHub.event.Revive.ContractEmitted.watch().subscribe(
        (event) => {
            console.log("Kusama Asset Hub ContractEmitted event:", event);
        }
    );

    // const parachains = await APIs.kusama.query.Paras.Heads.getEntries();
    const paraID = 2107;
    const chainhead = await APIs.kusama.query.Paras.Heads.getValue(paraID);
    const currentCodeHash =
        await APIs.kusama.query.Paras.CurrentCodeHash.getValue(paraID);
    const code = (await APIs.kusama.query.Paras.CodeByHash.getValue(
        currentCodeHash!
    ))!;

    console.log("head data:", chainhead?.asText());
    console.log("code hash:", currentCodeHash?.asHex());
    console.log("code: ", code.asText());

    await fs.mkdirSync(".cache/wasm", { recursive: true });
    await fs.writeFileSync(
        `.cache/wasm/kusama-para-${paraID}.wasm`,
        Buffer.from(code.asBytes())
    );
    console.log("Wrote wasm bytes:", code.asBytes().length);

    // const psp22Sdk = createReviveSdk(DESC.ksmcc3_asset_hub, DESC.contract);

    // const contractCodes =
    //     await APIs.kusamaAssetHub.query.Revive.PristineCode.getEntries();
    // console.log("Pristine contract codes:");
    // for (const code of contractCodes) {
    //     console.log(code.value.asText());
    // }
}

main().catch(console.error);
