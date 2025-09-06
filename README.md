# Software Distribution & Digital Rights Management (DRM) on-chain.

This would be a replacement for things like the <AppStore, Google Play Store, Steam, etc>. Our intention currently has been to focus on games (Steam), but there is no reason this cannot extend to general software distribution.

The main service these platforms provide you, a a game publisher, in exchange for usually ~30% of all sales is: they ensure that (1) only people who buy your <game> can play it, and (2) someone who buys it cannot *easily* redistribute it to to the rest of the world for free.

### How do we convert this Web2 service into a Web3 service?
#### (1) NFT collections on AssetHub
- When a game publisher `publish(...)`'s a game, a new *NFT collection* is created
- When a player `buy(...)`'s a game, a new *NFT* is *minted* from that game's NFT collection. This NFT is affectively the players access key to this software

#### (2) Open-Source DRM (off-chain)
- Here's a simplified TL;DR of what DRM is and how it works in Web2.
    - a) I am a game publisher, and I compile my game into an executable
    - b) I give my game executable to steam, and they wrap my executable in some logic which checks their database "did <player> buy <game>? if yes, run the game." creating a new *wrapped game executable*
    - c) Steam then obfuscates this executable, making it difficult to reverse engineer & extract the core game (basically, just jumble up the code a ton without changing its logic)
    - d) Now, anyone in the world can download this executable, yet only players who have bought <game> through steam can play <game>
- A Web3 solution does exactly the same thing except
    - Step (b) checks AssetHub NFT's instead of a Steam API
    - The game publisher runs steps (b) & (c) themselves

#### (3) Decentralized storage on Crust
- When a game publisher `publish(...)`'s a game, instead of publishing to a centralized database, it is stored on Crust for anyone to download. (But only *owners* can actually play it)

## Motivation

- For publishers: remove middlemen (Steam, AppStore, Google Play Store, etc) and facilitate fairer distribution.
- For players: in some instances, avoid having to lose their access if somehow the platform shuts down / goes out of business.
- More importantly, provide an improved ownership and consumer rights system.

### Security Risk Assessment

![Risk Assessment Comparison](/static/bpmn.svg)

## Quick Start

TODO

## Limitations and Future Development

- We intentionally do not allow for the transfer of ownership of the software (represented by an NFT) to primarily give publishers more control over their software's pricing, and not to lose out on potential revenue because of resale.
- Therefore, users must understand that losing their private keys means losing access to the software.
- There are ways to provide a recovery mechanism for users (e.g. [Kusama Social Recovery](https://wiki.polkadot.com/kusama/kusama-social-recovery/)), but it is out of scope for this project.

## Resources / Papers
**Crust (storage)**
- [Crust Economy Paper](https://crust.network/download/ecowhitepaper_en.pdf)
- [Crust Tech Paper](https://ipfsgw.live/ipfs/QmP9WqDYhreSuv5KJWzWVKZXJ4hc7y9fUdwC4u23SmqL6t)
    - *if you can find a better link... this one hosted on IPFS doesn't seem to be loading (not a great sign LOL)*
- [XCM on crust](https://wiki.crust.network/docs/en/buildXCMPBasedCrossChainSolution)

**Ink! Contracts**
- [Ink! w/ hardhat guide](https://use.ink/tutorials/ethereum-compatibility/hardhat-deployment/)
- [Setup environment for Ink! (not EVM)](https://use.ink/docs/v6/getting-started/setup/)

**DRM Papers**
- **2024 |** [Secure Rights (DRM on blockchain)](https://arxiv.org/abs/2403.06094)
    - doesn't mention NFT's ?
- **2018 |** [Secure DRM Scheme Based on Blockchain with
High Credibility](https://ietresearch.onlinelibrary.wiley.com/doi/pdf/10.1049/cje.2018.07.003)

**Obfuscation / Packing Papers**
- [UPX (packing)](https://github.com/upx/upx?tab=readme-ov-file)
