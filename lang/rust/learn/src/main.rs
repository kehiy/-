
const BEST_NUMBER: i8 = 7;
const PI: f32 = 3.14;

fn main() {
    println!("hello rust!!!");

    let age: i8 = 15;
    println!("this is the age: {}", age);

    let name: &str = "Kayhan";
    println!("this is the name: {}", name);

    // this is a single line comment

    /*
    this is a multi line comment
    */

    let mut will_change: i8 = 10;
    println!("this will change: {}", will_change);

    will_change = 100;
    println!("changed: {}", will_change);

    let is_valid: bool = true;
    println!("is valid? {}", is_valid);

    let fav_char: char = 'k';
    println!("I'm {}, and my fav char is: {}", name, fav_char);

    let(country, capital) = ("iran", "tehran");
    println!("capital of {} is {}.", country, capital);

    // math
    println!("7 + 7 = {}", 7 + 7);
    println!("7 - 7 = {}", 7 - 7);
    println!("7 * 7 = {}", 7 * 7);
    println!("7 / 7 = {}", 7 / 7);
    println!("7 % 7 = {}", 7 % 7);
    // yes I love 7 number.

    // let pi: f32 = 3.14;
    println!("pi number i: {}", PI);

    let _ignore: i8 = 7; // that's just a test, never ignore 7 in real life.

    println!("the best number is {}", BEST_NUMBER);

    // let x: i8 = 10;
    // let x: i8 = 21;
    let x: i8 = 7;
    if x == 7 {
        println!("x is seven");
    } else if x <= 20 {
        println!("x is <= 20");
    } else {
        println!("x is not <= 20 and 7");
    }

    // loop

    let mut n: i8 = 0;
    loop {
        n += 1;
        if n == 11 {
            continue;
        }

        println!("n is {}", n);
        if n == 14 {
            break;
        }
    }
}
