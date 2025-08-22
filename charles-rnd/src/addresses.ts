import { API } from "apis";
import * as fs from "fs";
import * as os from "os";

type Account = {
    accountId: string;
    networkId: string;
    publicKey: string;
    secretPhrase: string;
    secretSeed: string;
    ss58Address: string;
    ss58PublicKey: string;
};

const DEPLOY: Account = JSON.parse(
    fs.readFileSync(os.homedir() + "/.polkadot/address.json", "utf-8")
);
console.log("Using address:", DEPLOY.secretSeed);

console.log(
    "Balance: ",
    await API.paseoAssetHub.query.Balances.Account.getValue(DEPLOY.ss58Address)
);

console.log(
    "get balance at asset hub faucet: https://faucet.polkadot.io/?parachain=1000"
);

export { DEPLOY };

// 146E1z36kcMFQX8Nh8LzRRgoJ4fUaJCTFkVgfbY1qefBgkyG
// 1DdqDYEAC3RofEDPwBARLTPSqpV6tUW8Gh7b7HzjVm1LdQDC
// 0x022ee9f234f3b840de499b0f7aeec4424f11834d3ebb2735e3c1bf85a3483bff5b
// 0x88b0bf113f2c2de69e1d6cc1656c72df4679596fba23bc6c01179dddd0d60d35
