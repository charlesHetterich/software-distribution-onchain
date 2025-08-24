import { API } from "apis";
import * as fs from "fs";
import * as os from "os";

type AccountJSON = {
    accountId: string;
    networkId: string;
    publicKey: string;
    secretPhrase: string;
    secretSeed: string;
    ss58Address: string;
    ss58PublicKey: string;
};

class Account {
    account: AccountJSON;

    constructor() {
        this.account = JSON.parse(
            fs.readFileSync(os.homedir() + "/.polkadot/address.json", "utf-8")
        );
    }

    /**
     * Fetch the balance of this account on the Paseo Asset Hub
     */
    async balance() {
        const raw = await API.paseoAssetHub.query.System.Account.getValue(
            this.account.ss58Address
        );
        return Number(raw.data.free) / 10 ** 10;
    }
}
export const DEPLOY = new Account();
