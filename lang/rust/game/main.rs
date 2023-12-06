use std::{cmp::Ordering, io};
use rand::Rng;
use colored::*;

fn main() {
    let number: i32 = rand::thread_rng().gen_range(1, 7);
    println!("{}", "input number:".white());

    let mut attempt: isize = 0;
    loop {
        let mut guess: String = String::new();
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line.");

        let guess: i32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("{}", "Please enter a number!".red());
                continue;
            },
        };

        match guess.cmp(&number) {
            Ordering::Less => {
                println!("{}", "HINT: too small!".red());
                attempt = attempt + 1;
            },
            Ordering::Greater => {
                println!("{}", "HINT: too big!".red());
                attempt = attempt + 1;
            },
            Ordering::Equal => {
                let  msg: String = format!("you won with {} attempts!", attempt);
                println!("{}", msg.green());
                break;
            }
        };
    };
}
