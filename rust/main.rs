use std::io;

fn main() {
    let mut numbers = Vec::new();

    loop {
        println!("Enter a number (or 'q' to quit):");

        let mut input = String::new();
        io::stdin().read_line(&mut input).expect("Failed to read line");

        if input.trim() == "q" {
            break;
        }

        match input.trim().parse::<i32>() {
            Ok(number) => numbers.push(number),
            Err(_) => {
                println!("Invalid input. Please enter a valid number.");
                continue;
            }
        }
    }

    let sum: i32 = numbers.iter().sum();

    println!("Sum: {}", sum);
}
