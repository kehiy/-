use std::{error::Error, fs};

pub struct Config {
    pub query: String,
    pub file_name: String,
}

pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    let data = fs::read_to_string(config.file_name)?;

    for line in search(&config.query, &data) {
        println!("{}", line);
    }

    Ok(())
}

impl Config {
    pub fn new(args: &[String]) -> Result<Config, &str> {
        if args.len() < 3 {
            return Err("please provide both file name and query.");
        }

        let query = args[1].clone();
        let file_name = args[2].clone();

        Ok(Config { query, file_name })
    }
}

pub fn search<'a>(query: &str, data: &'a str) -> Vec<&'a str> {
    let mut results = Vec::new();

    for line in data.lines() {
        if line.contains(query) {
            results.push(line);
        }
    }

    results
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one_result() {
        let query = "hello";
        let data = "
hello, how are you?
are you ok?
or not?
jkdfhdsifhd
some text for test hel1o !
some text for test hel!o!
            ";

        assert_eq!(vec!["hello, how are you?"], search(query, data));
    }
}
