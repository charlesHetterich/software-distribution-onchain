import { API } from "apis";
import { DEPLOY } from "addresses";
import * as fs from "fs";
import chalk from "chalk";

const link = (text: string, url: string) =>
    `\u001B]8;;${url}\u0007${text}\u001B]8;;\u0007`; // OSC-8

async function main() {
    console.log("Using address:", chalk.bold(DEPLOY.account.ss58Address));
    console.log(
        "Get balance from the",
        chalk.blue.bold(
            link("Paseo AssetHub faucet", "https://faucet.polkadot.io/")
        )
    );

    console.log("Balance: ", chalk.green(await DEPLOY.balance()));
}
main().catch(console.error);

async function _main() {
    // console.log("Using address:", DEPLOY.secretSeed);
    API.kusama.event.System.Remarked.watch().subscribe((event) => {
        console.log("Kusama System.Remarked event:", event);
    });

    API.kusamaAssetHub.event.Revive.ContractEmitted.watch().subscribe(
        (event) => {
            console.log("Kusama Asset Hub ContractEmitted event:", event);
        }
    );

    // const parachains = await APIs.kusama.query.Paras.Heads.getEntries();
    const paraID = 2107;
    const chainhead = await API.kusama.query.Paras.Heads.getValue(paraID);
    const currentCodeHash =
        await API.kusama.query.Paras.CurrentCodeHash.getValue(paraID);
    const code = (await API.kusama.query.Paras.CodeByHash.getValue(
        currentCodeHash!
    ))!;

    console.log("head data:", chainhead?.asText());
    console.log("code hash:", currentCodeHash?.asHex());
    // console.log("code: ", code.asText());

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
