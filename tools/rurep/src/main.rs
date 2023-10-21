use std::{env, error::Error, fs, process};

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

    if let Err(e) = run(config) {
        println!("{}", e);
        process::exit(1);
    }
}

fn run(config: Config) -> Result<(), Box<dyn Error>> {
    let data = fs::read_to_string(config.file_name)?;
    println!("file data:\n{}", data);
    println!("query:\n{}", config.query);

    Ok(())
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
