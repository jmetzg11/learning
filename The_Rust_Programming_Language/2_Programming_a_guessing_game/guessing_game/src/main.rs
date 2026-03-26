use std::io;

use rand:Rng;

// On Generating a Random Number
// https://doc.rust-lang.org/book/ch02-00-guessing-game-tutorial.html#generating-a-random-number
fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().get_range(1..=100);

    println!("The secret number is: {secret_number");

    println!("Pl")

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {guess}")
}
