#![cfg_attr(not(feature = "std"), no_std, no_main)]

#[ink::contract]
mod ink_library {
    use ink::{
        prelude::vec::Vec,
        xcm::{
            latest::{send_xcm, AssetTransferFilter, SendError},
            prelude::*,
            DoubleEncoded,
        },
    };

    const ASSET_HUB_PARA: u32 = 1000; // Paseo Asset Hub (para id)

    #[ink(storage)]
    pub struct InkLibrary {}

    type Bytes32 = [u8; 32]; //ink::sol::FixedBytes<32>;

    impl InkLibrary {
        /// Constructor that initializes the `bool` value to the given `init_value`.
        #[ink(constructor)]
        pub fn new() -> Self {
            Self {}
        }

        fn xcm_transact_inner(
            as_account: Option<[u8; 32]>,
            call: Vec<u8>,
            ref_time: u64,
            proof_size: u64,
            fee: u128,
        ) -> Result<(), SendError> {
            let dest = Location::new(1, [Parachain(ASSET_HUB_PARA)]);
            let w = Weight::from_parts(ref_time, proof_size);
            let fee_asset = (Here, fee);

            // Build XCM program
            let mut instrs: Vec<Instruction<()>> = Vec::new();

            // Pay for execution on Asset Hub
            instrs.push(WithdrawAsset(fee_asset.clone().into()));
            instrs.push(BuyExecution {
                fees: fee_asset.into(),
                weight_limit: WeightLimit::Limited(w),
            });

            // Optionally change origin to a specific AccountId32 on Asset Hub
            if let Some(id32) = as_account {
                instrs.push(DescendOrigin(
                    // Origin becomes this AccountId32 on the destination
                    Junctions::X1(
                        [AccountId32 {
                            network: None,
                            id: id32,
                        }]
                        .into(),
                    ),
                ));
                // With DescendOrigin, the origin kind for Transact is the XCM origin:
                instrs.push(Transact {
                    origin_kind: OriginKind::Xcm,
                    fallback_max_weight: w.into(),
                    call: call.into(),
                });
            } else {
                // Execute as the parachain *sovereign account* on Asset Hub
                instrs.push(Transact {
                    origin_kind: OriginKind::SovereignAccount,
                    fallback_max_weight: w.into(),
                    call: call.into(),
                });
            }

            let xcm = Xcm::<()>(instrs);
            send_xcm::<()>(dest, xcm).map(|_| ())
        }

        #[ink(message, payable)]
        pub fn teleport(
            &mut self,
            para_id: u32,
            beneficiary: Bytes32,
            amount: u128,
        ) -> Result<(), SendError> {
            let destination = Location::new(1, [Parachain(para_id)]);
            let remote_fees =
                AssetTransferFilter::Teleport(Definite((Parent, amount.saturating_div(10)).into()));
            let preserve_origin = false;
            let mut transfer_assets = Vec::new();
            transfer_assets.push(AssetTransferFilter::Teleport(Wild(AllCounted(1))));
            let remote_xcm = Xcm::<()>::builder_unsafe()
                .deposit_asset(AllCounted(1), beneficiary)
                .build();

            let xcm = Xcm::<()>::builder()
                .withdraw_asset((Parent, amount))
                .pay_fees((Parent, amount.saturating_div(10)))
                .initiate_transfer(
                    destination.clone(),
                    remote_fees,
                    preserve_origin,
                    transfer_assets,
                    remote_xcm,
                )
                .build();

            send_xcm::<()>(destination, xcm).map(|_| ())
        }
    }
}
