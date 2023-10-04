use std::process;
use std::time::{SystemTime, UNIX_EPOCH};

fn main() {
    println!("{}", random(gen_seed()));
}

fn gen_seed() -> i32 {
    let pid = process::id();
    let time = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    let time = time.as_nanos() as i32;
    let i32pid = pid as i32;
    let seed = time + i32pid % i32pid - 100000;
    return seed;
}

fn random(seed: i32) -> i64 {
    let random = seed / 7 * 2;
    return random.into();
}
