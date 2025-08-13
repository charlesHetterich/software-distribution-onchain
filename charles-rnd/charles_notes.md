## Monorepo Structure
- **[Rust]** Smart contract project
- **[Typescript]** Chain interactions project
    - Use `papi` to submit transactions as well as read from chains
    - (we need to find some metadata stuff on crust to figure out how to read from there)
- **[Go ?]** DRM software
    - I say `Go` just because its a pretty easy/fast language for compiling things into neat binaries
- We can also make separate r&d dirs to dump a lot of bullshit into

## Misc Priorities
- Figure out how to read a piece of data from Crust
- Figure out how to upload a piece of data to Crust
- Baseline sample contracts compiled, tested, & deployed to Paseo
- Baseline ability to create NFT set, mint NFTs
- Read NFT's, explore different kinds of keygen/clever cryptography we can do with that
- Scope out all specs of complete functional system. We can do this in parallel w/ above r&d, but will need to explore all of the above to be able to actually complete this part