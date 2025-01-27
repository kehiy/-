// Prevent console window in addition to Slint window in Windows release builds when, e.g., starting the app via file manager. Ignored on other platforms.
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use nostr::prelude::*;
use std::error::Error;

slint::include_modules!();

fn generate() -> (String, String) {
    let keys = Keys::generate();
    (keys.public_key.to_hex(), keys.secret_key().to_secret_hex())
}

fn main() -> Result<(), Box<dyn Error>> {
    let ui = AppWindow::new()?;

    ui.on_generate_key({
        let ui_handle = ui.as_weak();
        move || {
            let ui = ui_handle.unwrap();
            let (pubkey, privkey) = generate();

            ui.set_privkey(privkey.into());
            ui.set_pubkey(pubkey.into());
        }
    });

    ui.on_clean({
        let ui_handle = ui.as_weak();
        move || {
            let ui = ui_handle.unwrap();
            ui.set_privkey("".into());
            ui.set_pubkey("".into());
        }
    });

    ui.run()?;

    Ok(())
}
