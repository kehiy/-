use std::{env, fs};

struct Config {
    query: String,
    file_name: String,
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let config: Config = Config::new(&args);

    println!("Searching for {:?}", config.query);
    println!("In: {:?}", config.file_name);

    let data = fs::read_to_string(config.file_name).expect("error while reading the file!");

    println!("file data:\n{}", data);
}

impl Config {
    fn new(args: &[String]) -> Config {
        let query = args[1].clone();
        let file_name = args[2].clone();

        Config { query, file_name }
    }
}
