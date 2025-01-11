// Prevent console window in addition to Slint window in Windows release builds when, e.g., starting the app via file manager. Ignored on other platforms.
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use std::error::Error;

use nostr::prelude::*;

slint::include_modules!();

fn decode(entity: &str) -> String {
    let mut result: String = "".to_string();

    let typ = Nip19::from_bech32(entity).unwrap_or(Nip19::EventId(EventId::all_zeros()));

    if let Nip19::Pubkey(data) = typ {
        result = data.to_hex();
    }

    if let Nip19::Profile(data) = &typ {
        result = format!(
            "pubkey: {}\nrelays: {:?}",
            data.public_key.to_hex(),
            data.relays
        );
    }

    if let Nip19::Event(data) = &typ {
        result = format!(
            "author: {}\nrelays: {:?}\nid: {}",
            data.author.unwrap().to_hex(),
            data.relays,
            data.event_id.to_hex()
        );
    }

    if let Nip19::EventId(data) = &typ {
        result = data.to_hex();
    }

    if let Nip19::Secret(data) = &typ {
        result = data.to_secret_hex();
    }

    if let Nip19::Coordinate(data) = &typ {
        result = format!(
            "pubkey: {}\nrelays: {:?}",
            data.public_key.to_hex(),
            data.relays
        );
    }

    // todo::: this is the most stupid way to prevent crash!
    if result == "0000000000000000000000000000000000000000000000000000000000000000" {
        return "invalid nip-19 entity!".to_string();
    }

    result
}

fn main() -> Result<(), Box<dyn Error>> {
    let ui = AppWindow::new()?;

    ui.on_decode({
        let ui_handle = ui.as_weak();
        move |entity| {
            let ui = ui_handle.unwrap();
            let result = decode(entity.as_str());
            ui.set_results(result.into());
        }
    });

    ui.run()?;

    Ok(())
}
