
const BEST_NUMBER: i8 = 7;
const _MUST_USELESS_NUMBER: i8 = 11;
const PI: f32 = 3.14;

enum ip_addr_versions {
    V4,
    V6,
}

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
        if n == 11 { // I hate 11
            continue;
        }

        println!("n is {}", n);
        if n == 14 {
            break;
        }
    }

    for i in 0..BEST_NUMBER + 1 {
        println!("{}", i);
    }

    println!("----------------------------------------");

    let r = 1..6;
    for i in r {
        println!("{}", i);
    }

    println!("----------------------------------------");

    let fruits = vec!["apple", "orange", "mango"];
    for (i, fruit) in fruits.iter().enumerate() {
        println!("{}-{}", i, fruit);
    }

    println!("----------------------------------------");

    let mut number: i8 = 1;
    while number <= 10 {
        println!("{}", number);
        number += 1;
    }

    println!("----------------------------------------");

    let mut for_check_even_nums: i8 = 0;
    while for_check_even_nums < 101 {
        let is_even = for_check_even_nums % 2 == 0;
        if is_even {
            println!("{} is even", for_check_even_nums)
        }
        for_check_even_nums += 1;
    }

    println!("----------------------------------------");

    // tuple
    let my_string_tuple = ("rice", "fish", "crab", "salad");
    println!("{:?}", my_string_tuple);

    let my_int_tuple = (2, 4 , 6, 8);
    println!("{:?}", my_int_tuple);

    let my_mixed_tuple = ("crab", 7, 3.14, 'x', 'k', true);
    println!("{:?}", my_mixed_tuple);

    println!("best number is {}, PI is {}, best animal is {}, best words are {} and {}.",
             my_mixed_tuple.1, my_mixed_tuple.2, my_mixed_tuple.0,
             my_mixed_tuple.3, my_mixed_tuple.4);

    let my_nested_tuple = (1, 2, true, false, "hi", (1, 3, 3, 9));
    println!("this will be nine: {}", (my_nested_tuple.5).3);

    println!("----------------------------------------");
    add_num(35, 7);
    add_num(15, 7);
    add_num(10, 29);
    add_num(10, 1);

    let device1:ip_addr_versions = ip_addr_versions::V6; // I have ipv4.

    match device1 {
        ip_addr_versions::V6 => println!("everything is good!"),
        ip_addr_versions::V4 => println!("change your useless ip."),
    }
}

fn add_num(x: i8, y: i8) {
    let z = x + y;
    println!("{} + {} = {}", x, y, z);
}
