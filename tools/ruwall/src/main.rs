use bdk::bitcoin::Network;
use bdk::database::MemoryDatabase;
use bdk::{SyncOptions, Wallet};
use std::env;
use bdk::blockchain::ElectrumBlockchain;
use bdk::electrum_client::Client;
use bdk::wallet::AddressIndex::New;


fn main() -> anyhow::Result<()> {
    dotenv::from_filename(".env").ok();
    let descriptor = env::var("WALLET_DESC").unwrap();

    let client = Client::new("ssl://electrum.blockstream.info:60002")?;
    let blockchain = ElectrumBlockchain::from(client);

    let wallet = Wallet::new(
        &descriptor.clone(),
        None,
        Network::Testnet,
        MemoryDatabase::default(),
    )?;

    wallet.sync(&blockchain, SyncOptions::default())?;

    println!("Descriptor balance: {} SAT", wallet.get_balance()?);
    println!("Descriptor address: {}", wallet.get_address(New)?);

    Ok(())
}
