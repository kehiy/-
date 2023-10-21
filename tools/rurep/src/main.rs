use std::{env, fs};

fn main() {
    let args: Vec<String> = env::args().collect();
    let query = &args[1];
    let file_name = &args[2];

    println!("Searching for {:?}", query);
    println!("In: {:?}", file_name);

    let data = fs::read_to_string(file_name).expect("error while reading the file!");

    println!("file data:\n{}", data);
}
