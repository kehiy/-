use std::{env, fs, process};

struct Config {
    query: String,
    file_name: String,
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let config = Config::new(&args).unwrap_or_else(|err| {
        println!("{}", err);
        process::exit(1);
    });

    println!("Searching for {:?}", config.query);
    println!("In: {:?}", config.file_name);

    let data = fs::read_to_string(config.file_name).expect("error while reading the file!");

    println!("file data:\n{}", data);
}

impl Config {
    fn new(args: &[String]) -> Result<Config, &str> {
        if args.len() < 3 {
            return Err("please provide both file name and query.");
        }

        let query = args[1].clone();
        let file_name = args[2].clone();

        Ok(Config { query, file_name })
    }
}
