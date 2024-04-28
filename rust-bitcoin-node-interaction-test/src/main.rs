use std::str::FromStr;

use bitcoincore_rpc::{bitcoin::Address, Auth, Client, RpcApi};

// Just an example on how to go through each block (for indexing purposes)
fn main() -> anyhow::Result<()> {
    let rpc = Client::new(
        "localhost:8332",
        Auth::UserPass("bitcoinduser".to_owned(), "bitcoinduser".to_owned()),
    )?;

    let blockchain_info = rpc.get_blockchain_info()?;
    dbg!(blockchain_info);

    match rpc.get_wallet_info() {
        Ok(s) => println!("Wallet info: {:?}", s),
        Err(_) => {
            println!("Wallet not found, creating...");
            let new_wallet = rpc.create_wallet(
                "my_wallet_tracker_wallet_lol",
                Some(true),
                Some(true),
                None,
                Some(false),
            )?;
            println!("Created new wallet: {:?}", new_wallet)
        }
    };

    let addr_info = rpc.get_address_info(
        &Address::from_str("bc1q5emlck4402nlctp3ss7cewkc4nsrewwujh0jd7")?.assume_checked(), // random wallet I grabbed from the internet just to cross-reference information on the internet with my node
    )?;
    dbg!(addr_info);

    // let a = rpc.get_block(&rpc.get_block_hash(0)?)?;
    // dbg!(a.header.target());

    let count = rpc.get_block_count()?;
    for i in 0..count {
        // let _block = rpc.get_block(&rpc.get_block_hash(i)?);
        rpc.get_block(&rpc.get_block_hash(i)?)?;
        println!("{i}");
    }

    // let latest = rpc.get_block(&rpc.get_block_hash(count)?)?;
    // dbg!(latest.);

    Ok(())
}
