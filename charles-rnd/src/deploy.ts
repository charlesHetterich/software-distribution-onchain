import { createReviveSdk } from "@polkadot-api/sdk-ink";
import * as DESC from "@polkadot-api/descriptors";
import { API } from "apis";
import * as fs from "fs";
import * as os from "os";
import { Binary } from "polkadot-api";
import { blake2b256, sr25519 } from "@polkadot-labs/hdkd-helpers";
import { secretFromSeed } from "@scure/sr25519";

// import { getPolkadotSigner, PolkadotSigner } from "polkadot-api/signer";
import { DEPLOY } from "addresses";

// import { sr25519CreateDerive } from "@polkadot-labs/hdkd"
import {
    DEV_PHRASE,
    entropyToMiniSecret,
    mnemonicToEntropy,
    ss58Address,
} from "@polkadot-labs/hdkd-helpers";
import { getPolkadotSigner } from "polkadot-api/signer";

// 5Ck6mWifJmp9hdnAnSVTnJYM81BSdrotatKg4jPucVVgcfdz

function getSigner() {
    const mnemonic = JSON.parse(
        fs.readFileSync(os.homedir() + "/.polkadot/address.json", "utf8")
    ).secretPhrase;
    const entropy = mnemonicToEntropy(mnemonic);
    const miniSecret = entropyToMiniSecret(entropy);
    const privateKey = secretFromSeed(miniSecret);
    const publicKey = sr25519.getPublicKey(privateKey);
    console.log("Using address:", ss58Address(publicKey, 0));
    return getPolkadotSigner(publicKey, "Sr25519", (signerPayload) => {
        return sr25519.sign(signerPayload, privateKey);
    });
}

export async function deploy() {
    const codeBytes = fs.readFileSync(
        "test-contracts/target/ink/ink_library.polkavm"
    );
    const code = Binary.fromBytes(codeBytes); // what the SDK needs :contentReference[oaicite:1]{index=1}

    const mnemonic = JSON.parse(
        fs.readFileSync(os.homedir() + "/.polkadot/address.json", "utf8")
    ).secretPhrase;
    const signer = getSigner();

    // createReviveSdk(DESC.passethub, DESC.contracts);
    console.log(
        await API.paseoAssetHubSC.tx.Revive.upload_code({
            code: code,
            storage_deposit_limit: 100_000_000_000n,
        })
            .signAndSubmit(signer)
            .then((result) => {})
    );
}

// function codeHash() {
//     const bytes = fs.readFileSync(
//         "test-contracts/target/ink/ink_library.polkavm"
//     );
//     const codeHashHex =
//         "0x" + Buffer.from(blake2b256(bytes, { dkLen: 32 })).toString("hex");
//     const codeInfo = await API.paseoAssetHubSC.query.Revive.CodeInfoOf.getValue(
//         Binary.fromHex(codeHashHex)
//     );
//     console.log("CodeInfoOf:", codeInfo?.toJSON?.() ?? codeInfo);
// }

export async function validate() {
    const contractAddress = Binary.fromHex(
        "0x82b1349240d695b3edc35f9b251dcb38a61865eb86c7a858a74d5de51efb6ca1"
    );
    console.log("contract:", contractAddress.asHex());
    const code = await API.paseoAssetHubSC.query.Revive.CodeInfoOf.getValue(
        contractAddress
    );

    console.log("Code info:", code?.owner, code?.code_len);
}
